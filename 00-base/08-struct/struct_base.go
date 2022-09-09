package main

import "fmt"

/*
	type 用法
*/

// 给int起了一个myint的别名
type myint int

// 所有类 方法 变量 首字母大写就是公有  首字母小写就是私有
type Book struct {
	name   string
	author string
}

func changeBook1(book Book) {
	book.name = "the car"
}

func changeBook2(book *Book) {
	book.name = "the bike"
}

func main() {

	// var book1 Book
	// book1.name = "the light of city"
	// book1.author = "remex"

	// fmt.Printf("%v \n", book1)
	// changeBook1(book1)
	// fmt.Printf("%v \n", book1)
	// changeBook2(&book1)
	// fmt.Printf("%v \n", book1)

	hero := Hero{
		Name:  "remex",
		Ad:    100,
		Level: 1,
	}

	hero.PrintInfo()
	hero.SetName("test")
	hero.PrintInfo()

}

type Hero struct {
	Name  string
	Ad    int
	Level int
}

// // this是当前对象的一个副本拷贝 所以Set Name并不会生效
// func (this Hero) GetName() string {
// 	return this.Name
// }

// func (this Hero) SetName(name string) {
// 	this.Name = name
// }

// func (this Hero) PrintInfo() {
// 	fmt.Println("Name = ", this.Name, " Ad = ", this.Ad, " Level = ", this.Level)
// }

// this *Object是当前对象的一个引用
func (this *Hero) GetName() string {
	return this.Name
}

func (this *Hero) SetName(name string) {
	this.Name = name
}

func (this *Hero) PrintInfo() {
	fmt.Println("Name = ", this.Name, " Ad = ", this.Ad, " Level = ", this.Level)
}
