package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type Server struct {
	Ip   string
	Port int

	// 在线用户列表
	OnlineMap map[string]*User
	mapLock   sync.RWMutex

	// 消息广播的channel
	Message chan string
}

// 服务创建
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
	return server
}

// 消息监听
func (this *Server) ListenMessage() {
	for {
		msg := <-this.Message
		// 将msg发送给全部客户端
		this.mapLock.Lock()
		for _, cli := range this.OnlineMap {
			cli.C <- msg
		}
		this.mapLock.Unlock()

	}
}

// 消息广播
func (this *Server) BroadCast(user *User, msg string) {
	sendMessage := "[" + user.Addr + "]" + user.Name + ":" + msg

	this.Message <- sendMessage
}

func (this *Server) Handle(conn net.Conn) {
	// 当前连接
	// fmt.Println("连接成功")
	user := NewUser(conn, this)
	user.Online()

	isAlive := make(chan bool)

	// 读取用户输入
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				user.Offline()
				return
			}

			if err != nil && err != io.EOF {
				fmt.Println("Conn read error:", err)
				return
			}

			// 把消息中的\n去掉
			msg := string(buf[:n-1])
			// 用户针对msg进行处理
			user.DoMessage(msg)
			isAlive <- true

		}
	}()

	// 当handle阻塞
	for {
		select {
		case <-isAlive:

		case <-time.After(time.Second * 60 * 5):
			// 已经超时
			// 将当前超时的User强制关闭
			user.SendMsg("你被踢了")
			// 销毁资源
			close(user.C)
			// 关闭连接
			conn.Close()
			//退出当前Handle
			return //runtime.Goexit()
		}
	}

}

// 服务启动
func (this *Server) Start() {
	// socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("server start err:", err)
	}

	// close
	defer listener.Close()

	fmt.Println("服务启动....")

	// 启用监听消息的goroutine
	go this.ListenMessage()

	for {
		// accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("server accept err:", err)
			continue
		}

		// do handler
		go this.Handle(conn)

	}

}
