# 网络扫描代理

## `0-port-scanner`
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

## `1-io-operation`
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