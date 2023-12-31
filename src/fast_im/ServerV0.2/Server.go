package main

import (
	"fmt"
	"net"
	"sync"
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
	user := NewUser(conn)
	this.mapLock.Lock()
	this.OnlineMap[user.Name] = user
	this.mapLock.Unlock()

	this.BroadCast(user, "已上线")
}

func (this *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + msg
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
