rc, err := S3{
	Auth:   // ...
	Region: aws.APSoutheast2,
	Bucket: "my-bucket",
}.Open("somepath/somefile.dat")