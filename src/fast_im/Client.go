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
	Conn       net.Conn
	flag       int
}

func NewClient(serverIp string, serverPort int) *Client {
	client := &Client{
		ServerIp:   serverIp,
		ServerPort: serverPort,
		flag:       999,
	}

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	if err != nil {
		fmt.Println("Client connect error: ", err)
		return nil
	}

	client.Conn = conn
	return client
}

func (client *Client) menu() bool {
	var flag int
	fmt.Println("1.公聊模式")
	fmt.Println("2.私聊模式")
	fmt.Println("3.更新用户名")
	fmt.Println("0.退出")

	fmt.Scanln(&flag)
	if flag < 0 || flag > 3 {
		return false
	}
	client.flag = flag
	return true
}

func (client *Client) Run() {
	for client.flag != 0 {
		for client.menu() != true {

		}

		switch client.flag {
		case 1:
			fmt.Println("公聊模式选择...")
			client.PublicChat()
			break
		case 2:
			fmt.Println("私聊模式选择...")
			client.PrivateChat()
			break
		case 3:
			fmt.Println("更新用户名...")
			client.UpdateName()
			break
		}
	}
}

func (client *Client) UpdateName() bool {
	fmt.Println("请输入用户名: ")
	fmt.Scanln(&client.Name)

	sendMsg := "rename|" + client.Name + "\n"
	_, err := client.Conn.Write([]byte(sendMsg))
	if err != nil {
		return false
	}
	return true
}

func (client *Client) DealResponse() {
	io.Copy(os.Stdout, client.Conn)
}

func (client *Client) PublicChat() {
	var chatMsg string
	for {
		fmt.Println("请输入群聊信息")
		fmt.Scanln(&chatMsg)
		client.Conn.Write([]byte(chatMsg + "\n"))
	}
}

func (client *Client) PrivateChat() {

	client.Conn.Write([]byte("who\n"))
	var chatMsg string
	for {
		fmt.Println("请输入群聊信息")
		fmt.Scanln(&chatMsg)
		client.Conn.Write([]byte("to|" + chatMsg + "\n"))
	}
}

func main() {
	flag.Parse()
	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println("Client connect error")
		return
	}
	fmt.Println("Client connect success")

	go client.DealResponse()

	client.Run()
}

var serverIp string
var serverPort int

func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "设置服务器IP地址")
	flag.IntVar(&serverPort, "port", 8888, "设置服务器端口")
}
