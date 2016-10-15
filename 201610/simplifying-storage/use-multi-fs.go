var (
	aws     = S3{Bucket: "model-files"}
	gcs     = CloudStorage("model-files")
	scratch = Local("/tmp/scratch-space")
)

fs := NewMultiFS(scratch, gcs, aws) // HL

// Will try scratch, then gcs, then aws.
rc, err := fs.Open("model.dat")