package main

import (
	"fmt"
	"time"
)

// 生产者
func producer(ch chan int) {
	for i := 0; i < 10; i++ {
		fmt.Printf("Producing %d\n", i)
		ch <- i                 // 发送数据
		time.Sleep(time.Second) // 休眠1秒
	}
	close(ch) // 关闭通道
	// TODO: 关闭通道发送消息
	ch <- 1
}

// 消费者函数
func consumer(ch <-chan int) {
	for num := range ch { // 从通道中接收数据
		fmt.Printf("Consuming: %d\n", num)
		time.Sleep(2 * time.Second) // 模拟处理时间
	}
}

func main() {
	ch := make(chan int, 5) // 创建一个带缓冲的整型通道，缓冲区大小为 5

	go producer(ch) // 启动生产者 goroutine
	go consumer(ch) // 启动消费者 goroutine

	// 等待消费者完成
	time.Sleep(20 * time.Second)
	fmt.Println("Done")
}
