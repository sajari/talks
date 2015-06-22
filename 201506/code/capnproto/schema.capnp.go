package main

// AUTO GENERATED - DO NOT EDIT

import (
	C "github.com/glycerine/go-capnproto"
	"math"
	"unsafe"
)

type ClueCapn C.Struct

func NewClueCapn(s *C.Segment) ClueCapn      { return ClueCapn(s.NewStruct(8, 2)) }
func NewRootClueCapn(s *C.Segment) ClueCapn  { return ClueCapn(s.NewRootStruct(8, 2)) }
func AutoNewClueCapn(s *C.Segment) ClueCapn  { return ClueCapn(s.NewStructAR(8, 2)) }
func ReadRootClueCapn(s *C.Segment) ClueCapn { return ClueCapn(s.Root(0).ToStruct()) }
func (s ClueCapn) Term() string              { return C.Struct(s).GetObject(0).ToText() }
func (s ClueCapn) SetTerm(v string)          { C.Struct(s).SetObject(0, s.Segment.NewText(v)) }
func (s ClueCapn) Intro() string             { return C.Struct(s).GetObject(1).ToText() }
func (s ClueCapn) SetIntro(v string)         { C.Struct(s).SetObject(1, s.Segment.NewText(v)) }
func (s ClueCapn) Potency() float32          { return math.Float32frombits(C.Struct(s).Get32(0)) }
func (s ClueCapn) SetPotency(v float32)      { C.Struct(s).Set32(0, math.Float32bits(v)) }

// capn.JSON_enabled == false so we stub MarshallJSON().
func (s ClueCapn) MarshalJSON() (bs []byte, err error) { return }

type ClueCapn_List C.PointerList

func NewClueCapnList(s *C.Segment, sz int) ClueCapn_List {
	return ClueCapn_List(s.NewCompositeList(8, 2, sz))
}
func (s ClueCapn_List) Len() int          { return C.PointerList(s).Len() }
func (s ClueCapn_List) At(i int) ClueCapn { return ClueCapn(C.PointerList(s).At(i).ToStruct()) }
func (s ClueCapn_List) ToArray() []ClueCapn {
	return *(*[]ClueCapn)(unsafe.Pointer(C.PointerList(s).ToArray()))
}
func (s ClueCapn_List) Set(i int, item ClueCapn) { C.PointerList(s).Set(i, C.Object(item)) }

type RevCapn C.Struct

func NewRevCapn(s *C.Segment) RevCapn      { return RevCapn(s.NewStruct(8, 0)) }
func NewRootRevCapn(s *C.Segment) RevCapn  { return RevCapn(s.NewRootStruct(8, 0)) }
func AutoNewRevCapn(s *C.Segment) RevCapn  { return RevCapn(s.NewStructAR(8, 0)) }
func ReadRootRevCapn(s *C.Segment) RevCapn { return RevCapn(s.Root(0).ToStruct()) }
func (s RevCapn) DocId() uint32            { return C.Struct(s).Get32(0) }
func (s RevCapn) SetDocId(v uint32)        { C.Struct(s).Set32(0, v) }
func (s RevCapn) Rank() uint16             { return C.Struct(s).Get16(4) }
func (s RevCapn) SetRank(v uint16)         { C.Struct(s).Set16(4, v) }
func (s RevCapn) InMeta() uint16           { return C.Struct(s).Get16(6) }
func (s RevCapn) SetInMeta(v uint16)       { C.Struct(s).Set16(6, v) }

// capn.JSON_enabled == false so we stub MarshallJSON().
func (s RevCapn) MarshalJSON() (bs []byte, err error) { return }

type RevCapn_List C.PointerList

func NewRevCapnList(s *C.Segment, sz int) RevCapn_List {
	return RevCapn_List(s.NewCompositeList(8, 0, sz))
}
func (s RevCapn_List) Len() int         { return C.PointerList(s).Len() }
func (s RevCapn_List) At(i int) RevCapn { return RevCapn(C.PointerList(s).At(i).ToStruct()) }
func (s RevCapn_List) ToArray() []RevCapn {
	return *(*[]RevCapn)(unsafe.Pointer(C.PointerList(s).ToArray()))
}
func (s RevCapn_List) Set(i int, item RevCapn) { C.PointerList(s).Set(i, C.Object(item)) }

type RevsCapn C.Struct

func NewRevsCapn(s *C.Segment) RevsCapn      { return RevsCapn(s.NewStruct(0, 1)) }
func NewRootRevsCapn(s *C.Segment) RevsCapn  { return RevsCapn(s.NewRootStruct(0, 1)) }
func AutoNewRevsCapn(s *C.Segment) RevsCapn  { return RevsCapn(s.NewStructAR(0, 1)) }
func ReadRootRevsCapn(s *C.Segment) RevsCapn { return RevsCapn(s.Root(0).ToStruct()) }
func (s RevsCapn) Data() RevCapn_List        { return RevCapn_List(C.Struct(s).GetObject(0)) }
func (s RevsCapn) SetData(v RevCapn_List)    { C.Struct(s).SetObject(0, C.Object(v)) }

// capn.JSON_enabled == false so we stub MarshallJSON().
func (s RevsCapn) MarshalJSON() (bs []byte, err error) { return }

type RevsCapn_List C.PointerList

func NewRevsCapnList(s *C.Segment, sz int) RevsCapn_List {
	return RevsCapn_List(s.NewCompositeList(0, 1, sz))
}
func (s RevsCapn_List) Len() int          { return C.PointerList(s).Len() }
func (s RevsCapn_List) At(i int) RevsCapn { return RevsCapn(C.PointerList(s).At(i).ToStruct()) }
func (s RevsCapn_List) ToArray() []RevsCapn {
	return *(*[]RevsCapn)(unsafe.Pointer(C.PointerList(s).ToArray()))
}
func (s RevsCapn_List) Set(i int, item RevsCapn) { C.PointerList(s).Set(i, C.Object(item)) }

type ShotgunCapn C.Struct

func NewShotgunCapn(s *C.Segment) ShotgunCapn      { return ShotgunCapn(s.NewStruct(8, 1)) }
func NewRootShotgunCapn(s *C.Segment) ShotgunCapn  { return ShotgunCapn(s.NewRootStruct(8, 1)) }
func AutoNewShotgunCapn(s *C.Segment) ShotgunCapn  { return ShotgunCapn(s.NewStructAR(8, 1)) }
func ReadRootShotgunCapn(s *C.Segment) ShotgunCapn { return ShotgunCapn(s.Root(0).ToStruct()) }
func (s ShotgunCapn) Term() string                 { return C.Struct(s).GetObject(0).ToText() }
func (s ShotgunCapn) SetTerm(v string)             { C.Struct(s).SetObject(0, s.Segment.NewText(v)) }
func (s ShotgunCapn) Potency() float32             { return math.Float32frombits(C.Struct(s).Get32(0)) }
func (s ShotgunCapn) SetPotency(v float32)         { C.Struct(s).Set32(0, math.Float32bits(v)) }

// capn.JSON_enabled == false so we stub MarshallJSON().
func (s ShotgunCapn) MarshalJSON() (bs []byte, err error) { return }

type ShotgunCapn_List C.PointerList

func NewShotgunCapnList(s *C.Segment, sz int) ShotgunCapn_List {
	return ShotgunCapn_List(s.NewCompositeList(8, 1, sz))
}
func (s ShotgunCapn_List) Len() int             { return C.PointerList(s).Len() }
func (s ShotgunCapn_List) At(i int) ShotgunCapn { return ShotgunCapn(C.PointerList(s).At(i).ToStruct()) }
func (s ShotgunCapn_List) ToArray() []ShotgunCapn {
	return *(*[]ShotgunCapn)(unsafe.Pointer(C.PointerList(s).ToArray()))
}
func (s ShotgunCapn_List) Set(i int, item ShotgunCapn) { C.PointerList(s).Set(i, C.Object(item)) }

type TermCapn C.Struct

func NewTermCapn(s *C.Segment) TermCapn          { return TermCapn(s.NewStruct(24, 3)) }
func NewRootTermCapn(s *C.Segment) TermCapn      { return TermCapn(s.NewRootStruct(24, 3)) }
func AutoNewTermCapn(s *C.Segment) TermCapn      { return TermCapn(s.NewStructAR(24, 3)) }
func ReadRootTermCapn(s *C.Segment) TermCapn     { return TermCapn(s.Root(0).ToStruct()) }
func (s TermCapn) TermStr() string               { return C.Struct(s).GetObject(0).ToText() }
func (s TermCapn) SetTermStr(v string)           { C.Struct(s).SetObject(0, s.Segment.NewText(v)) }
func (s TermCapn) Slot() uint32                  { return C.Struct(s).Get32(0) }
func (s TermCapn) SetSlot(v uint32)              { C.Struct(s).Set32(0, v) }
func (s TermCapn) NumDocuments() uint32          { return C.Struct(s).Get32(4) }
func (s TermCapn) SetNumDocuments(v uint32)      { C.Struct(s).Set32(4, v) }
func (s TermCapn) NumWords() uint8               { return C.Struct(s).Get8(8) }
func (s TermCapn) SetNumWords(v uint8)           { C.Struct(s).Set8(8, v) }
func (s TermCapn) Shotgun() ShotgunCapn_List     { return ShotgunCapn_List(C.Struct(s).GetObject(1)) }
func (s TermCapn) SetShotgun(v ShotgunCapn_List) { C.Struct(s).SetObject(1, C.Object(v)) }
func (s TermCapn) Clues() ClueCapn_List          { return ClueCapn_List(C.Struct(s).GetObject(2)) }
func (s TermCapn) SetClues(v ClueCapn_List)      { C.Struct(s).SetObject(2, C.Object(v)) }
func (s TermCapn) InteractionsPos() uint16       { return C.Struct(s).Get16(10) }
func (s TermCapn) SetInteractionsPos(v uint16)   { C.Struct(s).Set16(10, v) }
func (s TermCapn) InteractionsNeg() uint16       { return C.Struct(s).Get16(12) }
func (s TermCapn) SetInteractionsNeg(v uint16)   { C.Struct(s).Set16(12, v) }
func (s TermCapn) HardcodedScore() int8          { return int8(C.Struct(s).Get8(9)) }
func (s TermCapn) SetHardcodedScore(v int8)      { C.Struct(s).Set8(9, uint8(v)) }
func (s TermCapn) Infogain() float32             { return math.Float32frombits(C.Struct(s).Get32(16)) }
func (s TermCapn) SetInfogain(v float32)         { C.Struct(s).Set32(16, math.Float32bits(v)) }

// capn.JSON_enabled == false so we stub MarshallJSON().
func (s TermCapn) MarshalJSON() (bs []byte, err error) { return }

type TermCapn_List C.PointerList

func NewTermCapnList(s *C.Segment, sz int) TermCapn_List {
	return TermCapn_List(s.NewCompositeList(24, 3, sz))
}
func (s TermCapn_List) Len() int          { return C.PointerList(s).Len() }
func (s TermCapn_List) At(i int) TermCapn { return TermCapn(C.PointerList(s).At(i).ToStruct()) }
func (s TermCapn_List) ToArray() []TermCapn {
	return *(*[]TermCapn)(unsafe.Pointer(C.PointerList(s).ToArray()))
}
func (s TermCapn_List) Set(i int, item TermCapn) { C.PointerList(s).Set(i, C.Object(item)) }
