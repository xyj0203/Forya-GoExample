package main

import "context"

//// | 运算会把 0 变为 1 然后变不回去
//func minimumSubarrayLength(nums []int, k int) int {
//
//}

func main() {
	for i := 0; i < 3; i++ {
		defer func() { println(i) }()
	}
	recover()
	context.WithCancel()
}
