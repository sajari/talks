// HTTP is an implementation of FS which fetches paths from HTTP servers.
type HTTP string

func (h HTTP) Open(path string) (io.ReadCloser, error) {
	resp, err := http.Get(fmt.Sprintf("%v%v", h, path)) // HL
	if err != nil { /* handle error */ }

	switch resp.Status {
	case http.StatusNotFound: // 404
		return nil, ErrNotFound // HL

	case http.StatusServiceUnavailable: // 503
		return nil, ErrUnavailable // HL
	// ...
	}	
	return resp.Body, nil
}
