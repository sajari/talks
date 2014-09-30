package main

import (
	"fmt"
	"sync"
	"runtime"
	"time"
)

func main () {
		numcpu := runtime.NumCPU()
		runtime.GOMAXPROCS(numcpu)
		queue := make(chan int, 50)

		// Create collector/s first
		var waitcollect sync.WaitGroup
		waitcollect.Add(1)
		go func() {
			for {
				_, ok := <- queue; if ok {
					fmt.Println("chan received...")
				} else {
					fmt.Printf("queue is BOTH drained and closed")
					waitcollect.Done()
					break
				}
			}
		}()
	
		// Now create processors to send to the collector
		var wait sync.WaitGroup
		for i := 0; i < numcpu; i++ {
			wait.Add(1)
			go func() {
				for i := 0; i < 10; i++ {
					queue <- i
				}
				time.Sleep(10 * time.Microsecond)
				wait.Done()
			}()
		}

		wait.Wait() // The queue sends have finished
		close(queue)

		fmt.Println("\n queue closed")
		waitcollect.Wait() // The collector will exit when it's finished draining the closed queue

		fmt.Println("\n finished collecting and processing the queue")
}