// PrefixFS creates a FS which wraps an FS and prefixes all paths with Prefix.
type PrefixFS struct {
	FS     // HL
	Prefix string
}

func (p PrefixFS) Open(path string) (io.ReadCloser, error) {
	return p.FS.Open(filepath.Join(p.Prefix, path)) // HL
}
