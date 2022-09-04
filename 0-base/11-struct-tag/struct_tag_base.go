package main

import (
	"fmt"
	"reflect"
)

/*
	结构体的注解标签 以及标签内容获取
*/

type resume struct {
	Name string `info:"name" doc:"我的名字"`
	Sex  int    `info:"sex" l:"性别"`
}

func findTag(str interface{}) {
	t := reflect.TypeOf(str).Elem()

	for i := 0; i < t.NumField(); i++ {
		tagInfo := t.Field(i).Tag.Get("info")
		tagDoc := t.Field(i).Tag.Get("doc")
		fmt.Println("info: ", tagInfo, " doc: ", tagDoc)

	}
}

func main() {
	re := resume{
		Name: "remex",
		Sex:  1,
	}

	findTag(&re)
}
