import (
	"golang.org/x/net/context"
	"github.com/mitchellh/goamz/aws"
	"github.com/mitchellh/goamz/s3"
)

// ReadFromS3 reads the contents of a file from an S3 bucket using the given region and credentials.
func ReadFromS3(region aws.Region, auth aws.Auth, bucket, path string) ([]byte, error) { // HL
	rc, err := s3.New(auth, region).Bucket(bucket).GetReader(path) // HL
	if err != nil {
		return nil, err
	}
	defer rc.Close()
	return ioutil.ReadAll(rc)
}