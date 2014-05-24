package main

import (
	"os"
	"errors"
	"github.com/HouzuoGuo/tiedot/gommap"
	"sync"
	"sync/atomic"
	"encoding/binary"
	"unsafe"
	"fmt"
	"bytes"
)

// START OMIT
type MMapStore struct {
	Path string           // Path to the mmap'd file
	DataVersion string    // The data version to check on loading
	Slots gommap.MMap     // The []byte that the file is mmap'd to
	SlotSize uint32       // The capacity in bytes of each slot
	GrowthSize uint32     // The number of slots to grow by once we've exhausted what we have
	Chained bool          // Whether or not the store supports chained slots
	File *os.File         // The file that is mmap'd
	NextSlot uint32       // The next fresh slot to consume
	DeletedSlots []uint32 // The slice of slots that have been deleted for reuse
	slotOverhead uint8    // The number of bytes overhead per slot
	storeOverhead uint32  // The number of bytes overhead at the beginning of the store
	sync.RWMutex
}
// END OMIT

// Creates a new store at the given path
func CreateMMapStore(path string, growthSize uint32, slotSize uint32, chained bool, dataVersion string) (*MMapStore, error) {
	var err error
	
	store := new(MMapStore)
	store.Path = path
	store.SlotSize = slotSize
	store.GrowthSize = growthSize
	store.NextSlot = 1
	store.DataVersion = dataVersion
	store.Chained = chained
	store.storeOverhead = 6+uint32(len(store.DataVersion))
	
	if store.Chained {
		store.slotOverhead = 6
	} else {
		store.slotOverhead = 2
	}

	if store.File, err = os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0600); err != nil {
		return nil, errors.New("Couldn't create file: "+err.Error())
	}
	slotBytes := int64(store.slotOverhead)+int64(store.SlotSize)
	futureSlots := int64(store.GrowthSize)
	if err = store.File.Truncate(int64(store.storeOverhead)+(slotBytes*futureSlots)); err != nil {
		return nil, errors.New("Couldn't truncate file: "+err.Error())
	}
	if store.Slots, err = gommap.Map(store.File, gommap.RDWR, 0); err != nil {
		return nil, errors.New("Couldn't map file: "+err.Error())
	}

	// Write header
	binary.LittleEndian.PutUint32(store.Slots[0:4], slotSize)
	if store.Chained {
		store.Slots[4] = 1
	} else {
		store.Slots[4] = 0
	}
	store.Slots[5] = uint8(len(store.DataVersion))
	dataVersionBytes := []byte(store.DataVersion)
	for i := 0; i < len(dataVersionBytes); i++ {
		store.Slots[6+i] = dataVersionBytes[i]
	}
	
	go store.FindDeletedSlots()
	
	return store, nil
}

// Loads an existing store at the given path
func LoadMMapStore(path string, dataVersion string, growthSize uint32) (*MMapStore, error) {
	var err error
	store := new(MMapStore)
	store.Path = path
	store.GrowthSize = growthSize

// START2 OMIT
	if store.File, err = os.OpenFile(path, os.O_RDWR, 0600); err != nil {
		return nil, errors.New("Couldn't read file: "+err.Error())
	}
	if store.Slots, err = gommap.Map(store.File, gommap.RDWR, 0); err != nil {
		return nil, errors.New("Couldn't map file: "+err.Error())
	}

// END2 OMIT	// OMIT
