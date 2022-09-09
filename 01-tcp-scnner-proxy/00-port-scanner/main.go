package main

import (
	"fmt"
	"net"
	"sort"
)

// 工人池
func worker(ports, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("www.baidu.com:%d", p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			// 有数据时，接收动作阻塞
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

func main() {

	ports := make(chan int, 100)
	results := make(chan int)
	var openports []int

	// 启动worker
	for i := 0; i < 1024; i++ {
		go worker(ports, results)
	}
	// goroutine写入ports
	go func() {
		for i := 1; i <= 1024; i++ {
			ports <- i
		}
	}()

	for i := 0; i < 1024; i++ {
		// 无数据时，接收动作阻塞
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}

	fmt.Println("扫描结束")

	close(ports)
	close(results)

	sort.Ints(openports)
	for _, port := range openports {
		fmt.Printf("%d open!\n", port)
	}
}
