package base

import (
	"github.com/astaxie/beego/orm"
	"fmt"
	"github.com/astaxie/beego"
	"database/sql"
	"log"
	"hiServer/models"
)

var o orm.Ormer

func init(){
	orm.RegisterModel(
		new(models.WechatUser),
		new(models.User),
		new(models.PersonMsg),
		new(models.PunchTheClock),
		new(models.Tag),
		new(models.AchieveMsg),
		new(models.NoticeMsg),
		new(models.Feedback),
		new(models.Course),
		new(models.Score),
		new(models.StuCard),
		new(models.SelfStudyRoom),
		new(models.Teacher),
		new(models.LostAndFound),
		new(models.Active),
		new(models.JoinMsg),
		new(models.Vote),
		new(models.LuckDrawEvent),
		new(models.LuckDrawItems),
		new(models.LuckDrawDetail),
		new(models.PattedWorks),
		new(models.Secret),
		new(models.Topic),
		new(models.TechQuestion),
		new(models.SchoolOfficialWebMsg),
		new(models.TechNews),
		new(models.Comment),
		new(models.Praise),
	)
	_,dns:=getBase()
	orm.RegisterDataBase("default","mysql",dns)
}

type dbConf struct{
	user string
	pass string
	name string
	host string
	port string
}

func getBase()(dbConf,string){
	dbconf:=new(dbConf)
	dbconf.user=beego.AppConfig.String("db_user")
	dbconf.pass=beego.AppConfig.String("db_pass")
	dbconf.name=beego.AppConfig.String("db_name")
	dbconf.host=beego.AppConfig.String("db_host")
	dbconf.port=beego.AppConfig.String("db_port")
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",dbconf.user,dbconf.pass,dbconf.host,dbconf.port,dbconf.name)
	return *dbconf,dns
}

func Syncdb(force bool){
	dbconf,dns:=getBase()
	createDb(dbconf.name,dns)
	connect()
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

func connect(){
	orm.RegisterDriver("mysql",orm.DRMySQL)
}