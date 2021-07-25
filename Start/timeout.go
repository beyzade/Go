package main

import (
	"fmt"
	"time"
)

func timeoutSample() {

	ch1 := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "Message 1"
	}()

	select {
	case res1 := <-ch1:
		fmt.Println("res1:", res1)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	ch2 := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Message 2"
	}()

	select {
	case res2 := <-ch2:
		fmt.Println("res2: ", res2)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}

}
