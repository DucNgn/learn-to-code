package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func main() {
	p := person{name: "Duke", age: 21}
	fmt.Println(p.age)
}
