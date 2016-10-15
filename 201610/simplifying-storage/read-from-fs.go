func ReadFile(fs FS, path string) ([]byte, error) {
	rc, err := fs.Open(path) // HL
	if err != nil {
		return nil, err
	}
	defer rc.Close()
	return ioutil.ReadAll(rc)
}