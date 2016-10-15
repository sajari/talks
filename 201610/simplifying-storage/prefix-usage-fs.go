func UserProjectFS(fs FS, user, project string) FS {
	return PrefixFS{FS: fs, Prefix: filepath.Join(user, project)}
}