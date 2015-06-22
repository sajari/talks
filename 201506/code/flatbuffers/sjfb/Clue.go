// automatically generated, do not modify

package sjfb

import (
	flatbuffers "github.com/google/flatbuffers/go"
)
type Clue struct {
	_tab flatbuffers.Table
}

func (rcv *Clue) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Clue) Term() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Clue) Intro() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Clue) Potency() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0
}

func ClueStart(builder *flatbuffers.Builder) { builder.StartObject(3) }
func ClueAddTerm(builder *flatbuffers.Builder, Term flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(Term), 0) }
func ClueAddIntro(builder *flatbuffers.Builder, Intro flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(Intro), 0) }
func ClueAddPotency(builder *flatbuffers.Builder, Potency float32) { builder.PrependFloat32Slot(2, Potency, 0) }
func ClueEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT { return builder.EndObject() }
