# 网络扫描代理

## `00-port-scanner`
端口扫描

**知识点**
`sync.waitGroup`
主goroutine结束后程序直接终止，会造成子goroutine不能正常执行完成通过WaitGroup，让主goroutine等待子goroutine执行结束后再停止
```go
var wg sync.waitGroup
wg.Add(1)
wg.Done()
wg.Wait()

```

## `01-io-operation`
io操作练习

所有的io操作都需要继承如下接口
```go
type Reader interface{
    Read(p []byte) (n int, err error)
}

type Writer interface{
    Write(p []byte) (n int, err error)
}

```

## `02-tcp-proxy`
tcp代理


## `03-netcat`
本地测试 `go run main.go` 后 通过 `nc 127.0.0.1 7778` 便可执行shell命令


## 知识点
显示刷新Flush()
io.Copy 进行io输出优化
io.Pipe() 同步内存管道