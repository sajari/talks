// Mem is an in-memory FS implementation.
type Mem map[string][]byte

// Open implements FS.
func (m Mem) Open(path string) (io.ReadCloser, error) {
	b, ok := m[path]
	if !ok {
		return nil, ErrNotFound
	}
	return ioutil.NopCloser(bytes.NewReader(b)), nil
}
