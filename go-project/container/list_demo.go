// Package container 容器
package main

import (
	"container/list"
)

/**
 * 双向链表
 */
func main() {
	// 声明列表
	tmp1 := list.List{}
	var tm2 list.List
	tm2 = list.List{}
	tm3 := list.New()
	tmp1.PushBack(1)
	tm2.PushFront(1)
	tm3.PushFront(1)

}
