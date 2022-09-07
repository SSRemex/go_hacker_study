package main

import (
	"net"
	"strings"
)

type User struct {
	Name   string
	Addr   string
	C      chan string
	conn   net.Conn
	server *Server
}

// 创建一个用户的API
func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()
	user := &User{
		Name:   userAddr,
		Addr:   userAddr,
		C:      make(chan string),
		conn:   conn,
		server: server,
	}

	// 监听当前user channel消息的goroutine
	go user.ListenMessage()

	return user
}

func (this *User) Online() {
	// 用户上线 将用户添加至OnlineMap中
	this.server.mapLock.Lock()
	this.server.OnlineMap[this.Name] = this
	this.server.mapLock.Unlock()

	// 用户上线广播
	this.server.BroadCast(this, "已上线")
}

func (this *User) Offline() {
	// 用户下线
	this.server.mapLock.Lock()
	delete(this.server.OnlineMap, this.Name)
	this.server.mapLock.Unlock()

	this.server.BroadCast(this, "下线")
}

func (this *User) SendMsg(msg string) {
	this.conn.Write([]byte(msg))
}

func (this *User) DoMessage(msg string) {

	if msg == "/who" {
		for _, cli := range this.server.OnlineMap {
			onlineMessage := "[" + cli.Addr + "]" + cli.Name + ":" + "在线...\n"
			this.SendMsg(onlineMessage)
		}
	} else if len(msg) > 8 && msg[:8] == "/rename|" {
		// 消息格式 /rename|SSRemex
		newName := msg[8:]

		// 判断name是否存在
		_, ok := this.server.OnlineMap[newName]
		if ok {
			this.SendMsg("当前用户名被使用\n")
		} else {
			this.server.mapLock.Lock()
			delete(this.server.OnlineMap, this.Name)
			this.Name = newName
			this.server.OnlineMap[this.Name] = this
			this.server.mapLock.Unlock()
			this.SendMsg("用户名" + this.Name + "修改成功\n")
		}

	} else if len(msg) > 4 && msg[:4] == "/to|" {
		// 消息格式 /to|remex|content
		// 获取对方用户名
		remoteName := strings.Split(msg, "|")[1]
		if remoteName == "" {
			this.SendMsg("格式错误，格式应为消息格式 /to|remex|content\n")
			return
		}
		// 根据用户名获取对方User对象
		remoteUser, ok := this.server.OnlineMap[remoteName]
		if !ok {
			this.SendMsg("该用户名不存在\n")
			return
		}

		// 获取消息内容通过对方的User对象将消息内容发送出去
		content := strings.Split(msg, "|")[2]
		if content == "" {
			this.SendMsg("无消息内容，请重新发送\n")
			return
		} else {
			remoteUser.SendMsg(this.Name + "对你说：" + content + "\n")
		}

	} else {
		this.server.BroadCast(this, msg)
	}

}

// 监听当前User channel的方法，一旦有消息就直接发送给对端客户端
func (this *User) ListenMessage() {
	for {
		msg := <-this.C
		this.conn.Write([]byte(msg + "\n"))
	}
}
