package main

/*
	go 指针
*/
import (
	"fmt"
)

func swap(pa *int, pb *int) {
	var temp int
	temp = *pa
	*pa = *pb
	*pb = temp
}

func main() {
	var a int = 1
	var b int = 2

	fmt.Println("a = ", &a)
	fmt.Println("a = ", *&a)
	fmt.Println("a = ", a, "b = ", b)
	swap(&a, &b)
	fmt.Println("a = ", a, "b = ", b)

	// 二级指针
	var p *int
	var pp **int
	p = &a
	pp = &p

	fmt.Println("----------------")
	fmt.Println("p = ", p)
	fmt.Println("*p = ", *p)
	fmt.Println("&p = ", &p)
	fmt.Println("pp = ", pp)
	fmt.Println("*pp = ", *pp)
	fmt.Println("**pp = ", **pp)
	fmt.Println("&pp = ", &pp)

}
