package main

import "fmt"

func a() {
	fmt.Println("A")
}

func b() {
	fmt.Println("B")
}

func c() {
	fmt.Println("C")
}

// defer 和 return return先执行
func deferFunc() int {
	fmt.Println("defer call....")
	return 0
}

func returnFunc() int {
	fmt.Println("return call....")
	return 0
}

func returnAndDefer() int {
	defer deferFunc()
	return returnFunc()
}

func main() {
	// defer 在函数完全执行结束后执行
	// defer执行顺序先定义后执行  以下先执行end2 再执行end
	// defer fmt.Println("defer end1")
	// defer fmt.Println("defer end2")

	// fmt.Println("start .....")
	// 执行顺序 CBA
	defer a()
	defer b()
	defer c()

	returnAndDefer()
}
