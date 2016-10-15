b1, _ := ReadFile(Local("/tmp"), "somewhere/test.dat")
b2, _ := ReadFile(CloudStorage("bucket"), "somewhere/test.dat")
b3, _ := ReadFile(S3{Bucket:"bucket", Region: aws.APSouthEast2, Auth: /* ... */}, "somewhere/test.dat")