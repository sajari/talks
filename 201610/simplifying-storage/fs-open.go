// FS is an interface which abstracts common file system functionality.
type FS interface {
	// Open accesses the data located at path.
	// The caller should close the returned io.ReadCloser when done.
	Open(path string) (io.ReadCloser, error)
}