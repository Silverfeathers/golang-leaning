package main

// 用户消息广播
func main() {
	server := NewServer("127.0.0.1", 8888)
	server.Start()
}
