package main

import (
	"fmt"
	"time"
)

//func main() {
//	go task()
//
//	i := 0
//	for {
//		i++
//		fmt.Println("main: ", i)
//		time.Sleep(1)
//	}
//}

func task() {
	i := 0
	for {
		i++
		fmt.Println("task goroutine: ", i)
		time.Sleep(1)
	}
}
