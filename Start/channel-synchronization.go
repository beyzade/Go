package main

import (
	"fmt"
	"time"
)

func worker(ch chan bool) {
	fmt.Println("Started...")
	time.Sleep(time.Second)
	fmt.Println("Done ...")
	ch <- true
}

func channelSynchronizationSample() {
	ch := make(chan bool, 1)
	go worker(ch)
	fmt.Println(<-ch)
}
