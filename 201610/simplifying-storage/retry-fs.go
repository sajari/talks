type RetryFS struct {
	FS                 // HL
	N    int           // Number of retries before giving up.
	Wait time.Duration // Duration to wait between each attempt.
}

func (r RetryFS) Open(path string) (io.ReadCloser, error) {
	for i := 0; i < N; i++ {
		if rc, err := r.FS.Open(path); err != ErrUnavailable { // HL
			return rc, err
		}
		time.Sleep(r.Wait)
	}
	return nil, ErrUnavailable
}