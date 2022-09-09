package main

import "fmt"

/*
	map 即 字典
*/

// map的形参是引用
func printMap(mymap map[int]string) {
	for key, value := range mymap {
		fmt.Println("key = ", key)
		fmt.Println("value = ", value)
	}
	mymap[3] = "go"
}

func main() {
	var map1 = make(map[string]string)
	// map1 := make(map[string]string)
	map1["a"] = "go"
	map1["b"] = "python"
	fmt.Println(map1)

	map2 := map[int]string{
		1: "python",
		2: "c++",
	}
	fmt.Println(map2)
	printMap(map2)

	fmt.Println("------------")

	delete(map2, 1)
	printMap(map2)
}
