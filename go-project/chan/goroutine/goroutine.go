package main

import (
	"fmt"
	"strconv"
	"sync/atomic"
	"time"
)

/**
 * 保证goroutine执行
 */
func main() {
	main3()
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

/**
 * 保证按顺序执行
 */
func main3() {
	var count uint32 = 0
	trigger := func(i uint32, fn func()) {
		for {
			if n := atomic.LoadUint32(&count); n == i {
				fn()
				atomic.AddUint32(&count, 1)
				break
			}
			time.Sleep(time.Nanosecond * 10)
		}
	}
	for i := uint32(0); i < 10; i++ {
		go func(i uint32) {
			fn := func() {
				fmt.Println(i)
			}
			trigger(i, fn)
		}(i)
	}
	time.Sleep(time.Second * 5)
}
