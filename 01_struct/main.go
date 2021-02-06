package main

import "fmt"

type person struct {
	name string
	age  int
}

type agent struct {
	person
	secret bool
}

func (p person) speak() {
	fmt.Printf("Person %v speaks\n", p)
}

func (a agent) speak() {
	fmt.Printf("Agent %v speaks\n", a)
}

type human interface {
	speak()
}

func talk(h human) {
	h.speak()
}

func main() {
	p := person{"Pete", 20}

	a := agent{p, true}

	fmt.Println(p.name)
	p.speak()

	fmt.Println(a.secret)
	a.speak()

	a.person.speak()

	talk(p)
	talk(a)
}
