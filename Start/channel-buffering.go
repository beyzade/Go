package main

import "fmt"

func channelBufferingSammple() {

	messages := make(chan string, 2)
	messages <- "Message 1"
	fmt.Printf("Len: %d\n", len(messages))
	fmt.Printf("Capacity: %d\n", cap(messages))
	messages <- "Message 2"
	fmt.Printf("Len: %d\n", len(messages))
	fmt.Printf("Capacity: %d\n", cap(messages))
	fmt.Println("Element count: ", len(messages))

	fmt.Println(<-messages)
	fmt.Println(<-messages)

}
