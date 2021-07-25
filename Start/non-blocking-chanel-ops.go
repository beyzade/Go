package main

import "fmt"

func nonBlockingChannelOperations() {
	messages := make(chan string)
	signals := make(chan string)

	select {
	case msg1 := <-messages:
		fmt.Println("Message received from messages", msg1)
	default:
		fmt.Println("No messages received from messages")
	}

	msg2 := "Hi"
	select {
	case signals <- msg2:
		fmt.Println("Message sent to signals")
	default:
		fmt.Println("No messages sent to signals")
	}

	select {
	case msg3 := <-messages:
		fmt.Println("Message received from messages", msg3)
	case msg4 := <-signals:
		fmt.Println("Message received from signals", msg4)
	default:
		fmt.Println("No activity (multiple way non blocking)")
	}

}
