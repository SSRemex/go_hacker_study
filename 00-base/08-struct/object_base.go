package main

import "fmt"

type Human struct {
	name string
	sex  string
}

func (this *Human) Eat() {
	fmt.Println(this.name, " eat .........")
}

func (this *Human) Walk() {
	fmt.Println(this.name, " walk .........")
}

// 继承Human
type SuperMan struct {
	Human
	Level int
}

// 重定义父类方法
func (this *SuperMan) Walk() {
	fmt.Println(this.name, " super walk .........")
}

// 新增子类方法
func (this *SuperMan) Fly() {
	fmt.Println(this.name, " fly .........")
}

func main() {
	man := Human{
		name: "remex",
		sex:  "male",
	}
	man.Walk()
	man.Eat()

	// 继承组合定义 方式一
	// superman := SuperMan{
	// 	Human{
	// 		name: "ssremex",
	// 		sex:  "male",
	// 	},
	// 	1,
	// }

	// 方式二
	var superman SuperMan
	superman.name = "ssremex"
	superman.sex = "male"
	superman.Level = 1

	superman.Walk()
	superman.Fly()
}
