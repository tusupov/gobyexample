package main

import "fmt"

type person struct {
	name string
	age int
}

func main() {

	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "Alica", age: 30})

	fmt.Println(person{name: "Fed"})

	fmt.Println(&person{name: "Ann", age: 40})

	s := person{name: "Sean", age: 45}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 50
	fmt.Println(s.age)

}
