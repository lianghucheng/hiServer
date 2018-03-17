package controllers

import "github.com/astaxie/beego"

type BaseController struct {
	beego.Controller
}

type UserController struct {
	BaseController
}

type ManagerController struct {
	BaseController
}
func (this *BaseController) ReturnJson(status int,message string,args ...interface{}){
	result:=make(map[string]interface{})
	result["status"]=status
	result["message"]=message

	key:=""

	for _,arg:=range args{
		switch arg.(type){
		case string:
			key=arg.(string)
		default:
			result[key]=arg
		}
	}
	this.Data["json"]=result
	this.ServeJSON()
}

func (this *BaseController) ReturnSuccess(args ...interface{}) {

	result := make(map[string]interface{})
	result["status"] = 10000
	result["message"] = "success"
	key := ""
	for _, arg := range args {
		switch arg.(type) {
		case string:
			key = arg.(string)
		default:
			result[key] = arg
		}
	}
	this.Data["json"] = result
	this.ServeJSON()

}