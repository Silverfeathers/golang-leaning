package main

import (
	"fmt"
	"sync"
)

var wp sync.WaitGroup

func waitGroup() {
	for i := 0; i < 10; i++ {
		go sendMsg(i)
		wp.Add(1)
	}
	wp.Wait()
	fmt.Println("end")
}

func sendMsg(i int) {
	defer wp.Done()
	fmt.Println(i)
}
