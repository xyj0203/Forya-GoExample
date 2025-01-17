package main

import "fmt"

/*
*
给你一个下标从 0 开始长度为 n 的整数数组 nums 。
如果以下描述为真，那么 nums 在下标 i 处有一个 合法的分割 ：

前 i + 1 个元素的和 大于等于 剩下的 n - i - 1 个元素的和。
下标 i 的右边 至少有一个 元素，也就是说下标 i 满足 0 <= i < n - 1 。
请你返回 nums 中的 合法分割 方案数。
*/

func main() {
	nums := []int{2, 3, 1, 0}
	fmt.Println(waysToSplitArray(nums))
}

func waysToSplitArray(nums []int) int {
	count := 0
	length := len(nums)
	sum := make([]int64, length)
	sum[0] = int64(nums[0])
	for i := 1; i < length; i++ {
		sum[i] = sum[i-1] + int64(nums[i])
	}
	for i := 0; i < length-1; i++ {
		if sum[i] >= (sum[length-1] - sum[i]) {
			count++
		}
	}
	return count
}
