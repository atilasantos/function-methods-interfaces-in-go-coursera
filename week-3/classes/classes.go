package main

import "fmt"

type Person struct {
	name string
	age  int32
}

func (p Person) getName() string {
	return p.name
}

func main() {

	p := Person{name: "Átila", age: 31}
	fmt.Println(p.getName())

}
