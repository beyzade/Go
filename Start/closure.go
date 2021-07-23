package main

import "fmt"

func closureSample() {

	nextIntFunc := intSeq()

	fmt.Println("Current value: ", nextIntFunc())
	fmt.Println("Current value: ", nextIntFunc())
	fmt.Println("Current value: ", nextIntFunc())
	fmt.Println("Current value: ", nextIntFunc())

}

func intSeq() func() int {
	i := 0

	return func() int {
		i++
		return i
	}
}
