package main

import "fmt"

type rect struct {
	height, width int
}

func (r rect) area() int {
	return r.height * r.height
}

func (r *rect) perim() int {
	return 2*r.height + 2*r.width
}

func methodSample() {
	r := rect{height: 10, width: 20}
	fmt.Println("r: ", r)
	fmt.Println("area: ", r.area())
	fmt.Println("perim: ", r.perim())

	rp := &r
	fmt.Println("rp area: ", rp.area())
	fmt.Println("rp perim: ", rp.perim())

}
