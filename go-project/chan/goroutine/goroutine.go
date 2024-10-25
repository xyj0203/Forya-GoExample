package main

import (
	"fmt"
	"strconv"
	"time"
)

/**
 * 保证goroutine执行
 */
func main() {
	main2()
}

/**
 * 通过等待保证
 */
func main1() {
	for i := 0; i < 10; i++ {
		go func() {
			println(i)
		}()
	}
	time.Sleep(time.Second * 5)
}

/**
 * 通过管道保证
 */
func main2() {
	// 为什么用这个属性， 因为他的数据值为 0 字节
	chans := make(chan struct{}, 10)
	for i := 0; i < 10; i++ {
		// 调度顺序不一定
		go func() {
			println(i)
			chans <- struct{}{}
		}()
	}
	for i := 0; i < 10; i++ {
		<-chans
		fmt.Println("我是接收的数据： " + strconv.Itoa(i))
	}
	close(chans)
}
