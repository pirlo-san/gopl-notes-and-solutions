package main

import (
	"fmt"
	"time"
)

func main() {
	msg1, msg2 := make(chan string), make(chan string)

	go func() {
		for {
			fmt.Println(<-msg1)
			msg2 <- "message from new goroutine"
		}
	}()

	count := 0
	t := time.NewTicker(time.Second)
	defer t.Stop()
	for {
		select {
			case <-t.C:
				fmt.Printf("%v communications per second\n", count)
				count = 0
			case msg1 <- "message from main":
			case msg := <-msg2:
				fmt.Println(msg)
				count++
		}
	}
}
