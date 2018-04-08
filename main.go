package main

import (
	_ "hiServer/routers"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
	"os"
	"hiServer/models/base"
)

func init(){
	initArgs()
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.SetStaticPath("/user/hiserver/test_pages","test_pages")
	beego.Run()
}

func initArgs() {
	args := os.Args
	for _, v := range args {
		if v == "-syncdb" {
			base.Syncdb(false)
			os.Exit(0)
		}
		if v == "-syncdb-force" {
			base.Syncdb(true)
			os.Exit(0)
		}
	}
}