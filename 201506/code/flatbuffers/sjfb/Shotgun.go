// automatically generated, do not modify

package sjfb

import (
	flatbuffers "github.com/rw/flatbuffers/go"
)

type Shotgun struct {
	_tab flatbuffers.Table
}

func (rcv *Shotgun) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Shotgun) Term() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Shotgun) Potency() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0
}

func ShotgunStart(builder *flatbuffers.Builder) { builder.StartObject(2) }
func ShotgunAddTerm(builder *flatbuffers.Builder, Term flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(Term), 0)
}
func ShotgunAddPotency(builder *flatbuffers.Builder, Potency float32) {
	builder.PrependFloat32Slot(1, Potency, 0)
}
func ShotgunEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT { return builder.EndObject() }
