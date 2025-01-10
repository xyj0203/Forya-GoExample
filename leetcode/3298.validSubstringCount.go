package main

import (
	"fmt"
)

func main() {
	fmt.Println(validSubstringCount("bcca", "a"))
}

func validSubstringCount(word1 string, word2 string) int64 {
	// 计算word2 中各字符的个数
	template := make(map[rune]int)
	for _, v := range word2 {
		template[v]++
	}
	tempCount := make(map[rune]int)
	var count int64
	count = 0
	j := 0
	i := 0
	for i = 0; i < len(word1); {
		for j = i; j < len(word1); {
			if validCheck(tempCount, template) {
				count = count + 1 + int64(len(word1)-j)
				tempCount[rune(word1[i])]--
				i++
			} else {
				if j >= len(word1) {
					break
				}
				tempCount[rune(word1[j])]++
				j++
			}
		}
		if j >= len(word1) {
			break
		}
	}
	if validCheck(tempCount, template) {
		count = count + 1 + int64(len(word1)-j)
		// 处理到达末尾的情况
		tempCount[rune(word1[i])]--
		i++
		for i < len(word1) {
			if validCheck(tempCount, template) {
				count++
				tempCount[rune(word1[i])]--
				i++
			} else {
				break
			}
		}
	}
	return count
}

// 校验相等
func validCheck(map1, template map[rune]int) bool {
	fmt.Printf("map1:%+v, template:%+v\n", map1, template)
	// 遍历map
	for k, v := range template {
		if _, ok := map1[k]; !ok {
			return false
		}
		if map1[k] < v {
			return false
		}
	}
	return true
}
