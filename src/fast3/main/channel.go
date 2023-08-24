package main

import "fmt"

func main() {
	//c := make(chan int)
	//
	//go func() {
	//	defer fmt.Println("goroutine defer")
	//	fmt.Println("goroutine")
	//	c <- 666
	//}()
	//
	//num := <-c
	//fmt.Println(num)

	//c := make(chan int, 3)
	//go func() {
	//	defer fmt.Println("goroutine defer")
	//	fmt.Println("goroutine")
	//	for i := 0; i < 5; i++ {
	//		c <- i
	//		fmt.Println("goroutine add: ", i)
	//	}
	//}()
	//
	////time.Sleep(1000)
	//
	//defer fmt.Println("main defer")
	//fmt.Println("main")
	//for i := 0; i < 6; i++ {
	//	a := <-c
	//	fmt.Println("main get: ", a)
	//}

	//c := make(chan int)
	//go func() {
	//	defer fmt.Println("goroutine defer")
	//	fmt.Println("goroutine")
	//	for i := 0; i < 5; i++ {
	//		c <- i
	//		fmt.Println("goroutine add: ", i)
	//	}
	//	close(c)
	//}()
	//
	////for {
	////	data, ok := <-c
	////	if ok {
	////		fmt.Println("main get: ", data)
	////	} else {
	////		break
	////	}
	////}
	//
	//for data := range c {
	//	fmt.Println("main get: ", data)
	//}

	a, b := make(chan int), make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			t := <-a
			fmt.Println(t)
		}
		b <- 0
	}()
	fubonacii(a, b)
}

func fubonacii(a, b chan int) {
	x, y := 1, 1
	for {
		select {
		case a <- x:
			x = y
			y = x + y
		case <-b:
			fmt.Println("quit")
			return
		}
	}
}
