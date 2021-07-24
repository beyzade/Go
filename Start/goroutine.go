package main

import (
	"fmt"
	"time"
)

func f(message string) {
	for i := 0; i < 3; i++ {
		fmt.Println(message, ": ", i)
	}
}

func goroutineSample() {
	f("Hi there")
	go f("goroutine")

	go func(message string) {
		// for i := 0; i < 3; i++ {
		// 	fmt.Println(message, ": ", i)
		// }
		fmt.Println(message)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("Done!")
}
