package main

import "fmt"

func main() {
	// 切片声明 给定初始值
	slice1 := []int{1, 2, 3}
	// %v 打印出全部信息
	fmt.Printf("len = %d, slice = %v\n", len(slice1), slice1)
	// 声明 不分配空间
	var slice2 []int
	fmt.Printf("len = %d, slice = %v\n", len(slice2), slice2)
	// 声明 不做初始值同时分配空间
	var slice3 []int = make([]int, 3)
	// slice3 []int := make([]int, 3)
	fmt.Printf("len = %d, slice = %v\n", len(slice3), slice3)
	slice3 = append(slice3, 1)
	// cap 代表当前容量，如果超过则会增加初始的cap的长度
	fmt.Printf("len = %d, cap=%d, slice = %v\n", len(slice3), cap(slice3), slice3)
	// : 用法，代表引用
	// slice4 引用了slice的0-1个元素
	slice4 := slice3[0:2]
	fmt.Printf("len = %d, cap=%d, slice = %v\n", len(slice4), cap(slice4), slice4)
	slice3[0] = 100
	fmt.Printf("len = %d, cap=%d, slice = %v\n", len(slice4), cap(slice4), slice4)
	// 深拷贝copy
	var slice5 = make([]int, 3)
	copy(slice5, slice3)
	slice3[1] = 100
	fmt.Printf("len = %d, cap=%d, slice = %v\n", len(slice5), cap(slice5), slice5)

}
