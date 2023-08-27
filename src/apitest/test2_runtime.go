package main

import (
	"fmt"
	"runtime"
	"time"
)

func testRuntime() {
	fmt.Println(runtime.NumCPU())
	runtime.GOMAXPROCS(3)
	fmt.Println(runtime.NumCPU())

	go show()
	for i := 0; i < 10; i++ {
		runtime.Gosched()
		fmt.Println("main: ", i)
		time.Sleep(time.Second / 5)
	}
	fmt.Println("end")
}

func show() {
	for i := 0; i < 100; i++ {
		if i > 5 {
			runtime.Goexit()
		}
		time.Sleep(time.Second / 5)
		fmt.Println("go: ", i)
	}
}
