package main

// 基础服务构建
func main() {
	server := NewServer("127.0.0.1", 8888)
	server.Start()
}
