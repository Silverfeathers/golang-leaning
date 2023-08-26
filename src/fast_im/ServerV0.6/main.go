package main

// 修改用户名
func main() {
	server := NewServer("127.0.0.1", 8888)
	server.Start()
}
