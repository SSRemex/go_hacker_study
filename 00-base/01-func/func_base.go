package main

import (
	"fmt"
)

// 函数定义
// func 函数名(参数1 类型, 参数2 类型) (返回类型1, 返回类型2)（没有则不写) { ... }
// 匿名返回值
func foo1(a, b string) (int, int) {

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	c := 100
	return c, c
}

// 形参返回值
// 当返回类型相同时，也可以这么写func foo2() (r1, r2 int)
func foo2() (r1 int, r2 int) {
	r1 = 1
	r2 = 2
	return
}

func main() {
	c, d := foo1("abc", "123")
	fmt.Println("c = ", c, " d = ", d)

	e, f := foo2()
	fmt.Println("e = ", e, "f = ", f)
}
