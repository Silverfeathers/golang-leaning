package main

//func main() {
//
//	go func() {
//		defer fmt.Println("A defer")
//		fmt.Println("A")
//
//		//return
//		func() {
//			defer fmt.Println("B defer")
//			runtime.Goexit()
//			fmt.Println("B")
//		}()
//	}()
//
//	go func(a int, b int) bool {
//		fmt.Println(a, b)
//		return true
//	}(10, 20)
//
//	for {
//		time.Sleep(1 * time.Second)
//	}
//}
