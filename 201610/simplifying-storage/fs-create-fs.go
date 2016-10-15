// RWFS is an interface which defines common file system functionality.
type RWFS interface {
	// Open accesses the data located at path.  The caller should close the returned
	// io.ReadCloser when done.
	Open(ctx context.Context, path string) (io.ReadCloser, error)

	// Create creates a file at the given path.  The caller should close
	// the returned io.WriteCloser when done to ensure that the file is written.
	Create(ctx context.Context, path string) (io.WriteCloser, error) // HL
	
	// Delete removes the file at the given path.
	Delete(ctx context.Context, path string) error // HL
}