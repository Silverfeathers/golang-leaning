package main

import (
	"fmt"
	lib2 "goproject/fast1/lib2"
)

func main() {
	//c, e := test3("aaa", 10)
	//fmt.Println(c, e)

	//test6()

	//fmt.Println(test7())
	test12()
}

//const (
//	a = iota
//	b
//	c
//
//	d, e = iota * 2, iota * 3
//	f, g
//	h, i
//)

//// 常量
//func test1() {
//	const length = 10
//	fmt.Println(length)
//
//	fmt.Println(a)
//	fmt.Println(b)
//	fmt.Println(c)
//
//	fmt.Println(d)
//	fmt.Println(e)
//	fmt.Println(f)
//	fmt.Println(g)
//	fmt.Println(h)
//	fmt.Println(i)
//}

// 函數
func test2(a string, b int) (int, string) {
	fmt.Println(a)
	fmt.Println(b)
	c := 100
	return c, "hahaha"
}

// 函數
func test3(a string, b int) (c int, d string) {
	fmt.Println(a)
	fmt.Println(b)
	c = 100
	//d = "abc"
	return c, d
}

// 导包
func test5() {
	//lib1.Lib1Test()
	lib2.Lib2Test()
}

// 指针引用传递
func test6() {
	a := 10
	b := 20
	fmt.Println(&a, &b)
	swapper(&a, &b)
	fmt.Println(a, b)
	fmt.Println(&a, &b)
}

func swapper(pa *int, pb *int) {
	var temp int
	temp = *pa
	*pa = *pb
	*pb = temp
}

// defer
func test7() string {
	defer fmt.Println("a")
	defer fmt.Println("b")
	return "c"
}

// slice
func test9() {
	var arr [10]int
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	arr2 := [10]int{1, 2, 3, 4}
	for i, i2 := range arr2 {
		fmt.Println(i, i2)
	}

}

// slice
func test10() {
	arr1 := [10]int{1}
	chageArray(arr1)
	for _, v := range arr1 {
		fmt.Println(v)
	}
	fmt.Println()

	arr2 := []int{1, 1}
	chageArray2(arr2)
	for _, v := range arr2 {
		fmt.Println(v)
	}
}

func chageArray(arr [10]int) {
	arr[1] = 2
}

func chageArray2(arr []int) {
	arr[1] = 2
}

// slice
func test11() {
	arr1 := make([]int, 3)
	fmt.Println(arr1)

	arr2 := make([]int, 3, 5)
	arr2 = append(arr2, 1)
	arr2 = append(arr2, 2)
	arr2 = append(arr2, 3)
	fmt.Println(arr2)

	arr2 = arr2[1:5]
	fmt.Println(arr2)

	arr3 := make([]int, len(arr2), cap(arr2))
	copy(arr3, arr2)
	fmt.Println(arr3)
}

// map
func test12() {
	//第一种声明
	var test1 map[string]string
	//在使用map前，需要先make，make的作用就是给map分配数据空间
	test1 = make(map[string]string, 10)
	test1["one"] = "php"
	test1["two"] = "golang"
	test1["three"] = "java"
	fmt.Println(test1) //map[two:golang three:java one:php]

	//第二种声明
	test2 := make(map[string]string)
	test2["one"] = "php"
	test2["two"] = "golang"
	test2["three"] = "java"
	fmt.Println(test2) //map[one:php two:golang three:java]

	//第三种声明
	test3 := map[string]string{
		"one":   "php",
		"two":   "golang",
		"three": "java",
	}
	fmt.Println(test3) //map[one:php two:golang three:java]
}
