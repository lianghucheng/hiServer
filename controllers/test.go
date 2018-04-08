package controllers

import (
	"hiServer/parser/ccsu"
)

//@router /user/hiserver/engine/test [*]
func (user *UserController)EngineTest(){
	ccsu.Fectch("http://jwcxxcx.ccsu.cn/jwxt/Logon.do?method=logon","B20160304219","134531")
	user.Data["json"]="success"
	user.ServeJSON()
}