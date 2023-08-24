package main

import (
	"fmt"
	"io"
	"os"
)

//func main() {
//	test2()
//}

func test1() {
	tty, err := os.OpenFile("E:/a.txt", os.O_RDWR, 0)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	var r io.Reader
	r = tty

	var w io.Writer
	w = r.(io.Writer)

	w.Write([]byte("hello"))
}

type Read interface {
	ReadBook()
}

type Write interface {
	WriteBook()
}

type Book struct {
}

func (book *Book) ReadBook() {
	fmt.Println("读书")
}

func (book *Book) WriteBook() {
	fmt.Println("写字")
}

func test2() {
	book := &Book{}
	// var book2 Book
	book2 := Book{}
	fmt.Println(book2)

	var r Read
	r = book
	r.ReadBook()

	var w Write
	w = r.(Write)
	w.WriteBook()
}
