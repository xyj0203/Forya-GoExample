package main

import "fmt"

var container = []string{"h", "e", "l", "l", "o"}

func main() {
	container := map[int]string{
		1: "hello",
		2: "world",
	}
	// 类型断言
	value, ok := interface{}(container).(map[int]string)
	if ok {
		fmt.Printf("%v\n", value)
	}
}
