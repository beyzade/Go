package main

import "fmt"

func channelDirectionsSample() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "Sample message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

func ping(pings chan<- string, message string) {
	pings <- message
}

func pong(pings <-chan string, pongs chan<- string) {
	message := <-pings
	pongs <- message
}
