package main

// 用户业务封装
func main() {
	server := NewServer("127.0.0.1", 8888)
	server.Start()
}
