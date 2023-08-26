package main

// 在线用户查询
func main() {
	server := NewServer("127.0.0.1", 8888)
	server.Start()
}
