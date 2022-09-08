package main

import (
	"io"
	"net"
)

func handle(src net.Conn) {

	dst, err := net.Dial("tcp", "127.0.0.1:7788")
	if err != nil {

	}

	defer dst.Close()

	// 在goroutine中运行 防止io.Copy被阻塞
	go func() {
		// 将源的输出复制到目标
		if _, err := io.Copy(dst, src); err != nil {

		}
	}()

	// 将目标输出返回给源
	go func() {
		if _, err := io.Copy(src, dst); err != nil {

		}
	}()
}

func main() {

	// 监听
	listen, err := net.Listen("tcp", "127.0.0.1:80")
	if err != nil {

	}

	for {
		conn, _ := listen.Accept()
		go handle(conn)
	}

}
