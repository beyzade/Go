package main

import "fmt"

type channelMessage struct {
	code       int
	definition string
}

func sendMessage() {
	ch := make(chan channelMessage)
	cms := channelMessage{code: 1, definition: "Errro code 1"}
	go func() { ch <- cms }()

	cmr := <-ch
	fmt.Println("cmr", cmr)
}

func channelSample() {
	sendMessage()

}
