package main

import "fmt"

type Pet interface {
	SetName(name string)
	Name() string
	Category() string
}

type Dog struct {
	name     string
	category string
}

func (d Dog) Name() string {
	return d.name
}

func (d *Dog) SetName(name string) {
	d.name = name
}

func (d Dog) Category() string {
	return d.category
}

func main() {
	var pet Pet
	pet = &Dog{
		name:     "dog",
		category: "mammal",
	}
	fmt.Println("指针类型====================================>")
	fmt.Println(pet.Name())
	pet.SetName("dog2")
	fmt.Println(pet.Name())
	fmt.Println(pet.Category())
	fmt.Println("值类型====================================>")
	var dog Dog
	dog.SetName("dog3")
	fmt.Println(dog.Name())
	fmt.Println(dog.Category())
}
