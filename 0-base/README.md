# GO基础
本目录是对go的基础语法学习

## 运行
`go run target.go` 直接运行main函数
`go build` 编译
`go env -w GO111MODULE=on` 修改go环境变量开启go modules

## 特殊用法记录

`go build -ldflags "-w -s" target.go` 禁用默认情况下生成二进制文件中所包含的调试信息和字符表，该命令可以使编译后的可执行文件体积减小大约30%

**交叉编译**
bash下可以 `GOSO="windows" GOARCH="amd64" go build .\hello.go`
windows下需修改env 过于复杂

`go doc fmt.Println` 查询有关包，函数，方法或变量的文档

`go get 包` 下载包并放在$GOPATH/src 目录中

`go fmt target.go` 更改文件格式 `golint targt.go` 报告样式格式错误  `go vet`识别可疑构造


## import机制
go在import导包时会采取如下执行顺序
当前文件(import pkg_A) -> pkg_A 中(import -> init() )

**`go mod`使用**

|命令|作用|
|:-:|:-:|
|go mod init|生成go.mod文件|
|go mod download|下载go.mod的所有依赖|
|go mod tidy|整理现有的依赖|
|go mod graph|查看现有的依赖结构|
|go mod edit|编辑go.mod文件|
|go mod vendor|导出项目所有的依赖到vendorm目录|
|go mod verify|校验一个模块是否被篡改过|
|go mod why|查看为什么需要依赖某模块|


