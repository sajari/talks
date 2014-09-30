package compress


import (
	"github.com/pwaller/go-clz4"
	"compress/gzip"
	"compress/zlib"
	"compress/flate"
	"bytes"
)

const (
        NoCompression = 0
        BestSpeed     = 1

        BestCompression    = 9
        DefaultCompression = -1
)

func CompressLZ4(input []byte) (output []byte, err error) {
	err = clz4.Compress(input, &output)
	return output, err
}

func CompressGZIP(input []byte) (output []byte, err error) {
	var buf bytes.Buffer
	w := gzip.NewWriter(&buf)
	w.Write(input)
	w.Close()
	return buf.Bytes(), nil
}

func CompressZLIB(input []byte) (output []byte, err error) {
	var buf bytes.Buffer
	w := zlib.NewWriter(&buf)
	w.Write(input)
	w.Close()
	return buf.Bytes(), nil
}

func CompressFlate(input []byte, compression int) (output []byte, err error) {
	var buf bytes.Buffer
	w, err := flate.NewWriter(&buf, compression)
	w.Write(input)
	w.Close()
	return buf.Bytes(), nil
}
