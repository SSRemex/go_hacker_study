package main

import (
	"fmt"
)

/*
	多态的实现
	接口用法
*/

// interface的本质是一个指针
// 通过创建接口指针指向接口实现对象
type AnimalIF interface {
	Sleep()
	GetColor() string
	GetType() string
}

//具体的类对接口的实现
// 接口中的方法必须完全实现，否则报错
type Cat struct {
	color string
}

func (this Cat) Sleep() {
	fmt.Println("Cat sleeps...")
}

func (this Cat) GetColor() string {
	return this.color
}

func (this Cat) GetType() string {
	return "Cat"
}

//具体的类对接口的实现
type Dog struct {
	color string
}

func (this *Dog) Sleep() {
	fmt.Println("Dog sleeps...")
}

func (this *Dog) GetColor() string {
	return this.color
}

func (this *Dog) GetType() string {
	return "Dog"
}

// interface {} 万能数据类型 无需再指定数据类型
func myFunc(arg interface{}) {
	fmt.Println("-------万能数据类型接口------------")
	// fmt.Println("the arg is ", arg)

	// interface{} 断言机制来判断数据类型
	// value 会将arg切换为指定类型，如果不能切换就为默认值
	// 强制类型转换
	value, ok := arg.(string)
	if !ok {
		fmt.Println("arg is not string")
		fmt.Printf("the arg type is %T\n", arg)
		fmt.Println("the value is ", value)
	} else {
		fmt.Println("arg is string, the value is ", value)
	}

}

func main() {
	// 定义一个接口
	var animal AnimalIF
	// 指向子类
	animal = &Cat{color: "red"}
	animal.Sleep()
	fmt.Println(animal.GetColor())
	fmt.Println(animal.GetType())

	var animal2 AnimalIF
	animal2 = &Dog{color: "blue"}
	animal2.Sleep()

	// var dog = Dog{color: "yellow"}

	// dog.Sleep()

	// myFunc(1)
	// myFunc("a")
	// myFunc(false)

}
