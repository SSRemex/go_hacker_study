package main

import (
	"fmt"
	"time"
)

// 子goroutine
func newTask() {
	i := 0
	for {
		i++
		fmt.Println("new goroutine: i = ", i)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	// 创建一个goroutine
	// go newTask()

	// 当主线程执行停止时，程序结束，goroutine不会再执行
	// fmt.Println("main goroutine exit")

	// // 主线程一起
	// i := 0
	// for {
	// 	i++
	// 	fmt.Println("main goroutine: i = ", i)
	// 	time.Sleep(1 * time.Second)
	// }

	// 用go创建一个形参为空 返回值为空的一个匿名函数
	// 匿名函数自调用
	// go func() {
	// 	defer fmt.Println("A.defer")

	// 	func() {
	// 		defer fmt.Println("B.defer")

	// 		fmt.Println("B")

	// 	}()
	// 	fmt.Println("A")
	// }()

	// 匿名带参函数goroutine
	// 如果想要获取返回值需要channel
	go func(a int, b int) bool {

		fmt.Println("a = ", a, " b = ", b)
		return true
	}(10, 20)

	for {
		time.Sleep(1 * time.Second)
	}

}
