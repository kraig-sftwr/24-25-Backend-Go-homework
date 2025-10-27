package main

import (
	"fmt"
)

type Person struct {
	name        string
	age         int
	nationality string
}

func (per *Person) GetName() string {
	return per.name
}

func NewPerson(name string, age int, nationality string) *Person {
	return &Person{
		name:        name,
		age:         age,
		nationality: nationality,
	}
}

func main() {
	per := NewPerson("John Doe", 19, "USA")
	fmt.Println(per)
}
