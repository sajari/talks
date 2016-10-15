import (
	"golang.org/x/net/context"
	"golang.org/x/net/oauth2/google"
	"cloud.google.com/go/storage"
)

// ReadFromGCS reads the contents of a file from a Google Cloud Storage bucket.
func ReadFromGCS(bucket, path string) ([]byte, error) { // HL
	ctx := context.Background()
	ts, err := google.DefaultTokenSource(ctx, storage.ScopeReadOnly)
	if err != nil { /* need to handle */ }
	
	client, err := storage.NewClient(ctx, option.WithTokenSource(ts))
	if err != nil { /* need to handle */ }
	
	rc, err := client.Bucket(bucket).Object(path).NewReader(ctx) // HL
	if err != nil {
		return nil, fmt.Errorf("error fetching object reader for '%v': %v", path, err)
	}
	defer rc.Close()
	return ioutil.ReadAll(rc)
}