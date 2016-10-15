import (
	"io"
	"log"
)

type LogFS struct {
	FS     // HL
	Name   string
	Logger *log.Logger
}

// Open implements FS.  All calls to Open are logged and errors are logged seperately.
func (l LogFS) Open(path string) (io.ReadCloser, error) {
	l.Logger.Printf("%v: open: %v", l.Name, path) // HL
	rc, err := l.FS.Open(path)
	if err != nil {
		l.Logger.Printf("%v: error %v: %v", l.Name, path, err) // HL
	}
	return rc, err
}
