package main

import "fmt"

/*
	pair概念 主要用在断言机制，类型一致即可断言成功
	pair 变量的内在属性，存储变量类型以及值
	其中type分为 静态类型statictype 以及 混合类型concretetype
	<type:, value:>
	赋值时会将pair传递
*/

type Read interface {
	ReadBook()
}
type Write interface {
	WriteBook()
}

type Book struct {
}

func (this *Book) ReadBook() {
	fmt.Println("read book ....")
}

func (this *Book) WriteBook() {
	fmt.Println("write book ....")
}

func main() {
	// var a interface{} = 123
	// value, _ := a.(int)
	// fmt.Println(value)

	// b: pari<type:Book, value:&Book{}>
	b := &Book{}

	// r: pari<type:, value:>
	var r Read
	// r: pari<type:Book, value:&Book{}>
	r = b
	r.ReadBook()
	// w: pari<type:, value:>
	var w Write
	// w: pari<type:Book, value:&Book{}>
	w = r.(Write) // 断言成功 因为 w,r type一致
	w.WriteBook()

}
