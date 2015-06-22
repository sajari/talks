package main

import (
	"fmt"
	"github.com/gogo/protobuf/proto"
	"time"
	"unsafe"
)

/*
type Term struct {
	TermStr         string
	Slot            uint32
	NumDocuments    uint32
	NumWords        uint8
	Shotgun         []Shotgun
	Clues           []Clue
	InteractionsPos uint16
	InteractionsNeg uint16
	HardcodedScore  int8 // -100 to +100 score strength overrides popularity function, Outside of -100 to 100 is ignored.
	Infogain        float32
	sync.RWMutex
}


type Shotgun struct {
	PartitionID uint16
	Slot        uint32
}

type Clue struct {
	PartitionID uint16
	Slot        uint32
}
*/

func EncodeTerm(term *Term) ([]byte, error) {
	return proto.Marshal(term)
}

func DecodeTerm(input []byte) (*Term, error) {
	term := new(Term)
	err := proto.Unmarshal(input, term)
	return term, err
}

func main() {
	var t Term
	t.TermStr = proto.String("sup sup!")
	t.NumWords = proto.Int32(78)
	t.NumDocuments = proto.Uint32(78)
	t.Infogain = proto.Float32(40.3)
	t.InteractionsNeg = proto.Int32(4)
	t.InteractionsNeg = proto.Int32(2)
	t0 := time.Now()
	data, err := EncodeTerm(&t)
	t1 := time.Now()
	fmt.Println("\n\nEncoding time: ", t1.Sub(t0))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Original (%v) GoGoProto: (%v) : %v\n", unsafe.Sizeof(t), len(data), data)

	// Decode
	t2 := time.Now()
	nt, err := DecodeTerm(data)
	t3 := time.Now()
	fmt.Println("Decoding time: ", t3.Sub(t2))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s", nt.GetTermStr())

	t.TermStr = proto.String("changing this")
	t.NumWords = proto.Int32(77)
	t.NumDocuments = proto.Uint32(48)
	t.Infogain = proto.Float32(20.3)
	t.InteractionsNeg = proto.Int32(3)
	t.InteractionsNeg = proto.Int32(3)

	t0 = time.Now()
	data2, err := EncodeTerm(&t)
	t1 = time.Now()
	fmt.Println("\n\nEncoding time: ", t1.Sub(t0))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Original (%v) GoGoProto: (%v) : %v\n", unsafe.Sizeof(t), len(data2), data2)

	// Decode
	t2 = time.Now()
	nt, err = DecodeTerm(data2)
	t3 = time.Now()
	fmt.Println("Decoding time: ", t3.Sub(t2))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s", nt.GetTermStr())
}
