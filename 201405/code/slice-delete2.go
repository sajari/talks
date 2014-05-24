package main

import (
	"time"
	"fmt"
)

type Rev struct {
	id 		uint32
	woff 	uint32
	poff 	uint32
	rank 	uint32
	inMeta 	uint32
}

// START OMIT
func RemoveElemWithCopy(r []int, i int) {
	newslice := make([]int, len(r)-1)
	copy(newslice, r[0:i])
	copy(newslice[i:], r[i+1:])
	r = newslice
	fmt.Println("Copy: ", r) //OMIT
}

func RemoveElemWithCopy2(r []int, i int) {
	r = r[:i+copy(r[i:], r[i+1:])]
	fmt.Println("Copy 2: ", r) //OMIT
}

func RemoveElemWithShift(r []int, i int) {
	r[i], r = r[len(r)-1], r[:len(r)-1]
	fmt.Println("Shift: ", r) //OMIT
}

func RemoveElemWithAppend(r []int, i int) {
	r = append(r[:i], r[i+1:]...)
	fmt.Println("Append: ", r) //OMIT
}
// END OMIT




func main() {

	testmax := 10
	r := make([]int, testmax)
	for i := 0; i < testmax; i++ {
		r[i] = i
	}
	t0 := time.Now()
	RemoveElemWithAppend(r, 4)
	t1 := time.Now()
	fmt.Println(t1.Sub(t0))
	

}