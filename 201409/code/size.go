package main

import(
	"fmt"
	"io/ioutil"
	"./compression"
)

const(
	TEST_SMALL = "This is a random string to test simple compression of small strings"
)

func main() {

	// Small string
	input := []byte(TEST_SMALL)
// START OMIT	
	output, _ := compress.CompressLZ4(input)	
	fmt.Printf("Small string (LZ4):  input=%d, output=%d, ratio=%.2f \n", len(input), len(output), float64(len(output))/float64(len(input)))

	output, _ = compress.CompressGZIP(input)	
	fmt.Printf("Small string (GZIP):  input=%d, output=%d, ratio=%.2f \n", len(input), len(output), float64(len(output))/float64(len(input)))	

	output, _ = compress.CompressZLIB(input)	
	fmt.Printf("Small string (ZLIB):  input=%d, output=%d, ratio=%.2f \n", len(input), len(output), float64(len(output))/float64(len(input)))

	output, _ = compress.CompressFlate(input, compress.BestCompression)	
	fmt.Printf("Small string (Flate):  input=%d, output=%d, ratio=%.2f \n", len(input), len(output), float64(len(output))/float64(len(input)))
// END OMIT
	// Medium sie article
	input, err := ioutil.ReadFile("data/test.txt")
	if err != nil {
		return
	}

	output, _ = compress.CompressLZ4(input)	
	fmt.Printf("Article (LZ4):  input=%d, output=%d, ratio=%.2f \n", len(input), len(output), float64(len(output))/float64(len(input)))

	output, _ = compress.CompressGZIP(input)	
	fmt.Printf("Article (GZIP):  input=%d, output=%d, ratio=%.2f \n", len(input), len(output), float64(len(output))/float64(len(input)))	

	output, _ = compress.CompressZLIB(input)	
	fmt.Printf("Article (ZLIB):  input=%d, output=%d, ratio=%.2f \n", len(input), len(output), float64(len(output))/float64(len(input)))

	output, _ = compress.CompressFlate(input, compress.BestCompression)	
	fmt.Printf("Article (Flate):  input=%d, output=%d, ratio=%.2f \n", len(input), len(output), float64(len(output))/float64(len(input)))


	// Large article
	input, err = ioutil.ReadFile("data/Fundamental-Data-Structures.pdf")
	if err != nil {
		return
	}

	output, _ = compress.CompressLZ4(input)	
	fmt.Printf("Large (LZ4):  input=%d, output=%d, ratio=%.2f \n", len(input), len(output), float64(len(output))/float64(len(input)))

	output, _ = compress.CompressGZIP(input)	
	fmt.Printf("Large (GZIP):  input=%d, output=%d, ratio=%.2f \n", len(input), len(output), float64(len(output))/float64(len(input)))	

	output, _ = compress.CompressZLIB(input)	
	fmt.Printf("Large (ZLIB):  input=%d, output=%d, ratio=%.2f \n", len(input), len(output), float64(len(output))/float64(len(input)))

	output, _ = compress.CompressFlate(input, compress.BestCompression)	
	fmt.Printf("Large (Flate):  input=%d, output=%d, ratio=%.2f \n", len(input), len(output), float64(len(output))/float64(len(input)))
}