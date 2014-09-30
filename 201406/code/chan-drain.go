package main

import "fmt"

// START OMIT
func main() {
	ch := make(chan int, 10)
	for i:=0; i<5; i++ {
		ch <- i
	} 
	fmt.Println("channel primed")
	close(ch)
	fmt.Println("channel closed")
	for {
		if v, ok := <- ch; ok {
			fmt.Println(v, ok)			
		} else {
			fmt.Println(v, ok)
			break
		}
	}	
	fmt.Println("channel drained")
}
// END OMIT