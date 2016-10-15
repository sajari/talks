import (
	"io/ioutil"
	"os"
)

// ReadFromLocal reads the contents of a file on the local filesystem.
func ReadFromLocal(path string) ([]byte, error) { // HL
	f, err := os.Open(path) // HL
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return ioutil.ReadAll(f)
}
