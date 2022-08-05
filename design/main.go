package main

import "fmt"
import "time"

type Action interface {
	Call() string
	GetName() string
}

type Animal struct {
	name string
}

func NewAnimal() *Animal {
	return &Animal{name: "cjp"}
}

func (a Animal) Call() string {
	return "Animal call"
}

func (a Animal) GetName() string {
	return a.name
}

type Dog struct {
	*Animal
	name string
}

func (d Dog) GetName() string {
	return d.name
}

func main() {
	dog := Dog{Animal: NewAnimal(), name: "child"}
	dog.Animal.name = dog.name
	fmt.Println(dog.GetName())
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().UnixNano())
}
