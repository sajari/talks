package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"expvar"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"golang.org/x/net/context"
	"golang.org/x/net/trace"
)

var (
	// ErrNotFound is returned when the requested resource cannot be found.
	ErrNotFound = errors.New("not found")

	// ErrPermissionDenied is returned when the caller does not have sufficient privileges to
	// complete a request.
	ErrPermissionDenied = errors.New("permission denied")

	// ErrUnavailable is returned when the requested resource is temporarily unavailable.
	ErrUnavailable = errors.New("not available")
)

// Mem is an in-memory FS implementation.
type Mem map[string][]byte

// Open implements FS.
func (m Mem) Open(ctx context.Context, path string) (io.ReadCloser, error) {
	b, ok := m[path]
	if !ok {
		return nil, ErrNotFound
	}
	return ioutil.NopCloser(bytes.NewReader(b)), nil
}

// MultiFS is an implementation of FS which uses an ordered list of FS implementations.
type MultiFS []FS

// Open implements FS.  All calls are directed to each FS in order until succesfull.
// In the case were every call fails then the last error is returned.
func (m MultiFS) Open(ctx context.Context, path string) (rc io.ReadCloser, err error) {
	for _, fs := range m {
		rc, err = fs.Open(ctx, path)
		if err == nil {
			return
		}
	}
	return
}

// NewMultiFS returns a MultiFS of the given FS implementations.
func NewMultiFS(fs ...FS) FS {
	return MultiFS(fs)
}

// FS is an interface which abstracts common file system functionality.
type FS interface {
	// Open accesses the data located at path.  The caller should close the returned
	// io.ReadCloser when done.
	Open(ctx context.Context, path string) (io.ReadCloser, error)
}

type Local string

func (l Local) Open(ctx context.Context, path string) (io.ReadCloser, error) {
	f, err := os.Open(filepath.Join(string(l), path))
	if os.IsNotExist(err) {
		return nil, ErrNotFound
	}
	if err == nil {
		if stat, err := f.Stat(); err == nil {
			if stat.IsDir() {
				f.Close()
				return nil, ErrNotFound
			}
		}
	}
	return f, err
}

// PrefixFS creates a FS which wraps an FS and prefixes all paths with Prefix.
type PrefixFS struct {
	FS     // HL
	Prefix string
}

func (p PrefixFS) Open(ctx context.Context, path string) (io.ReadCloser, error) {
	return p.FS.Open(ctx, filepath.Join(p.Prefix, path))
}

type LogFS struct {
	FS     // HL
	Name   string
	Logger *log.Logger
}

// Open implements FS.  All calls to Open are logged and errors are logged seperately.
func (l LogFS) Open(ctx context.Context, path string) (io.ReadCloser, error) {
	l.Logger.Printf("%v: open: %v", l.Name, path) // HL
	rc, err := l.FS.Open(ctx, path)
	if err != nil {
		l.Logger.Printf("%v: error: %v: %v", l.Name, path, err) // HL
	}
	return rc, err
}

// TraceFS is a type which wraps an FS, recording Open calls in a trace.
type TraceFS struct {
	FS
	Name string
}

func (t TraceFS) Open(ctx context.Context, path string) (rc io.ReadCloser, err error) {
	if tr, ok := trace.FromContext(ctx); ok {
		tr.LazyPrintf("%v: open: %v", t.Name, path)
		defer func() {
			if err != nil {
				tr.LazyPrintf("%v: error: %v", t.Name, err)
			}
		}()
	}
	return t.FS.Open(ctx, path) // HL
}

func NewStatCountFS(fs FS, name string, err error) FS {
	status := expvar.NewMap(name)        // HL
	status.Set("total", new(expvar.Int)) // HL
	status.Set("count", new(expvar.Int)) // HL
	return StatCountFS{
		FS:     fs,
		Status: status,
		Err:    err,
	}
}

type StatCountFS struct {
	FS
	Err    error
	Status *expvar.Map
}

func (s StatCountFS) Open(ctx context.Context, path string) (io.ReadCloser, error) {
	rc, err := s.FS.Open(ctx, path)
	if err == s.Err {
		s.Status.Add("count", 1) // HL
	}
	s.Status.Add("total", 1) // HL
	return rc, err
}

type DemoServer struct {
	FS
}

func (f DemoServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	x := struct {
		Path string `json:"path"`
	}{}
	if err := json.Unmarshal(b, &x); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tr := trace.New("demo-server", x.Path)
	defer tr.Finish()

	ctx := trace.NewContext(context.Background(), tr)
	rc, err := f.FS.Open(ctx, x.Path)
	if err != nil {
		status := http.StatusInternalServerError
		if err == ErrNotFound {
			status = http.StatusNotFound
		}
		http.Error(w, err.Error(), status)
		return
	}
	defer rc.Close()

	if n, err := io.Copy(w, rc); err != nil {
		log.Println("server: error writing response after %d bytes: %v", n, err)
	}
}

var DefaultLogger = log.New(os.Stderr, "", log.LstdFlags)

// START OMIT
func Debug(fs FS, name string) FS {
	fs = LogFS{FS: fs, Name: name, Logger: DefaultLogger} // HL
	fs = TraceFS{FS: fs, Name: name} // HL
	fs = NewStatCountFS(fs, fmt.Sprintf("storage(%v) not found", name), ErrNotFound) // HL
	return NewStatCountFS(fs, fmt.Sprintf("storage(%v) ok", name), nil) // HL
}

func main() {
	go func() {
		log.Println("Starting debug server on localhost:7778")
		log.Println("error from debug listen:", http.ListenAndServe("localhost:7778", nil))
	}()

	mem := Debug(Mem(map[string][]byte{
		"hello-world.txt": []byte("hello world!"),
	}), "mem-fs")
	local := Debug(Local("/tmp/demo"), "local:///tmp/demo")
	prefixLocal := Debug(PrefixFS{FS: local, Prefix: "gopher"}, "prefix(gopher)")

	fs := Debug(NewMultiFS(mem, local, prefixLocal), "multi") // HL

	log.Println("Starting demo HTTP server on localhost:7777")
	log.Fatal(http.ListenAndServe("localhost:7777", DemoServer{fs}))
}

// END OMIT
