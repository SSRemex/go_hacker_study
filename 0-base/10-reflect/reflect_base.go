package main

/*
	反射机制
*/
import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (this *User) Call() {
	fmt.Print("Id: ", this.Id, " name: ", this.Name, " Age: ", this.Age)
}

func reflectNum(arg interface{}) {
	fmt.Println("type: ", reflect.TypeOf(arg))
	fmt.Println("value: ", reflect.ValueOf(arg))
}

func main() {

	// var a float32 = 1.2345
	// b := 1.545
	// reflectNum(a)
	// reflectNum(b)

	user := User{
		Id:   1,
		Name: "remex",
		Age:  12,
	}
	DoFieldAndMethod(user)

}

func DoFieldAndMethod(input interface{}) {
	// 获取input type
	t := reflect.TypeOf(input)
	fmt.Println("the type of input is :", t.Name())

	// 获取input value
	v := reflect.ValueOf(input)
	fmt.Println("the type of input is :", v)

	// 通过type获取里面的字段
	// 1. 获取interface的reflect.Type，通过Type得到NumField，进行遍历
	// 2. 得到field的数据类型
	// 3. 通过field有一个Interface()方法得到对应的value
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Interface()

		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}

	// 通过type 获取里面的方法 调用
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%s: %v \n", m.Name, m.Type)
	}

}
