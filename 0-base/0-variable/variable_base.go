package main

import (
	"fmt"
)

// 声明全局变量， 可以使用方法一 方法二 方法三
var A int
var B int = 100
var C = 99

func single_variable() {
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
	// 只能用在函数体内，不能用在全局
	e := 100
	fmt.Println("e = ", e)
	fmt.Printf("type of e = %T\n", e)

}

func multiple_variable() {
	/*
		多变量声明
	*/
	var xx, yy = 100, 0
	fmt.Println("xx = ", xx, "yy = ", yy)
	var qq, ww = true, "asd"
	fmt.Println("qq = ", qq, "ww = ", ww)
	var (
		tt int  = 1
		ii bool = false
	)
	fmt.Println("tt = ", tt, "ii = ", ii)
}

func const_variable() {
	// const 常量即只读 将正常变量定义时的var替换为const即可
	const a int = 1
	fmt.Println("a = ", a)
	// const 枚举类型，iota只能在const使用
	// iota初始值为0，每次自增1
	const (
		BEIJING = iota
		TIANJING
		SHANGHAI = iota * 10
	)
	fmt.Println("BEIJING = ", BEIJING)
	fmt.Println("TIANJING = ", TIANJING)
	fmt.Println("SHNGHAI = ", SHANGHAI)

}

func main() {
	// single_variable()
	// multiple_variable()
	const_variable()
}
