package compress 

import (
	"testing"
	"time"
	"encoding/binary"
	"math/rand"
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	)

const(
	TEST_SMALL = "This is a random string to test simple compression of small strings"
)

// Random multi field struct
type Rev struct {
	DocumentId 		uint32
	Rank 			uint16
	InMeta 			uint16
	Next 			uint32 
}

func randomRev() *Rev {
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	r := new(Rev)
	r.DocumentId = uint32(rand.Int31())
	r.Rank = uint16(rand.Int31n(65355))
	r.InMeta = uint16(rand.Int31n(65355))
	r.Next = uint32(rand.Int31())
	return r
}

func randomRevBytes(n int) []byte {
	buf := bytes.NewBuffer([]byte{})
	for i := 0; i < 100; i++ {
		_ = binary.Write(buf, binary.LittleEndian, randomRev())	
	}
	return buf.Bytes()
}

func randomMap() map[string]string {
	meta := make(map[string]string)
	meta["title"] = "Water cooler technician"
	meta["date"] = "134402030428"
	meta["description"] = "Aqua Cooler is Australia's leading manufacturer of water coolers and industrial water cooling systems with an enviable record of quality products and service support. Located at Chester Hill our office is very close to public transport. Due to an exciting growth phase we are seeking a person to assist our service department with data entry, logging service jobs, invoicing and general office support. This sixteen hours per week part time position is available immediately for the right person who can work independently, be comfortable with software such as Excel, Word, Lotus Notes and Micronet (not mandatory).  Accuracy and attention to detail are critical to this role as well as a desire to provide our clients with the best possible service support. The successful applicant will also be well presented with satisfactory written and verbal skills."
	meta["bayes-industry"] = "Information and communication science"
	meta["location"] = "Alexandria, VA, USA"
	meta["lat"] = "38.2234288"
	meta["lng"] = "155.2004833"
	meta["url"] = "http://www.example.com/Jobs/JobDetails.aspx?siteid=homepage&Job_DID=JHL43F6VMH90DSZBY11&ipath=HPRJ"
	return meta
}

func encodeGob(input interface{}) []byte {
	var buf bytes.Buffer 
	// Create an encoder and send a value.
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(input)
	if err != nil {
		fmt.Println("encode:", err)
	}
	return buf.Bytes()
}

var result []byte

func benchmarkCompress(algo string, input []byte, rate int, b *testing.B) {
	var output []byte
	for n := 0; n < b.N; n++ {
		switch (algo) {
			case "lz4":
				output, _ = CompressLZ4(input)
			case "zlib":
				output, _ = CompressZLIB(input)
			case "gzip":
				output, _ = CompressGZIP(input)
			case "flate":
				output, _ = CompressFlate(input, rate)			
		}
    }
    result = output // This prevents compiler optimisation causing issues.
}

func BenchmarkSimpleLZ4(b *testing.B) { benchmarkCompress("lz4", []byte(TEST_SMALL), 0, b) }
func BenchmarkSimpleGZIP(b *testing.B) { benchmarkCompress("gzip", []byte(TEST_SMALL), 0, b) }
func BenchmarkSimpleZLIB(b *testing.B) { benchmarkCompress("zlib", []byte(TEST_SMALL), 0, b) }
func BenchmarkSimpleFlateHigh(b *testing.B) { benchmarkCompress("flate", []byte(TEST_SMALL), BestCompression, b) }
func BenchmarkSimpleFlateFast(b *testing.B) { benchmarkCompress("flate", []byte(TEST_SMALL), BestSpeed, b) }


func BenchmarkStructLZ4(b *testing.B) { benchmarkCompress("lz4", randomRevBytes(1000), 0, b) }
func BenchmarkStructGZIP(b *testing.B) { benchmarkCompress("gzip", randomRevBytes(1000), 0, b) }
func BenchmarkStructZLIB(b *testing.B) { benchmarkCompress("zlib", randomRevBytes(1000), 0, b) }
func BenchmarkStructFlateHigh(b *testing.B) { benchmarkCompress("flate", randomRevBytes(1000), BestCompression, b) }
func BenchmarkStructFlateFast(b *testing.B) { benchmarkCompress("flate", randomRevBytes(1000), BestSpeed, b) }

func BenchmarkMapGobLZ4(b *testing.B) { benchmarkCompress("lz4", encodeGob(randomMap()), 0, b) }
func BenchmarkMapGobGZIP(b *testing.B) { benchmarkCompress("gzip", encodeGob(randomMap()), 0, b) }
func BenchmarkMapGobZLIB(b *testing.B) { benchmarkCompress("zlib", encodeGob(randomMap()), 0, b) }
func BenchmarkMapGobFlateHigh(b *testing.B) { benchmarkCompress("flate", encodeGob(randomMap()), BestCompression, b) }
func BenchmarkMapGobFlateFast(b *testing.B) { benchmarkCompress("flate", encodeGob(randomMap()), BestSpeed, b) }

func BenchmarkFileLZ4(b *testing.B) { 
	input, _ := ioutil.ReadFile("../data/Fundamental-Data-Structures.pdf")
	benchmarkCompress("lz4", input, 0, b) 
}
func BenchmarkFileGZIP(b *testing.B) { 
	input, _ := ioutil.ReadFile("../data/Fundamental-Data-Structures.pdf")
	benchmarkCompress("gzip", input, 0, b) 
}
func BenchmarkFileZLIB(b *testing.B) { 
	input, _ := ioutil.ReadFile("../data/Fundamental-Data-Structures.pdf")
	benchmarkCompress("zlib", input, 0, b) 
}
func BenchmarkFileFlateHigh(b *testing.B) { 
	input, _ := ioutil.ReadFile("../data/Fundamental-Data-Structures.pdf")
	benchmarkCompress("flate", input, BestCompression, b) 
}
func BenchmarkFileFlateFast(b *testing.B) { 
	input, _ := ioutil.ReadFile("../data/Fundamental-Data-Structures.pdf")
	benchmarkCompress("flate", input, BestSpeed, b) 
}




