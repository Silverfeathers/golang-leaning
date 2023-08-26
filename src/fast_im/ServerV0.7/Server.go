package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

type Server struct {
	Ip   string
	Port int

	OnlineMap map[string]*User
	mapLock   sync.RWMutex

	Message chan string
}

func NewServer(ip string, port int) *Server {
	server := Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}

	return &server
}

func (this *Server) Start() {
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("链接创建失败:", err)
		return
	}
	defer listen.Close()

	go this.ListenMessager()

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("接收失败:", err)
			return
		}
		fmt.Println("链接创建成功!")
		go this.Handler(conn)
	}
}

func (this *Server) Handler(conn net.Conn) {
	user := NewUser(conn, this)
	user.Online()

	isLive := make(chan bool)
	go func() {
		bytes := make([]byte, 4096)
		for {
			n, err := conn.Read(bytes)
			if err != nil {
				fmt.Println("接收消息错误")
				return
			}
			msg := string(bytes[:n-1])
			if msg == "0" {
				user.Offline()
				return
			}
			user.DoMessage(msg)

			isLive <- true
		}
	}()

	select {
	case <-isLive:

	case <-time.After(time.Second * 10):
		user.SendMsg("get out")
		close(user.C)
		conn.Close()
		return
	}
}

func (this *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ": " + msg + "\n"
	this.Message <- sendMsg
}

func (this *Server) ListenMessager() {
	for {
		msg := <-this.Message
		this.mapLock.Lock()
		for _, v := range this.OnlineMap {
			v.C <- msg
		}
		this.mapLock.Unlock()
	}
}
