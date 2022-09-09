package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

// FooReader 定义一个从标准输入读取的io.Reader
type FooReader struct{}

func (fooReader *FooReader) Read(b []byte) (int, error) {
	fmt.Print("in >")
	return os.Stdin.Read(b)
}

// 写
type FooWriter struct{}

func (fooWriter *FooWriter) Write(b []byte) (int, error) {
	fmt.Print("out <")
	return os.Stdout.Write(b)
}

func main() {
	var (
		reader FooReader
		writer FooWriter
	)

	// 创建缓冲区
	// input := make([]byte, 4096)

	// // 使用reader 读取输入
	// s, err := reader.Read(input)
	// if err != nil {
	// 	log.Fatalln("Unable to read data")
	// }

	// // 使用writer 写入输出
	// s, err = writer.Write(input)
	// if err != nil {
	// 	log.Fatalln("Unable to write data")
	// }
	// fmt.Printf("Wrote %d bytes to stdout\n", s)

	// 上述代码可以直接替换成如下
	if _, err := io.Copy(&writer, &reader); err != nil {
		log.Fatalln("Unable to write data")
	}

}
