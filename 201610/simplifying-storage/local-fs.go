import (
	"io"
	"os"
	"path/filepath"
)

// Local is an FS implementation which uses the local file system.
type Local string

// Open implements FS.
func (l Local) Open(path string) (io.ReadCloser, error) {
	return os.Open(filepath.Join(string(l), path))
}
