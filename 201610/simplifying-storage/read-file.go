// ReadFile reads the file given by path and returns its contents.
func ReadFile(fs FS, path string) ([]byte, error) {
	rc, err := fs.Open(path)
	if err != nil {
		return nil, err
	}
	defer rc.Close()
	return ioutil.ReadAll(rc)
}