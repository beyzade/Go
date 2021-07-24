package main

import (
	"fmt"
	"time"
)

func selectSample() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func(message string) {
		time.Sleep(1 * time.Second)
		ch1 <- message
	}("one")

	go func(message string) {
		time.Sleep(2 * time.Second)
		ch2 <- message
	}("two")

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("From channel 1 message: ", msg1)
		case msg2 := <-ch2:
			fmt.Println("From channel 2 message: ", msg2)
		}
	}
}
