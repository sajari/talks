// CloudStorage is an implementation of FS that uses a Google Cloud Storage bucket.
type CloudStorage string

// Open implements FS.
func (c CloudStorage) Open(path string) (io.ReadCloser, error) {
	ctx := context.Background()
	ts, err := google.DefaultTokenSource(ctx, scope)
	if err != nil { /* handle error */ }

	client, err := storage.NewClient(ctx, option.WithTokenSource(ts))
	if err != nil { /* handle error */ }

	bh, err := client.Bucket(c).bucketHandle(ctx, storage.ScopeReadOnly)
	if err != nil { /* handle error */ }

	return bh.Object(path).NewReader(ctx)
}
