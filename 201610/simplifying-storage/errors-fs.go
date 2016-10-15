var (
	// ErrNotFound is returned when the requested resource cannot be found.
	ErrNotFound = errors.New("not found")

	// ErrPermissionDenied is returned when the caller does not have sufficient privileges to
	// complete a request.
	ErrPermissionDenied = errors.New("permission denied")

	// ErrUnavailable is returned when the requested resource is temporarily unavailable.
	ErrUnavailable = errors.New("not available")
)