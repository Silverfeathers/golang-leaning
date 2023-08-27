package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func testConnMysql() {
	// 数据库连接信息
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println("数据库连接错误:", err)
		return
	}
	defer db.Close()

	// 测试数据库连接
	err = db.Ping()
	if err != nil {
		fmt.Println("数据库连接错误:", err)
		return
	}

	fmt.Println("成功连接到数据库！")

	selectAllTest1(db)
}

func insertTest1(db *sql.DB) {
	sqlStr := "insert into test1 (name, num) values (?, ?)"
	res, err := db.Exec(sqlStr, "abc", 123)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("res:", res)
}

type Test1Table struct {
	id   int
	name string
	num  int
}

func selectOneTest1(db *sql.DB) {
	var test1Table Test1Table
	sqlStr := "select * from test1 where id = ?"
	err := db.QueryRow(sqlStr, 1).Scan(&test1Table.id, &test1Table.name, &test1Table.num)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(test1Table)
	}
}

func selectAllTest1(db *sql.DB) {
	sqlStr := "select * from test1"
	rows, err := db.Query(sqlStr)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer rows.Close()

	var test1Table Test1Table
	for rows.Next() {
		rows.Scan(&test1Table.id, &test1Table.name, &test1Table.num)
		fmt.Println(test1Table)
	}
}
