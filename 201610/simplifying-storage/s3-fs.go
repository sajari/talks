// S3 is an implementation of FS that uses an S3 bucket.
type S3 struct {
	Auth   aws.Auth
	Region aws.Region

	Bucket string
}

// Open implements FS.
func (s S3) Open(path string) (io.ReadCloser, error) {
	return s3.New(s.Auth, s.Region).Bucket(s.Bucket).GetReader(path)
}
