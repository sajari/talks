// UserProjectFS creates an FS scoped to a particular user/project.
func UserProjectFS(fs FS, userID, projectID string) FS {
	return PrefixFS{
		FS:     fs,
		Prefix: fmt.Sprintf("%v/%v", userID, projectID),
	}
}
