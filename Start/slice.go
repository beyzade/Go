package main

import "fmt"

func sliceSample() {

	s := make([]string, 3)
	fmt.Println(s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set: ", s)
	fmt.Println("get: ", s[2])

	fmt.Println("len: ", len(s))

	s = append(s, "d", "e")
	fmt.Println(s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("c: ", c)

}
