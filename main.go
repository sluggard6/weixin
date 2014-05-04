package main

import (
	"database/sql"
	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/astaxie/beego"
	_ "weixin/routers"
)

var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("mysql", "root:111111@tcp(192.168.1.199:3306)/ips?charset=utf8")
	checkErr(err)
}

func main() {
	beego.Trace("Log is ok")
	//stmt, err := db.Prepare("select * from account")
	rows, err := DB.Query("select id from account")
	checkErr(err)
	var id int
	for rows.Next() {
		rows.Scan(&id)
		beego.Trace("======================", id, "===========================")
	}
	rows.Close()
	tx, err := DB.Begin()
	checkErr(err)
	tx.Commit()
	//db.
	DB.Close()
	beego.Run()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
