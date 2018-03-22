package base

import (
	"github.com/astaxie/beego/orm"
	"fmt"
	"github.com/astaxie/beego"
	"database/sql"
	"log"
)

var o orm.Ormer

func init(){
	orm.RegisterModel(

	)
}

type dbConf struct{
	user string
	pass string
	name string
	host string
	port string
}

func Syncdb(force bool){
	dbconf:=new(dbConf)
	dbconf.user=beego.AppConfig.String("db_user")
	dbconf.pass=beego.AppConfig.String("db_pass")
	dbconf.name=beego.AppConfig.String("db_name")
	dbconf.host=beego.AppConfig.String("db_host")
	dbconf.port=beego.AppConfig.String("db_port")
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",dbconf.user,dbconf.pass,dbconf.host,dbconf.port,dbconf.name)
	createDb(dbconf.name,dns)
	connect(dns)
	o=orm.NewOrm()
	name:="default"
	verbose:=true

	err := orm.RunSyncdb(name,force,verbose)
	if err!=nil{
		fmt.Println(err)
	}
	if force{
		fmt.Println("database init is complete.\nPlease restart the application")
	}
}

func createDb(db_name string,dns string){
	var sqlstring string
	//sql ="create database beego_test;"
	beego.Debug("bbbb")
	db,err:=sql.Open("mysql",dns)
	if err!=nil{
		panic(err.Error())
	}
	r,err:=db.Exec(sqlstring)
	if err!=nil{
		beego.Debug(err)
		beego.Debug(r)
	}else{
		log.Println("Database",db_name,"created")
	}
	db.Close()
}

func connect(dns string){
	orm.RegisterDriver("mysql",orm.DRMySQL)
	orm.RegisterDataBase("default","mysql",dns)
}