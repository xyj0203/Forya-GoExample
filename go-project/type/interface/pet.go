package main

import "fmt"

type Pet interface {
	Name() string
	Category() string
	setName(string)
}

type Dog struct {
	name     string
	category string
}

func (d Dog) Name() string {
	return d.name
}

func (d Dog) setName(name string) {
	d.name = name
}

func (d Dog) Category() string {
	return d.category
}

func main() {
	main2()
}

/**
 * 测试不同类型饮用
 */
func main1() {
	var pet Pet
	pet = Dog{
		name:     "dog",
		category: "mammal",
	}
	fmt.Println("指针类型====================================>")
	fmt.Println(pet.Name())
	fmt.Println(pet.Name())
	fmt.Println(pet.Category())
	fmt.Println("值类型====================================>")
	var dog Dog
	fmt.Println(dog.Name())
	fmt.Println(dog.Category())
}

/**
 * 测试值类型调用，并没有更改，
 */
func main2() {
	// 值类型修改复制
	originDog := Dog{
		name:     "dog",
		category: "mammal",
	}
	fmt.Println("值类型修改复制====================================>")
	modifyDog := originDog
	modifyDog.setName("dog2")
	fmt.Println("原始值： " + originDog.Name())
	fmt.Println("修改值： " + modifyDog.Name())
}
