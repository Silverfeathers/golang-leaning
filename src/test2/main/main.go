package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

func main() {
	test15()
}

func test1() {
	var age int
	age = 18
	fmt.Println(age)
}

func test2() {
	var age int = 19
	fmt.Println(age)
}

func test3() {
	var age int
	fmt.Println(age)

	var age2 = 10
	fmt.Println(age2)

	age3 := 12
	fmt.Println(age3)

	var n1, n2, n3 int
	n1 = 10
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)

	var n4, str = 1, "abc"
	fmt.Println(n4)
	fmt.Println(str)

	fmt.Println(n5)
	fmt.Println(str2)
}

var (
	n5   = 11
	str2 = "hahaha"
)

func test5() {
	var n1 int8 = 12
	fmt.Println(n1)

	var n2 = 1200000000000000011
	fmt.Printf("%T", n2)
	fmt.Println()
	fmt.Println(unsafe.Sizeof(n2))
}

// 浮点数
func test6() {
	var num1 float32 = 3.14
	fmt.Println(num1)
	var num2 float32 = -3.14
	fmt.Println(num2)
	var num3 float32 = 314e-2
	fmt.Println(num3)
}

// 字符
func test7() {
	var c4 int = '中'
	fmt.Println(c4)
}

// 类型转换
func test9() {
	//进行类型转换：
	var n1 int = 100
	//var n2 float32 = n1  在这里自动转换不好使，比如显式转换
	fmt.Println(n1)
	//fmt.Println(n2)
	var n2 float32 = float32(n1)
	fmt.Println(n2)
	//注意：n1的类型其实还是int类型，只是将n1的值100转为了float32而已，n1还是int的类型
	fmt.Printf("%T", n1) //int
	fmt.Println()
	//将int64转为int8的时候，编译不会出错的，但是会数据的溢出
	var n3 int64 = 888888
	var n4 int8 = int8(n3)
	fmt.Println(n4) //56
	var n5 int32 = 12
	var n6 int64 = int64(n5) + 30 //一定要匹配=左右的数据类型
	fmt.Println(n5)
	fmt.Println(n6)
	var n7 int64 = 12
	var n8 int8 = int8(n7) + 127 //编译通过，但是结果可能会溢出
	//var n9 int8 = int8(n7) + 128 //编译不会通过
	fmt.Println(n8)
	//fmt.Println(n9)

	var n10 int64 = int64(n5 + 30) //一定要匹配=左右的数据类
	fmt.Println(n10)

	n11 := n10 / 30
	fmt.Println(n11)

	var n13 float64 = float64(n10) / 30
	fmt.Println(n13)
}

// 字符串
func test10() {
	var n1 int = 19
	var n2 float32 = 4.78
	var n3 bool = false
	var n4 byte = 'a'
	var s1 string = fmt.Sprintf("%d", n1)
	fmt.Printf("s1对应的类型是：%T ，s1 = %q \n", s1, s1)
	var s2 string = fmt.Sprintf("%f", n2)
	fmt.Printf("s2对应的类型是：%T ，s2 = %q \n", s2, s2)
	var s3 string = fmt.Sprintf("%t", n3)
	fmt.Printf("s3对应的类型是：%T ，s3 = %q \n", s3, s3)
	var s4 string = fmt.Sprintf("%c", n4)
	fmt.Printf("s4对应的类型是：%T ，s4 = %q \n", s4, s4)
}

// 字符串
func test11() {
	var n1 int = 18
	var s1 string = strconv.FormatInt(int64(n1), 10) //参数：第一个参数必须转为int64类型 ，第二个参数指定字面值的进制形式为十进制
	fmt.Printf("s1对应的类型是：%T ，s1 = %q \n", s1, s1)
	var n2 float64 = 4.29
	var s2 string = strconv.FormatFloat(n2, 'f', 9, 64)
	//第二个参数：'f'（-ddd.dddd）  第三个参数：9 保留小数点后面9位  第四个参数：表示这个小数是float64类型
	fmt.Printf("s2对应的类型是：%T ，s2 = %q \n", s2, s2)
	var n3 bool = true
	var s3 string = strconv.FormatBool(n3)
	fmt.Printf("s3对应的类型是：%T ，s3 = %q \n", s3, s3)

	s5 := "" + strconv.FormatInt(int64(n1), 8)
	fmt.Println(s5)
}

// 字符串转基本数据类型
func test12() {
	//string-->bool
	var s1 string = "true"
	var b bool
	//ParseBool这个函数的返回值有两个：(value bool, err error)
	//value就是我们得到的布尔类型的数据，err出现的错误
	//我们只关注得到的布尔类型的数据，err可以用_直接忽略
	b, _ = strconv.ParseBool(s1)
	fmt.Printf("b的类型是：%T,b=%v \n", b, b)

	//string---》int64
	var s2 string = "19"
	var num1 int64
	num1, _ = strconv.ParseInt(s2, 10, 64)
	fmt.Printf("num1的类型是：%T,num1=%v \n", num1, num1)

	//string-->float32/float64
	var s3 string = "3.14"
	var f1 float64
	f1, _ = strconv.ParseFloat(s3, 64)
	fmt.Printf("f1的类型是：%T,f1=%v \n", f1, f1)

	//注意：string向基本数据类型转换的时候，一定要确保string类型能够转成有效的数据类型，否则最后得到的结果就是按照对应类型的默认值输出
	var s4 string = "golang"
	var b1 bool
	b1, _ = strconv.ParseBool(s4)
	fmt.Printf("b1的类型是：%T,b1=%v \n", b1, b1)

	var s5 string = "golang"
	var num2 int64
	num2, _ = strconv.ParseInt(s5, 10, 64)
	fmt.Printf("num2的类型是：%T,num2=%v \n", num2, num2)
}

// 指针
func test13() {
	age := 18
	fmt.Println(&age)

	//定义一个指针变量：
	//var代表要声明一个变量
	//ptr 指针变量的名字
	//ptr对应的类型是：*int 是一个指针类型 （可以理解为 指向int类型的指针）
	//&age就是一个地址，是ptr变量的具体的值
	var ptr *int = &age
	fmt.Println(ptr)
	fmt.Println("ptr本身这个存储空间的地址为：", &ptr)
	//想获取ptr这个指针或者这个地址指向的那个数据：
	fmt.Printf("ptr指向的数值为：%v", *ptr) //ptr指向的数值为：18
}

// 指针
func test15() {
	var num = 15
	var ptr *int = &num
	*ptr = 20
	fmt.Println(num)
	fmt.Println(ptr)
}
