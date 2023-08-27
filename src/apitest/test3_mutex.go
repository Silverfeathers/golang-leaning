package main

import (
	"fmt"
	"sync"
)

var num = 100
var lock sync.Mutex
var wg sync.WaitGroup

func add() {
	lock.Lock()
	num++
	fmt.Println(num)
	lock.Unlock()
	defer wg.Done()
}

func sub() {
	lock.Lock()
	num--
	fmt.Println(num)
	lock.Unlock()
	defer wg.Done()
}

func testNutex() {
	for i := 0; i < 100; i++ {
		go add()
		wg.Add(1)
		go sub()
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println("end: ", num)
}
