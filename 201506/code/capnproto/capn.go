package main

import (
	"bytes"
	capn "github.com/glycerine/go-capnproto"
)

type Term struct {
	TermStr         string    `capid:"0"`
	Slot            uint32    `capid:"1"`
	NumDocuments    uint32    `capid:"2"`
	NumWords        uint8     `capid:"3"`
	Shotgun         []Shotgun `capid:"4"`
	Clues           []Clue    `capid:"5"`
	InteractionsPos uint16    `capid:"6"`
	InteractionsNeg uint16    `capid:"7"`
	HardcodedScore  int8      `capid:"8"` // -100 to +100 score strength overrides popularity function, Outside of -100 to 100 is ignored.
	Infogain        float32   `capid:"9"`
}

type Shotgun struct {
	Term    string  `capid:"0"`
	Potency float32 `capid:"1"`
}

type Clue struct {
	Term    string  `capid:"0"`
	Intro   string  `capid:"1"`
	Potency float32 `capid:"2"`
}

type Rev struct {
	DocId  uint32 `capid:"0"`
	Rank   uint16 `capid:"1"`
	InMeta uint16 `capid:"2"`
}

type Revs struct {
	Data []Rev `capid:"0"`
}

func (term *Term) Encode(version string) []byte {
	var o bytes.Buffer
	term.Save(&o)
	// now we have saved!
	return o.Bytes()
}

func (term *Term) Decode(version string, data []byte) {
	o := bytes.NewBuffer(data)
	term.Load(o)
}

// Use Cap-n-proto object instead -----------------------
func (term *Term) EncodeCapN(version string) []byte {
	seg := capn.NewBuffer(nil)
	termcapn := NewTermCapn(seg)
	termcapn.SetTermStr(term.TermStr)
	termcapn.SetSlot(term.Slot)
	termcapn.SetNumDocuments(term.NumDocuments)
	termcapn.SetNumWords(term.NumWords)

	// Shotgun -> ShotgunCapn (go slice to capn list)
	if len(term.Shotgun) > 0 {
		typedList := NewShotgunCapnList(seg, len(term.Shotgun))
		plist := capn.PointerList(typedList)
		i := 0
		for _, ele := range term.Shotgun {
			plist.Set(i, capn.Object(ShotgunGoToCapn(seg, &ele)))
			i++
		}
		termcapn.SetShotgun(typedList)
	}

	// Clues -> ClueCapn (go slice to capn list)
	if len(term.Clues) > 0 {
		typedList := NewClueCapnList(seg, len(term.Clues))
		plist := capn.PointerList(typedList)
		i := 0
		for _, ele := range term.Clues {
			plist.Set(i, capn.Object(ClueGoToCapn(seg, &ele)))
			i++
		}
		termcapn.SetClues(typedList)
	}

	termcapn.SetInteractionsPos(term.InteractionsPos)
	termcapn.SetInteractionsNeg(term.InteractionsNeg)
	termcapn.SetInfogain(term.Infogain)
	termcapn.SetHardcodedScore(term.HardcodedScore)
	buf := bytes.Buffer{}
	seg.WriteTo(&buf)
	return buf.Bytes()
}

func (term TermCapn) Decode(version string, data []byte) {
	seg, _, err := capn.ReadFromMemoryZeroCopy(data)
	if err != nil {
		panic(err)
	}
	term = ReadRootTermCapn(seg)
}

// -------------------------------------------------------

func (revs *Revs) Encode(version string) []byte {
	var o bytes.Buffer
	revs.Save(&o)
	// now we have saved!
	return o.Bytes()
}

func (revs *Revs) Decode(version string, data []byte) {
	o := bytes.NewBuffer(data)
	revs.Load(o)
}

func main() {

}
