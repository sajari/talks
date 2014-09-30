package main
	
import (
	"fmt"
	"sync"
	"reflect"
	"errors"
)

// START OMIT
// Basic structure of a data bucket. Next and Prev are optional
type Bucket struct {
	Datatype	string				// Description only
	Length		uint32				// Max number of objects 4,294,967,295
	Offset		chan uint32 		// Write offset. If Offset=Length, auto extends to 1.5x
	Data		interface{}			// Any struct as the data type, cannot change
	StringMap	map[string]uint32 	// Optional string to slot lookup
	Uint32Map	map[uint32]uint32 	// Optional uint32 to slot lookup
	DeletedSlots	[]uint32		// Reuses slots when they are deleted
	UsedCount	uint32				// Object count
	sync.RWMutex
}
// END OMIT

// Create a new bucket of type datastruct
// START2 OMIT
func NewBucket(datatype string, hashmap string, length uint32, datastruct interface{}) *Bucket {
	b := new(Bucket)
	b.Datatype = datatype
	b.Length = length
	b.Offset = make(chan uint32, 1)
    b.Offset <- 1
	b.DeletedSlots = make([]uint32, 0, 100) 
	myType := reflect.TypeOf(datastruct)
    b.Data = reflect.MakeSlice(reflect.SliceOf(myType), int(length), int(length)).Interface()
    
    switch hashmap { 
    	case "string":
        	b.StringMap = make(map[string]uint32, length)
    	case "uint32":
        	b.Uint32Map = make(map[uint32]uint32, length)
    }
	return b
}
// END2 OMIT

// Automatically grow the size of a bucket as required
// START3 OMIT
func (bucket *Bucket) ExtendBucket() {		
	bucket.Lock()
	oldlength := bucket.Length
	newlength := uint32(float32(oldlength) * 1.5)
	extension := newlength - oldlength
	bucket.Length = newlength

	// Extend the datablock
	myType := reflect.ValueOf(bucket.Data).Index(0).Type()
	newslice := reflect.MakeSlice(reflect.SliceOf(myType), int(extension), int(extension)).Interface()
	bucket.Data = reflect.AppendSlice(reflect.ValueOf(bucket.Data), reflect.ValueOf(newslice)).Interface()

	fmt.Printf("Bucket (%v) capacity extended from %v to %v objects...\n", bucket.Datatype, oldlength, newlength)
    bucket.Unlock()
}
// END3 OMIT

// Allocates a slot for writing, either new or from an old deleted slot
func (bucket *Bucket) GetSlot() uint32 {
    var slot uint32
    bucket.Lock()
	if len(bucket.DeletedSlots) > 0 {
		slot = bucket.DeletedSlots[0]
		bucket.DeletedSlots[0], bucket.DeletedSlots = bucket.DeletedSlots[len(bucket.DeletedSlots)-1], bucket.DeletedSlots[:len(bucket.DeletedSlots)-1]
		bucket.Unlock()
	} else {
		slot = <- bucket.Offset
		bucket.UsedCount++
		bucket.Unlock()
		if slot == bucket.Length {
	        bucket.ExtendBucket()
	    } 
	    bucket.Offset <- slot+1
	}
    return slot
}


// Add an object to a bucket
// If there is a deleted slot, use it, otherwise write to the offset
func (bucket *Bucket) Add(object interface{}, next uint32, prev uint32) (uint32, error) {
	slot := bucket.GetSlot()

    bucket.RLock()
	defer bucket.RUnlock()

	bD := reflect.ValueOf(bucket.Data).Index(int(slot))
	val := reflect.ValueOf(object)
	bD.Set(val) // TODO: Should check the data type is ok before adding...

	return slot, nil
}

// Add an object using a hashmap lookup value 
func (bucket *Bucket) AddMapped(hashmap interface{}, object interface{}) (uint32, error) {
	slot := bucket.GetSlot()

	bucket.Lock()
	switch hashmap.(type) {
        case string:
     		bucket.StringMap[hashmap.(string)] = slot        
        case uint32:
            bucket.Uint32Map[hashmap.(uint32)] = slot
        default:
        	return 0, errors.New("Incorrect hashmap type used in Add() function...")   
    }

	bD := reflect.ValueOf(bucket.Data).Index(int(slot))
	val := reflect.ValueOf(object)
	bD.Set(val) // TODO: Should check the data type is ok before adding...

	bucket.Unlock()

	return slot, nil
}

// Remove an object from a bucket and correct the linkedin list linking structure if necessary
func (bucket *Bucket) Delete(hashmap interface{}) error {
	var slot uint32
	var ok bool
	var err error
	switch hashmap.(type) {
        case string:
     		slot, ok = bucket.StringMap[hashmap.(string)]; if ok {
     			bucket.Lock()
     			delete(bucket.StringMap, hashmap.(string)) 
     			err = bucket.deleteSlot(slot)
     			bucket.Unlock()
 			} else {
 				return errors.New("That hashmap does not exist...") 
 			}
     		     
        case uint32:
            slot, ok = bucket.Uint32Map[hashmap.(uint32)]; if ok {
            	bucket.Lock()
            	delete(bucket.Uint32Map, hashmap.(uint32)) 
            	err = bucket.deleteSlot(slot)
            	bucket.Unlock()
            } else {
 				return errors.New("That hashmap does not exist...") 
 			}

        default:
        	fmt.Println("incorrect hash function...")
    }
    
    return err
}

func (bucket *Bucket) deleteSlot(slot uint32) error {
	if slot > bucket.UsedCount {
		return errors.New("Slot out of data range")
	}
	for _, delslot := range bucket.DeletedSlots {
		if delslot == slot {
			// Already been deleted and queued for reuse...
			return nil
		}
	}
	bucket.DeletedSlots = append(bucket.DeletedSlots, slot)
	return nil
}

// Return the pointer to an object
func (bucket *Bucket) Get(slot uint32) (interface{}, bool) {
	bucket.RLock()
	defer bucket.RUnlock()
	if slot > bucket.UsedCount {
		return 0, false
	}
	read := reflect.ValueOf(bucket.Data).Index(int(slot)).Interface()
    if reflect.ValueOf(read).IsNil() {
        return 0, false
    }
	return read, true
}

// Return the pointer to an object using it's map lookup
func (bucket *Bucket) Map(input interface{}) (interface{}, bool) {
	bucket.RLock()
	var slot uint32
	switch input.(type) {
        case string:
     		slot = bucket.StringMap[input.(string)]           
        case uint32:
            slot = bucket.Uint32Map[input.(uint32)]
        default:
        	return nil, false    
    }
	bucket.RUnlock()
	return bucket.Get(slot)
}

// Gets the number of objects in the bucket currently
func (bucket *Bucket) Count() uint32 {
    bucket.RLock()
    defer bucket.RUnlock()
    return bucket.UsedCount - uint32(len(bucket.DeletedSlots))
}

// Stream all data in the bucket's Data array
func (bucket *Bucket) StreamAll() <- chan interface{} {
	bucket.RLock()
	defer bucket.RUnlock()
	out := make(chan interface{}, 10)
	go func() {
		for i := 1; i <= int(bucket.UsedCount); i++ {
			read := reflect.ValueOf(bucket.Data).Index(int(i)).Interface()
		    if reflect.ValueOf(read).IsNil() {
		    	close(out)
		    	return
		    }
		    out <- read
		}
		close(out) 
	}()
	return out
} 

// Dump a bucket for debug purposes
func (bucket *Bucket) Dump() {
	bucket.Lock()
	defer bucket.Unlock()
	fmt.Println(bucket)
}