package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

// echo 回显函数
func echo(conn net.Conn) {
	defer conn.Close()

	// ==========================
	// 低级函数版本
	// b := make([]byte, 512)
	// for {
	// 	size, err := conn.Read(b[0:])
	// 	if err == io.EOF {
	// 		log.Println("Clinet disconnected")
	// 		break
	// 	}
	// 	if err != nil {
	// 		log.Println("Unexpected error")
	// 		break
	// 	}
	// 	log.Printf("Received %d bytes: %s\n", size, string(b))

	// 	// 通过 conn.Write 发送数据
	// 	if _, err := conn.Write(b[0:size]); err != nil {
	// 		log.Fatalln("Unable to write data")
	// 	}
	// }
	// =============================

	// bufio 版本优化缓冲区操作
	// reader := bufio.NewReader(conn)
	// s, _ := reader.ReadString('\n')
	// log.Printf("Read %d bytes: %s\n", len(s), s)
	// log.Println("Writing data")

	// writer := bufio.NewWriter(conn)
	// if _, err := writer.WriteString(s); err != nil {

	// }

	// // 显式调用 使得所有数据写入底层的writer
	// writer.Flush()

	// io.Copy

	if _, err := io.Copy(conn, conn); err != nil {
		log.Fatalln("Unable to read/write data")
	}

}

func main() {
	// 在所有接口绑定TCP 20080 端口
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {

	}

	log.Println("listening successful ")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept error")
		}
		go echo(conn)
	}

}
