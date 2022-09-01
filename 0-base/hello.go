package main

import (
	"fmt"
)

func variable() {
	/*
		变量的声明方式
	*/
	// 方法一：声明一个变量 默认值为0
	var a int
	fmt.Println("a = ", a)
	fmt.Printf("type of a = %T\n", a)
	// 方法二：声明一个变量 初始化一个值
	var b int = 1
	fmt.Println("b = ", b)
	fmt.Printf("type of b = %T\n", b)

	var bb string = "abcd"
	fmt.Printf("bb = %s, type of bb = %T\n", bb, bb)

	// 方法三：在初始化时候省去数据类型 自动适配类型
	var c = 100
	fmt.Println("c = ", c)
	fmt.Printf("type of c = %T\n", c)

	// 方法四：省去var关键字，直接自动匹配
	e := 100
	fmt.Println("e = ", e)
	fmt.Printf("type of e = %T\n", e)

}

func main() {
	variable()
}
