var (
	src     = HTTPFS("https://sajari.com/model-files")
	aws     = S3{Bucket: "model-files"}
	gcs     = CloudStorage("model-files")
	scratch = Local("/tmp/scratch-space")
)

func retry(fs FS) FS {
	return RetryFS{
		FS:   fs,
		N:    3,
		Wait: 1 * time.Second,
	}
}

fs := NewMultiFS(scratch, retry(gcs), retry(aws), retry(src)) // HL