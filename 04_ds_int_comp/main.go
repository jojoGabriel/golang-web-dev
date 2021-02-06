package main

import "fmt"

func slicer() {

	si := []int{3, 5, 7, 9}

	for i, _ := range si {
		fmt.Printf("slice index: %v\n", i)
	}

	for i, v := range si {
		fmt.Printf("slice index: %v value: %v\n", i, v)
	}
}

func mapper() {

	ids := map[string]int{"john": 1, "jane": 2}

	for k, v := range ids {
		fmt.Printf("key: %v ", k)
		fmt.Printf("val: %v\n", v)
		fmt.Printf("kv: %v\n", ids[k])
	}

}

type person struct {
	fname   string
	lname   string
	favFood []string
}

func personify() {

	p := person{fname: "jojo", favFood: []string{"meat", "veg"}}
	p.lname = "gabriel"

	fmt.Printf("p %v\n", p)
	fmt.Printf("lname %v\n", p.lname)
	fmt.Printf("food %v\n", p.favFood)

	for _, v := range p.favFood {
		fmt.Printf("%v ", v)
	}
	fmt.Println()

	fmt.Printf(p.walk())

}

func (p person) walk() string {
	return fmt.Sprintln(p.fname + " is walking")
}

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func (t truck) transportationDevice() string {
	return "deliver"
}

func (s sedan) transportationDevice() string {
	return "travel"
}

type transportation interface {
	transportationDevice() string
}

func report(t transportation) string {
	return t.transportationDevice()
}

func vehicular() {

	t := truck{vehicle{doors: 4, color: "blue"}, true}
	s := sedan{vehicle{doors: 4, color: "red"}, false}

	fmt.Printf("truck doors % v\n", t.vehicle.doors)
	fmt.Printf("sedan luxy %v\n", s.luxury)

	fmt.Println("truck " + t.transportationDevice())
	fmt.Println("sedan " + s.transportationDevice())

	fmt.Println("truck " + report(t))
	fmt.Println("sedan " + report(s))

}

type gator int

func (g gator) greeting() {
	fmt.Println("Hello, I am a gator")
}

type flamingo bool

func (f flamingo) greeting() {
	fmt.Println("Hello, I am pink and beautiful and wonderful")
}

type swampCreature interface {
	greeting()
}

func bayou(s swampCreature) {
	s.greeting()
}

func typer() {

	var g1 gator
	g1 = 12

	fmt.Printf("g1 %T %v\n", g1, g1)
	g1.greeting()

	var x int = 10
	fmt.Printf("x %T %v\n", x, x)

	x = int(g1)
	fmt.Printf("x %T %v\n", x, x)

	var f flamingo = true

	bayou(g1)
	bayou(f)

}

func byteMe() {
	s := "i'm sorry dave i can't do that"

	fmt.Println(s)

	fmt.Println([]byte(s))

	fmt.Println(string([]byte(s)))

	fmt.Println(s[:14])

	fmt.Println(s[10:22])

	for _, c := range []byte(s) {
		fmt.Println(string(c))
	}
}

func main() {
	// slicer()
	// mapper()
	// personify()
	// vehicular()
	// typer()
	byteMe()
}
