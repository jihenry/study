package server

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var DBEngine *xorm.Engine

func InitDB() {
	engine, err := xorm.NewEngine("mysql", "root:123456@tcp(127.0.0.1:3306)/mysql")
	if err != nil {
		fmt.Printf("xorm.NewEngine err:%s\n", err)
		return
	}
	engine.ShowSQL()
	err = engine.Ping()
	if err != nil {
		fmt.Printf("Ping err:%s\n", err)
		return
	}
	fmt.Printf("mysql connected\n")
	engine.Close()
}
