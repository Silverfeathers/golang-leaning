package main

// 超时强踢
func main() {
	server := NewServer("127.0.0.1", 8888)
	server.Start()
}
