package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func NormalRequest() {
	// 直接使用通用可全设置的方式 NewRequest
	req, _ := http.NewRequest("GET", "https://www.baidu.com", nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36")
	var clinet http.Client
	// 发起请求
	resp, _ := clinet.Do(req)
	// 解析返回
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func ResponseJson() {
	// {"code":0,"message":"0","ttl":1,"data":{"cover":"","channel_id":0,"channel_name":"","notify":false,"ctype":0,"subscribed_count":0}}
	url := "https://api.bilibili.com/x/web-interface/web/channel/red"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36")
	var client http.Client
	resp, _ := client.Do(req)

	// 用来接收json的map
	json_str := make(map[string]interface{})

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	// 解析json到json_str
	err := json.Unmarshal(body, &json_str)
	if err != nil {

	}

	defer resp.Body.Close()
	fmt.Printf("%v\n", json_str)

}

func main() {
	ResponseJson()
}
