func NewStatCountFS(fs FS, name string, err error) FS {
	m := expvar.NewMap(name)        // HL
	m.Set("total", new(expvar.Int)) // HL
	m.Set("count", new(expvar.Int)) // HL
	return StatCountFS{
		FS:    fs,
		Stats: m,
		Err:   err,
	}
}

type StatCountFS struct {
	FS          // HL
	Err   error // The error to count.
	Stats *expvar.Map
}

func (s StatCountFS) Open(path string) (io.ReadCloser, error) {
	rc, err := s.FS.Open(path)
	if err == s.Err { s.Stats.Add("count", 1) } // HL
	s.Stats.Add("total", 1) // HL
	return rc, err
}
