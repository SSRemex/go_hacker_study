package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	conn       net.Conn
	flag       int // 当前用户的模式
}

func NewClient(serverIp string, serverPort int) *Client {
	// 创建client对象
	client := &Client{
		ServerIp:   serverIp,
		ServerPort: serverPort,
		flag:       999,
	}

	// 连接server
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	if err != nil {
		fmt.Println("net.Dial error: ", err)
		return nil
	}

	client.conn = conn

	return client

}

// 处理返回结果
func (this *Client) DealResponse() {
	// 一旦this.conn 有数据，就直接copy显示到stdout输出上，永久阻塞监听
	io.Copy(os.Stdout, this.conn)

}

func (this *Client) menu() bool {
	var flag int

	fmt.Println("1.公聊模式")
	fmt.Println("2.私聊模式")
	fmt.Println("3.修改名称")
	fmt.Println("0.退出")

	fmt.Scanln(&flag)
	if flag >= 0 && flag <= 3 {
		this.flag = flag
		return true
	} else {
		fmt.Println("请输入合法的数字范围")
		return false
	}
}

func (this *Client) UpdateName() bool {
	fmt.Println("请输入用户名：")
	fmt.Scanln(&this.Name)

	sendMsg := "/rename|" + this.Name + "\n"
	_, err := this.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn.Write err: ", err)
		return false
	}
	return true
}

// 查询在线用户
func (this *Client) SelectUsers() {
	sendMsg := "/who\n"
	_, err := this.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn.Write err: ", err)
		return
	}

}

// 私聊模式
func (this *Client) PrivateChat() {
	var remoteName string
	var chatMsg string

	this.SelectUsers()

	fmt.Println(">>>>> 请输入聊天对象[用户名]，exit退出：")
	fmt.Scanln(&remoteName)

	for remoteName != "exit" {
		fmt.Print(">>>>> 请输入聊天内容，exit退出：")
		fmt.Scanln(&chatMsg)

		for chatMsg != "exit" {
			chatMsg = "/to|" + remoteName + "|" + chatMsg + "\n"
			_, err := this.conn.Write([]byte(chatMsg))
			if err != nil {
				fmt.Println("conn.Write err: ", err)
			}

			chatMsg = ""
			fmt.Scanln(&chatMsg)
		}
		fmt.Println(">>>>> 请输入聊天对象[用户名]，exit退出：")
		this.SelectUsers()
		fmt.Scanln(&remoteName)
	}

}

func (this *Client) PublicChat() {
	// 提示用户输入消息
	var chatMsg string
	fmt.Println(">>> 请输入聊天内容，exit推出")
	fmt.Scanln(&chatMsg)

	for chatMsg != "exit" {
		// 发送给服务器
		if len(chatMsg) != 0 {
			sendMsg := chatMsg + "\n"
			_, err := this.conn.Write([]byte(sendMsg))
			if err != nil {
				fmt.Println("conn.Write err: ", err)
				break
			}
		}
		chatMsg = ""
		fmt.Scanln(&chatMsg)

	}

}

func (this *Client) Run() {
	for this.flag != 0 {
		for this.menu() != true {
		}

		// 根据不同的flag处理不同的业务
		// go switch case 不需要break
		switch this.flag {
		case 1:
			// 公聊模式
			this.PublicChat()
		case 2:
			// 私聊模式
			this.PrivateChat()
		case 3:
			// 更新名称
			this.UpdateName()

		}

	}
}

var serverIp string
var serverPort int

func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "默认服务器IP:127.0.0.1")
	flag.IntVar(&serverPort, "port", 7788, "默认服务器Port: 7788")
}

func main() {
	// 命令行解析
	flag.Parse()

	// 创建客户端
	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println(">>>>>> 连接服务器失败")
		return
	}

	fmt.Println(">>>>>> 连接服务器成功")

	go client.DealResponse()

	// 启动客户端
	client.Run()

}
