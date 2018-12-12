package main

import (
	_ "app/cmd"
	"app/config"

	_ "github.com/go-sql-driver/mysql"

	"github.com/ilibs/gosql"
	"github.com/ilibs/very/server"
)

func main() {
	//db connect
	gosql.Connect(config.App.DB)

	//command server
	cliServ := server.NewCliServer()
	cliServ.Run()
}
