# HTTP相关

## `00-http-base`
http基础使用代码不过多去写

## `01-shodan-client`
示例代码中，http的json解析常用struct + json.NewDecode().Decode()的方式进行

由于python的原因，本人更喜欢直接解析为map
```go
json_str := make(map[string]interface{})
body, _ := ioutil.ReadAll(resp.Body)
fmt.Println(string(body))
err := json.Unmarshal(body, &json_str)
```

同时该项目中，使用了go1.18的新特性，work模式进行包管理，运行时只需在`/cmd/shodan`下直接`go run main.go`即可
