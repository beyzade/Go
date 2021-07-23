package main

import "fmt"

func sums(nums ...int) {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	fmt.Println("Sum is: ", sum)
}

func sumsWithPrefix(prefix string, nums ...int) {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	fmt.Println(prefix, sum)
}

func functionSample() {
	sums(1, 2, 3)
	sums(10, 20, 30, 40)

	nums := []int{11, 12, 13, 14}
	sums(nums...)

	sumsWithPrefix("Toplam sayı: ", 11, 12, 13, 14, 15)

}
