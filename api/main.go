package main

import (
	_ "app/cmd"
	"app/config"

	_ "github.com/go-sql-driver/mysql"

	"github.com/goapt/gee"
	"github.com/ilibs/gosql"
)

func main() {
	//db connect
	gosql.Connect(config.App.DB)

	//command server
	cliServ := gee.NewCliServer()
	cliServ.Run()
}
