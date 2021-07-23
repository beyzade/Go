package main

import "fmt"

func pointerSample() {

	val := 1
	fmt.Println("Initial value: ", val)
	increase(val)
	fmt.Println("After increase: ", val)
	increasePtr(&val)
	fmt.Println("After increasePtr: ", val)
	fmt.Println("Memory address: ", &val)

}

func increase(val int) {
	val += 1
}

func increasePtr(val *int) {
	*val += 1
}
