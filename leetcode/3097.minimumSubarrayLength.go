package main

import (
	"fmt"
	"math"
	"strconv"
)

func minimumSubarrayLength(nums []int, k int) int {
	n := len(nums)
	var bits [30]int

	var ans = math.MaxInt32
	// 移动右数组， 当满足条件时移动左数组
	for left, right := 0, 0; right < n; right++ {
		for i := 0; i < 30; i++ {
			bits[i] += nums[right] >> i & 1
		}
		for ; left <= right && cal(bits) >= k; left++ {
			length := right - left + 1
			ans = min(ans, length)
			// 恢复操作
			for i := 0; i < 30; i++ {
				bits[i] -= nums[left] >> i & 1
			}
		}
	}
	if ans == math.MaxInt32 {
		return -1
	}
	return ans
}

func main() {
	var nums = []int{2, 1, 8}
	fmt.Printf(strconv.Itoa(minimumSubarrayLength(nums, 10)))
}

// 计算对应的二进制数组中的值
func cal(nums [30]int) int {
	ans := 0
	for i, v := range nums {
		if v != 0 {
			ans |= 1 << i
		}
	}
	return ans
}
