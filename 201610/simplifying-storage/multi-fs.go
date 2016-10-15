// MultiFS is an FS implementation defined by an ordered list of FS implementations.
type MultiFS []FS

// Open implements FS.  All calls are directed to each FS in order until succesfull.
// In the case were every call fails then the last error is returned.
func (m MultiFS) Open(path string) (rc io.ReadCloser, err error) {
	for _, fs := range m {
		rc, err = fs.Open(path)
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
