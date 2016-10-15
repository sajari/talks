import (
	"golang.org/x/net/context"
	"golang.org/x/net/trace" // HL
)

// TraceFS is a type which wraps an FS, recording Open calls in a trace.
type TraceFS struct {
	FS // HL
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
