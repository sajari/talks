// automatically generated, do not modify

package sjfb

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Term struct {
	_tab flatbuffers.Table
}

func GetRootAsTerm(buf []byte, offset flatbuffers.UOffsetT) *Term {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Term{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Term) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Term) TermStr() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Term) Slot() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Term) NumDocuments() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Term) NumWords() int16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetInt16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Term) Shotgun(obj *Shotgun, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		if obj == nil {
			obj = new(Shotgun)
		}
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *Term) ShotgunLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Term) Clues(obj *Clue, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		if obj == nil {
			obj = new(Clue)
		}
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *Term) CluesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Term) InteractionPos() int16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.GetInt16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Term) InteractionNeg() int16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.GetInt16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Term) HardcodedScore() int16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return rcv._tab.GetInt16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Term) Infogain() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0
}

func TermStart(builder *flatbuffers.Builder) { builder.StartObject(10) }
func TermAddTermStr(builder *flatbuffers.Builder, TermStr flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(TermStr), 0)
}
func TermAddSlot(builder *flatbuffers.Builder, Slot uint32) { builder.PrependUint32Slot(1, Slot, 0) }
func TermAddNumDocuments(builder *flatbuffers.Builder, NumDocuments uint32) {
	builder.PrependUint32Slot(2, NumDocuments, 0)
}
func TermAddNumWords(builder *flatbuffers.Builder, NumWords int16) {
	builder.PrependInt16Slot(3, NumWords, 0)
}
func TermAddShotgun(builder *flatbuffers.Builder, Shotgun flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(Shotgun), 0)
}
func TermStartShotgunVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func TermAddClues(builder *flatbuffers.Builder, Clues flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(Clues), 0)
}
func TermStartCluesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func TermAddInteractionPos(builder *flatbuffers.Builder, InteractionPos int16) {
	builder.PrependInt16Slot(6, InteractionPos, 0)
}
func TermAddInteractionNeg(builder *flatbuffers.Builder, InteractionNeg int16) {
	builder.PrependInt16Slot(7, InteractionNeg, 0)
}
func TermAddHardcodedScore(builder *flatbuffers.Builder, HardcodedScore int16) {
	builder.PrependInt16Slot(8, HardcodedScore, 0)
}
func TermAddInfogain(builder *flatbuffers.Builder, Infogain float32) {
	builder.PrependFloat32Slot(9, Infogain, 0)
}
func TermEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT { return builder.EndObject() }
