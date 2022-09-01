# GO基础
本目录是对go的基础语法学习

## 运行
`go run target.go` 直接运行main函数
`go build` 编译

## 特殊用法记录

`go build -ldflags "-w -s" target.go` 禁用默认情况下生成二进制文件中所包含的调试信息和字符表，该命令可以使编译后的可执行文件体积减小大约30%

**交叉编译**
bash下可以 `GOSO="windows" GOARCH="amd64" go build .\hello.go`
windows下需修改env 过于复杂

`go doc fmt.Println` 查询有关包，函数，方法或变量的文档

`go get 包` 下载包并放在$GOPATH/src 目录中

`go fmt target.go` 更改文件格式 `golint targt.go` 报告样式格式错误  `go vet`识别可疑构造