package main

import (
	"fmt"
	"time"
)

/*
	channel goroutine通信机制
*/

func noneCache() {

	// 定义一个channel
	// 无缓存用例
	c := make(chan int)

	go func() {
		defer fmt.Println("goroutine 结束")

		fmt.Println("goroutine 正在运行")

		// 将666发送给channel，如果channel没有被获取，则sub goroutine阻塞
		c <- 666
	}()
	// 从channel获取，如果此时没有数据，则main goroutine会阻塞
	num := <-c
	fmt.Println("num = ", num)
	fmt.Println("main goroutine 结束")
}

func cacheChannel() {
	// 定义一个缓存为3的channel
	c := make(chan int, 3)

	go func() {
		defer fmt.Println("goroutine 结束")

		fmt.Println("goroutine 正在运行")
		for i := 0; i < 3; i++ {
			c <- i
			fmt.Println("sub goroutine 正在运行，发送的元素 = ", i, "len = ", len(c), " cap = ", cap(c))
		}
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 3; i++ {
		num := <-c

		fmt.Println("num = ", num)
	}

}

func closeChannel() {
	// 定义一个channel
	c := make(chan int)

	go func() {
		defer fmt.Println("goroutine 结束")

		fmt.Println("goroutine 正在运行")
		for i := 0; i < 3; i++ {
			fmt.Println("sub goroutine 正在运行，发送的元素 = ", i, "len = ", len(c), " cap = ", cap(c))
			c <- i

		}
		// 关闭channel
		// 关闭后无法再使用该channel 发数据 会报错
		// 但是可以接收channel缓存的数据
		close(c)
	}()

	for {
		if data, ok := <-c; ok {
			fmt.Println(data)
		} else {
			break
		}
	}

	fmt.Println("Main Finish .....")

}

func rangeChannel() {
	c := make(chan int)

	go func() {
		defer fmt.Println("goroutine 结束")

		fmt.Println("goroutine 正在运行")
		for i := 0; i < 3; i++ {
			//fmt.Println("sub goroutine 正在运行，发送的元素 = ", i, "len = ", len(c), " cap = ", cap(c))
			c <- 1
		}
		close(c)
	}()

	// channel range特殊使用
	// 通过循环进行迭代
	// 只要go程在运行，就会一直等待，所有go程结束，如果chan没关闭就会异常
	for data := range c {
		fmt.Println(data)
	}
}

func fibonacii(c, quit chan int) {
	x, y := 1, 1

	for {
		select {
		case c <- x:
			// 如果c可写入则会进入该case
			tmp := y
			y = x + y
			x = tmp
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {

	// // 单流程下一个go监控一个channel的状态，select可以完成监控多个channel的状态
	// // channel select 特殊用法
	// c := make(chan int)
	// quit := make(chan int)

	// go func() {
	// 	for i := 0; i < 6; i++ {
	// 		fmt.Println("c: ", <-c)
	// 	}

	// 	quit <- 0
	// }()

	// fibonacii(c, quit)

	rangeChannel()

}
