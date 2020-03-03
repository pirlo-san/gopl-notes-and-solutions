package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan struct{})
	count := 0
	begin := time.Now()
	defer func(begin time.Time) {
		end := time.Now()
		fmt.Printf("\nDone after %v, with %v goroutines created\n",
			end.Sub(begin), count)
		if r := recover(); r != nil {
			begin := time.Now()
			fmt.Printf("Sending data to channels\n")
			for i := 0; i < count; i++ {
				ch <- struct{}{}
			}
			end := time.Now()
			fmt.Printf("Done after %v", end.Sub(begin))
		}
	}(begin)

	fmt.Println("Number of goroutines:")
	for ; count < 1000000; count++ {
		fmt.Printf("\r%d", count)
		go func() {
			// all routine will block at reading from ch
			// and finally run out of memory, which will lead to a panic
			<-ch
		}()
	}
	time.Sleep(time.Second)
	panic("pipeline\n")
}
