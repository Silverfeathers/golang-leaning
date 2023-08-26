package main

// 广播上线功能
func main() {
	server := NewServer("127.0.0.1", 8888)
	server.Start()
}
