package _chan

import (
	"fmt"
	"time"
)

// 生产者
func producer(ch chan<- int) {
	for i := 0; i < 5; i++ {
		fmt.Printf("Producing %d\n", i)
		ch <- i                 // 发送数据
		time.Sleep(time.Second) // 休眠1秒
	}
	close(ch) // 关闭通道
}

// 消费者
func consumer(ch <-chan int) {
	for num := range ch {
		// 接受数据
		fmt.Printf("Consuming %d\n", num)
		time.Sleep(time.Second)
	}
}

func main() {
	ch := make(chan int)
	// 启动生产者
	go producer(ch)
	// 启动消费者
	go consumer(ch)

	// 等待消费完成
	time.Sleep(10 * time.Second)
	fmt.Println("Done")
}
