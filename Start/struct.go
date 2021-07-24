package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 39

	return &p
}

func structSample() {
	fmt.Println(person{"Kerem", 4})
	fmt.Println(person{name: "Eren", age: 1})
	fmt.Println(person{name: "Yavuz"})
	p := person{name: "Hakan", age: 12}
	fmt.Println(&p)
	fmt.Println(&person{name: "Hakan", age: 12})

	fmt.Println(newPerson("Sultan"))
	s := newPerson("Seher")
	fmt.Println("s.age: ", s.age)
	fmt.Println("*s.age: ", (*s).age)
	sp := &(*s)
	sp.age = 35
	fmt.Println(s)
	fmt.Println(sp)

}
