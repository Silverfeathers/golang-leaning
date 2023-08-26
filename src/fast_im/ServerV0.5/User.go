package main

import "net"

type User struct {
	Name   string
	Addr   string
	C      chan string
	conn   net.Conn
	server *Server
}

func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()
	user := &User{
		Name:   userAddr,
		Addr:   userAddr,
		C:      make(chan string),
		conn:   conn,
		server: server,
	}

	go user.ListenMessage()
	return user
}

func (this *User) ListenMessage() {
	for {
		msg := <-this.C
		this.conn.Write([]byte(msg + "\n"))
	}
}

func (this *User) Online() {
	this.server.mapLock.Lock()
	this.server.OnlineMap[this.Name] = this
	this.server.mapLock.Unlock()
	this.server.BroadCast(this, "online")
}

func (this *User) Offline() {
	this.server.mapLock.Lock()
	delete(this.server.OnlineMap, this.Name)
	this.server.mapLock.Unlock()
	this.server.BroadCast(this, "offline")
}

func (this *User) DoMessage(msg string) {
	if msg == "who" {
		for _, v := range this.server.OnlineMap {
			msg = "[" + v.Name + "]" + v.Addr + ": is online" + "\n"
			this.SendMsg(msg)
		}
	} else {
		this.server.BroadCast(this, msg)
	}
}

func (this *User) SendMsg(msg string) {
	this.conn.Write([]byte(msg))
}
