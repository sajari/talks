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

type Term struct {
	Name		string
	Rev 		[]*Rev
}

// START OMIT
func (term *Term) RemoveElemWithCopy(i int) {
	newslice := make([]*Rev, len(term.Rev)-1)
	copy(newslice, term.Rev[0:i])
	copy(newslice[i:], term.Rev[i+1:])
	term.Rev = newslice
}
func (term *Term) RemoveElemWithCopy2(i int) {
	term.Rev = term.Rev[:i+copy(term.Rev[i:], term.Rev[i+1:])]
}
func (term *Term) RemoveElemWithShift(i int) {
	term.Rev[i], term.Rev = term.Rev[len(term.Rev)-1], term.Rev[:len(term.Rev)-1]
}
func (term *Term) RemoveElemWithAppend(i int) {
	term.Rev = append(term.Rev[:i], term.Rev[i+1:]...)
}
// END OMIT




func main() {

	testmax := 1000000

	term := new(Term)
	term.Rev = make([]*Rev, testmax)

	t0 := time.Now()

	for i := 0; i < 10000; i++ {
		term.RemoveElemWithCopy2(i)
	}
	
	t1 := time.Now()

	fmt.Println(t1.Sub(t0))

}