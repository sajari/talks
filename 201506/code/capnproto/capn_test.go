package main

import (
	"fmt"
	"testing"
)

var term = &Term{
	TermStr:      "本語 本語本語本語本語本語 whatever",
	Slot:         12345,
	NumDocuments: 2323244,
	NumWords:     3,
	Shotgun: []Shotgun{
		Shotgun{
			Term:    "shotgun",
			Potency: 0.44,
		},
		Shotgun{
			Term:    "shotgun two",
			Potency: 0.232333,
		},
	},
	Clues: []Clue{
		Clue{
			Term:    "clue",
			Intro:   "intro",
			Potency: 0.33,
		},
		Clue{
			Term:    "clue two",
			Intro:   "intro two",
			Potency: 0.22819,
		},
	},
	InteractionsPos: 23,
	InteractionsNeg: 456,
	HardcodedScore:  89,
	Infogain:        34.3232,
}

var simple_term = &Term{
	TermStr:         "本語 本語本語本語本語本語 whatever",
	Slot:            12345,
	NumDocuments:    2323244,
	NumWords:        3,
	InteractionsPos: 23,
	InteractionsNeg: 456,
	HardcodedScore:  89,
	Infogain:        34.3232,
}
var revs *Revs

func init() {
	N := 1000
	revs = &Revs{
		Data: make([]Rev, N),
	}
	for i := 0; i < N; i++ {
		revs.Data[i] = Rev{
			DocId:  12344,
			Rank:   1234,
			InMeta: 2347,
		}
	}

}

func TestCapNDecodeTerm(t *testing.T) {
	encoded := term.Encode("")
	term2 := new(Term)
	term2.Decode("", encoded)

	fmt.Printf("%v\n%v\n", term, term2)

	fmt.Printf("Encoded Length: %v\n", len(encoded))

	simple_encoded := simple_term.Encode("")
	fmt.Printf("Encoded Length (simple): %v\n", len(simple_encoded))
}

func BenchmarkCapNEncodeTerm(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		term.Encode("")
	}
}

func BenchmarkCapNDecodeTerm(b *testing.B) {
	b.ReportAllocs()
	encoded := term.Encode("")
	term2 := new(Term)
	for n := 0; n < b.N; n++ {
		term2.Decode("", encoded)
	}
}

func BenchmarkCapNEncodeSimpleTerm(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		simple_term.Encode("")
	}
}

func BenchmarkCapNDecodeSimpleTerm(b *testing.B) {
	b.ReportAllocs()
	encoded := simple_term.Encode("")
	term2 := new(Term)
	for n := 0; n < b.N; n++ {
		term2.Decode("", encoded)
	}
}

func BenchmarkCapNEncodeRev1000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		revs.Encode("")
	}
}

func BenchmarkCapNDecodeRev1000(b *testing.B) {
	encoded := revs.Encode("")
	revs2 := new(Revs)
	for n := 0; n < b.N; n++ {
		revs2.Decode("", encoded)
	}
}

func BenchmarkCapNZeroEncodeTerm(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		term.EncodeCapN("")
	}
}

func BenchmarkCapNZeroDecodeTerm(b *testing.B) {
	b.ReportAllocs()
	encoded := term.EncodeCapN("")
	termcapn2 := new(TermCapn)
	for n := 0; n < b.N; n++ {
		termcapn2.Decode("", encoded)
	}
}
