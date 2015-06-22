package main

import (
	"encoding/binary"
	"github.com/gogo/protobuf/proto"
	"testing"
)

func OldRead(data []byte) (uint32, uint16, uint16, uint32) {
	return binary.LittleEndian.Uint32(data[0:4]),
		binary.LittleEndian.Uint16(data[4:6]),
		binary.LittleEndian.Uint16(data[6:8]),
		binary.LittleEndian.Uint32(data[8:12])
}

func ProtoRead(data []byte) *Rev {
	rev := new(Rev)
	proto.Unmarshal(data, rev)
	return rev
}

func BenchmarkOld(b *testing.B) {
	data := []byte("123412121234")
	// run the function b.N times
	for n := 0; n < b.N; n++ {
		OldRead(data)
	}
}

func BenchmarkNew(b *testing.B) {
	data := []byte("123412121234")
	// run the function b.N times
	for n := 0; n < b.N; n++ {
		ProtoRead(data)
	}
}
