package main

import (
	"fmt"
	"testing"
)

const CURRENT_VERSION = ""

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

func BenchmarkFlatbuffersEncodeTerm(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		term.Encode("")
	}
}

func BenchmarkFlatbuffersDecodeTerm(b *testing.B) {
	b.ReportAllocs()
	encoded := term.Encode("")
	term2 := new(Term)
	for n := 0; n < b.N; n++ {
		term2.Decode("", encoded)
	}
}

func BenchmarkFlatbuffersEncodeSimpleTerm(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		simple_term.Encode("")
	}
}

func BenchmarkFlatbuffersDecodeSimpleTerm(b *testing.B) {
	b.ReportAllocs()
	encoded := simple_term.Encode("")
	term2 := new(Term)
	for n := 0; n < b.N; n++ {
		term2.Decode("", encoded)
	}
}

func TestFlatbuffersNDecodeTerm(t *testing.T) {
	encoded := term.Encode("")
	term2 := new(Term)
	term2.Decode("", encoded)

	fmt.Printf("%v\n%v\n", term, term2)

	fmt.Printf("Encoded Length: %v\n", len(encoded))

	simple_encoded := simple_term.Encode("")
	fmt.Printf("Encoded Length (simple): %v\n", len(simple_encoded))
}

func TestEncodeTerm(t *testing.T) {
	encoded := term.Encode(CURRENT_VERSION)

	term2 := new(Term)
	term2.Decode(CURRENT_VERSION, encoded)

	fmt.Println(term, "\n", string(encoded), "\n", term2)

	if term.TermStr != term2.TermStr {
		t.Errorf("Failed to encode and decode term. Got %v, expected %v", term2.TermStr, term.TermStr)
	}
	if term.Slot != term2.Slot {
		t.Errorf("Failed to encode and decode doc Slot. Got %v, expected %v", term2.Slot, term.Slot)
	}
	if term.NumDocuments != term2.NumDocuments {
		t.Errorf("Failed to encode and decode doc NumDocuments. Got %v, expected %v", term2.NumDocuments, term.NumDocuments)
	}
	if term.NumWords != term2.NumWords {
		t.Errorf("Failed to encode and decode doc NumWords. Got %v, expected %v", term2.NumWords, term.NumWords)
	}
	if term.InteractionsPos != term2.InteractionsPos {
		t.Errorf("Failed to encode and decode doc InteractionsPos. Got %v, expected %v", term2.InteractionsPos, term.InteractionsPos)
	}
	if term.InteractionsNeg != term2.InteractionsNeg {
		t.Errorf("Failed to encode and decode doc InteractionsNeg. Got %v, expected %v", term2.InteractionsNeg, term.InteractionsNeg)
	}
	if term.HardcodedScore != term2.HardcodedScore {
		t.Errorf("Failed to encode and decode doc HardcodedScore. Got %v, expected %v", term2.HardcodedScore, term.HardcodedScore)
	}
	if term.Infogain != term2.Infogain {
		t.Errorf("Failed to encode and decode doc Infogain. Got %v, expected %v", term2.Infogain, term.Infogain)
	}

	// Shotgun
	if len(term.Shotgun) != len(term2.Shotgun) {
		t.Errorf("Failed. Shotgun count has changed after decode")
	} else {
		if term.Shotgun[1].Term != term2.Shotgun[1].Term {
			t.Errorf("Failed to encode and decode doc Shotgun term. Got %v, expected %v", term2.Shotgun[1].Term, term.Shotgun[1].Term)
		}
		if term.Shotgun[1].Potency != term2.Shotgun[1].Potency {
			t.Errorf("Failed to encode and decode doc Shotgun potency. Got %v, expected %v", term2.Shotgun[1].Potency, term.Shotgun[1].Potency)
		}
	}

	// Clues
	if term.Clues[1].Term != term2.Clues[1].Term {
		t.Errorf("Failed to encode and decode doc Clue term. Got %v, expected %v", term2.Clues[1].Term, term.Clues[1].Term)
	}
	if term.Clues[1].Intro != term2.Clues[1].Intro {
		t.Errorf("Failed to encode and decode doc Clue intro. Got %v, expected %v", term2.Clues[1].Intro, term.Clues[1].Intro)
	}
	if term.Clues[1].Potency != term2.Clues[1].Potency {
		t.Errorf("Failed to encode and decode doc Clue potency. Got %v, expected %v", term2.Clues[1].Potency, term.Clues[1].Potency)
	}

}
