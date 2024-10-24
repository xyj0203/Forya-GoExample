// Package short 短变量
package main

import "fmt"

/**
 * 注意不同层级的值覆盖问题
 */
func main() {
	sameLevel()
	fmt.Printf("==============================\n")
	diffLevel()
}

// 同一层级
func sameLevel() {
	value, err := getResult(1)
	out, err := getResult(2)
	// 预期打印结果为2
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("%v\n", value)
	fmt.Printf("%v\n", out)
}

// 不同层级
func diffLevel() {
	value, err := getResult(1)
	// 预期打印结果为2
	if err != nil {
		out, err := getResult(2)
		fmt.Printf("%v\n", err)
		fmt.Printf("%v\n", out)
	}
	fmt.Printf("%v\n", value)
	// 此处的err还是外面的error
	fmt.Printf("%v\n", err)
}

func getResult(value int) (result int, err error) {
	return value, fmt.Errorf("error code: %d\n", value)
}
