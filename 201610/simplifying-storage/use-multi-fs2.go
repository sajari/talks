var (
	src     = HTTP("https://sajari.com/model-files") // HL
	aws     = S3{Bucket: "model-files"}
	gcs     = CloudStorage("model-files")
	scratch = Local("/tmp/scratch-space")
)

fs := NewMultiFS(scratch, gcs, aws, src) // HL

// Will try scratch, then gcs, aws, then src.
rc, err := fs.Open("model.dat")
