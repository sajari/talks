package main

import (
	//"fmt"
	flatbuffers "github.com/google/flatbuffers/go"
	"talks/201506/code/flatbuffers/sjfb"
)

/*
type Term struct {
	sjfb.Term
}
*/

type Term struct {
	TermStr         string
	Slot            uint32
	NumDocuments    uint32
	NumWords        uint8
	Shotgun         []Shotgun
	Clues           []Clue
	InteractionsPos uint16
	InteractionsNeg uint16
	HardcodedScore  int8
	Infogain        float32
}

type Shotgun struct {
	Term    string
	Potency float32
}

type Clue struct {
	Term    string
	Intro   string
	Potency float32
}

func (term *Term) Encode(version string) []byte {
	builder := flatbuffers.NewBuilder(0)

	// Preprocess the shotguns
	shotguns := make([]flatbuffers.UOffsetT, len(term.Shotgun))
	for i := 0; i < len(term.Shotgun); i++ {
		sjfb.ShotgunStart(builder)
		sjfb.ShotgunAddTerm(builder, builder.CreateString(term.Shotgun[i].Term))
		sjfb.ShotgunAddPotency(builder, term.Shotgun[i].Potency)
		shotguns[i] = sjfb.ShotgunEnd(builder)
	}
	sjfb.TermStartShotgunVector(builder, len(term.Shotgun))
	for i := len(term.Shotgun) - 1; i >= 0; i-- {
		builder.PrependUOffsetT(shotguns[i])
	}
	shotgun_vec := builder.EndVector(len(term.Shotgun))
	// Preprocess the clues
	clues := make([]flatbuffers.UOffsetT, len(term.Clues))
	//for i := len(term.Clues) - 1; i >= 0; i-- {
	for i := 0; i < len(term.Clues); i++ {
		sjfb.ClueStart(builder)
		sjfb.ClueAddTerm(builder, builder.CreateString(term.Clues[i].Term))
		sjfb.ClueAddIntro(builder, builder.CreateString(term.Clues[i].Intro))
		sjfb.ClueAddPotency(builder, term.Clues[i].Potency)
		clues[i] = sjfb.ClueEnd(builder)
	}
	sjfb.TermStartCluesVector(builder, len(term.Clues))
	for i := len(term.Clues) - 1; i >= 0; i-- {
		builder.PrependUOffsetT(clues[i])
	}
	clues_vec := builder.EndVector(len(term.Clues))
	// Start packing the final term
	sjfb.TermStart(builder)

	termStr := builder.CreateString(term.TermStr)
	sjfb.TermAddTermStr(builder, termStr)
	sjfb.TermAddSlot(builder, term.Slot)
	sjfb.TermAddNumDocuments(builder, term.NumDocuments)
	sjfb.TermAddNumWords(builder, int16(term.NumWords))
	sjfb.TermAddShotgun(builder, shotgun_vec)
	sjfb.TermAddClues(builder, clues_vec)

	sjfb.TermAddInteractionPos(builder, int16(term.InteractionsPos))
	sjfb.TermAddInteractionNeg(builder, int16(term.InteractionsNeg))
	sjfb.TermAddHardcodedScore(builder, int16(term.HardcodedScore))
	sjfb.TermAddInfogain(builder, term.Infogain)

	obj := sjfb.TermEnd(builder)
	builder.Finish(obj)

	return builder.Bytes[builder.Head():]
}

func (term *Term) Decode(version string, data []byte) {

	t := sjfb.Term{}
	t.Init(data, flatbuffers.GetUOffsetT(data))
	term.TermStr = string(t.TermStr())
	term.Slot = t.Slot()
	term.NumDocuments = t.NumDocuments()
	term.NumWords = uint8(t.NumWords())

	//fmt.Println("LEN: ", t.ShotgunLength())

	term.Shotgun = make([]Shotgun, t.ShotgunLength())
	/*
		shotvecs := make([]*sjfb.Shotgun, t.ShotgunLength())
		for i := 0; i < t.ShotgunLength(); i++ {
			t.Shotgun(shotvecs[i], i)
			term.Shotgun[i].Term = string(shotvecs[i].Term())
			term.Shotgun[i].Potency = shotvecs[i].Potency()
		}
	*/

	//term.Clues = t.Clues(obj, j)

	term.InteractionsPos = uint16(t.InteractionPos())
	term.InteractionsNeg = uint16(t.InteractionNeg())
	term.HardcodedScore = int8(t.HardcodedScore())
	term.Infogain = t.Infogain()

}

func main() {

}
