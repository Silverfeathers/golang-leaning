package main

import "fmt"

func main() {
	test7()
}

// 面向对象
func test1() {
	var book Book
	book.title = "golang"
	book.auth = "张三"
	fmt.Println(book)

	changeTitle(book)
	fmt.Println(book)

	changeTitle2(&book)
	fmt.Println(book)
}

type Book struct {
	title string
	auth  string
}

func changeTitle(book Book) {
	book.title = "java"
}

func changeTitle2(book *Book) {
	book.title = "python"
}

// 面向对象: 封装
func test2() {
	hero := Hero{Name: "zhang3", Level: 4}
	fmt.Println(hero)
	fmt.Println(hero.GetName())

	hero.SetName("lisi")
	fmt.Println(hero)
	fmt.Println(hero.GetName())
}

type Hero struct {
	Name  string
	Level int
}

func (this *Hero) SetName(name string) {
	this.Name = name
}

func (this *Hero) GetName() string {
	return this.Name
}

// 继承
func test3() {
	var superHero SuperHero
	superHero.SetName("wang5")
	superHero.Power = 99
	fmt.Println(superHero)
}

type SuperHero struct {
	Hero
	Power int
}

func (this *SuperHero) SetName(name string) {
	this.Name = "super " + name
}

func (this *SuperHero) Fly() {
	fmt.Println("SuperHero Fly")
}

// 多态
func test5() {
	var dongwu Dongwu
	dongwu = &Cat{"blue"}
	dongwu.Sleep()
}

type Dongwu interface {
	Sleep()
	GetColor() string
}

type Cat struct {
	color string
}

func (cat *Cat) Sleep() {
	fmt.Println("cat is sleep...")
}
func (cat *Cat) GetColor() string {
	return cat.color
}

// interface
func test7() {
	books := Books{"lisi"}
	myInter(books)
	myInter("abc")
}

func myInter(arg interface{}) {
	fmt.Println("myInter is calling")
	fmt.Println(arg)
	v, ok := arg.(string)
	//fmt.Printf("%T", arg)
	//fmt.Printf("%T", v)
	if ok {
		fmt.Println("Yes")
	} else {
		fmt.Println(v)
		fmt.Printf("%T", v)
		fmt.Println()
	}
}

type Books struct {
	auth string
}
