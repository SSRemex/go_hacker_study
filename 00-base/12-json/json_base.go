package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"Year"`
	Price  int      `josn:"rmb"`
	Actors []string `json:"actors"`
}

func main() {
	movie := Movie{
		Title: "大幅度",
		Year:  1234,
		Price: 123,
		Actors: []string{
			"asdf",
			"sadfas",
			"asdf",
		},
	}

	// 结构体 --> json
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("json marshal error", err)
		return
	}

	fmt.Printf("jsonStr = %s\n", jsonStr)

	// json --> 结构体
	myjson := "{\"title\":\"大幅度\",\"Year\":1234,\"Price\":123,\"actors\":[\"asdf\",\"sadfas\",\"asdf\"]}"
	myMovie := Movie{}
	err = json.Unmarshal([]byte(myjson), &myMovie)
	if err != nil {
		fmt.Println("json marshal error", err)
		return
	}
	fmt.Println(myMovie)

}
