package main

import (
	"fmt"
	"math"
)

type rect struct {
	height, width float64
}

type circle struct {
	radius float64
}

type geometry interface {
	area() float64
	perim() float64
}

func (r rect) area() float64 {
	return r.height * r.width
}

func (r rect) perim() float64 {
	return 2*r.height + 2*r.width
}

func (c circle) area() float64 {
	return 2 * math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func typeAssertion() geometry {
	return &rect{height: 15, width: 25}
}

func interfaceSample() {
	r := rect{height: 10, width: 20}
	c := circle{radius: 10}
	measure(r)
	measure(c)

	g := typeAssertion()
	rta, ok := g.(*rect)
	fmt.Println("ok: ", ok)
	fmt.Println("rta: ", rta)

	fmt.Println("rta.height: ", rta.height)
	fmt.Println("rta.width: ", rta.width)
	fmt.Printf("%+v", rta)

}
