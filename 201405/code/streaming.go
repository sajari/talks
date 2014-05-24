
package main

import (
	"fmt"
	"time"
	)
// START OMIT
func Stream(testsize int) <- chan int {
	out := make(chan int, 10)
	go func() {
		for i := 0; i < testsize; i++ {
			out <- i
		}
		close(out) 
	}()
	return out
} 

func StatePass(testsize int) []int {
	out := make([]int, testsize)
	for i := 0; i < testsize; i++ {
		out[i] = i
	}
	return out
}
// END OMIT
// START2 OMIT
func main() {
	testsize := 1000000
	t0 := time.Now()
	StatePass(testsize)
	t1 := time.Now()
	fmt.Println("StatePass: ", t1.Sub(t0))
	
	t2 := time.Now()
	stream := Stream(testsize)
	for {
		if _, ok := <- stream; !ok {
			break
		}
	}
	t3 := time.Now()	
	fmt.Println("Stream: ", t3.Sub(t2))
}
// END2 OMIT
