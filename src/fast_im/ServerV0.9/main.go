package main

// 客户端实现
func main() {
	server := NewServer("127.0.0.1", 8888)
	server.Start()
}
