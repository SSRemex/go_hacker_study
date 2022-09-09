package main

import "fmt"

// 动态数组传递为引用传递指针用法
// 若为array [4]int 固定数组作为形参则不代表指针引用
func printArray(array []int) {
	// _ 表示匿名变量可以不使用
	for _, value := range array {
		fmt.Println(" value = ", value)
	}
	array[0] = 999
}

func main() {
	// 固定长度数组
	var array1 = [4]int{1, 2, 3, 4}
	array2 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("array1-------")
	for i := 0; i < len(array1); i++ {
		fmt.Println("index = ", i, " value = ", array1[i])
	}

	fmt.Println("array2-------")
	for index, value := range array2 {
		fmt.Println("index = ", index, " value = ", value)
	}
	fmt.Printf("the type of array1 is %T\n", array1)
	fmt.Printf("the type of array2 is %T\n", array2)

	fmt.Println("==================")
	// 动态数组即为切片
	var array3 = []int{1, 2, 3}
	array4 := []int{1, 3, 4, 5, 6}
	fmt.Printf("the type of array3 is %T\n", array3)
	fmt.Printf("the type of array4 is %T\n", array4)
	printArray(array4)
	fmt.Println(array4[0])

}
