package main

import (
	// 起别名  improt xxx as a
	my "demo"
	// _ 代表匿名，匿名引入代码中不调用也不报错
	_ "fmt"
	// 忽略报名 from xxx import *
	. "fmt"
)

func main() {
	my.Callback()
	Print(". 别名  from xxx import *")
}
