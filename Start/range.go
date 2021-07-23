package main

import "fmt"

func rangeSample() {

	nums := []int{1, 2, 3, 4}
	sum := 0
	for _, v := range nums {
		sum += v
	}

	fmt.Println("Sum of numbers: ", sum)

	for i, v := range nums {
		if i == 3 {
			fmt.Println("value at selected index: ", v)
		}
	}

	kvs := map[string]string{"a": "banana", "b": "apple"}
	for key, value := range kvs {
		fmt.Printf("%s => %s\n", key, value)
	}
}
