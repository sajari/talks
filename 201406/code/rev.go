package main

import (
	"encoding/binary"
	"fmt"
)
// START OMIT
type RevStore struct {
	Store *MMapStore
}

func NewRevStore(path string) (*RevStore, error) {
	rev := new(RevStore)
	var err error
	if rev.Store, err = CreateMMapStore(path, 1024*1024, 12, false, "sjengine-rev-9.6.0"); err != nil {
		return nil, fmt.Errorf("Reverse index %s: %s", path, err.Error())
	}
	return rev, nil
}

func (rev *RevStore) Add(firstSlot uint32, documentId uint32, rank uint16, inMeta uint16) (uint32, error) {
	data := make([]byte, 12)
	binary.LittleEndian.PutUint32(data[0:4], documentId)
	binary.LittleEndian.PutUint16(data[4:6], rank)
	binary.LittleEndian.PutUint16(data[6:8], inMeta)
	binary.LittleEndian.PutUint32(data[8:12], firstSlot)
	return rev.Store.Add(data)
}
// END OMIT

func (rev *RevStore) Remove(slot uint32, removeDocumentId uint32) uint32 {
	originalSlot := slot
	documentId, _, _, nextSlot := rev.Get(slot)
	if documentId == removeDocumentId {
		rev.Store.ClearSlot(slot)
		return nextSlot
	} else {
		for {
			documentId, _, _, nextNextSlot := rev.Get(nextSlot)
			if documentId == removeDocumentId {
				data, _ := rev.Store.Get(slot)
				binary.LittleEndian.PutUint32(data[8:12], nextNextSlot)
				rev.Store.WriteSlot(slot, data)
				rev.Store.ClearSlot(nextSlot)
				return originalSlot
			}
			if nextSlot == 0 {
				return originalSlot
			} else {
				slot = nextSlot
				nextSlot = nextNextSlot
			}
		}
	}
}

// Returns the documentId, rank, inMeta, and the next slot
func (rev *RevStore) Get(slot uint32) (uint32, uint16, uint16, uint32) {
	//data := *(*[12]byte)(rev.Store.GetPointer(slot))
	data, _ := rev.Store.Get(slot)
	return binary.LittleEndian.Uint32(data[0:4]),
	       binary.LittleEndian.Uint16(data[4:6]),
	       binary.LittleEndian.Uint16(data[6:8]),
	       binary.LittleEndian.Uint32(data[8:12])
}

