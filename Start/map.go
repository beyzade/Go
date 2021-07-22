package main

import "fmt"

func mapSample() {

	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	m["three"] = 3
	fmt.Println(m)
	fmt.Println("len: ", len(m))

	v1 := m["two"]
	fmt.Println("v1: ", v1)

	delete(m, "three")
	fmt.Println(m)
	fmt.Println("len: ", len(m))

	v2, d2 := m["two"]
	fmt.Println("v2: ", v2, ", d2: ", d2)

	v4, d4 := m["four"]
	fmt.Println("v4: ", v4, ", d4: ", d4)

	m2 := map[string]int{"foo": 1, "bar": 2}
	fmt.Println(m2)

}
