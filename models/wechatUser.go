package models

import "github.com/astaxie/beego/orm"

//微信的基本信息

type WechatUser struct{
	Id int64
	OpenId string
	User *User `orm:"reverse(one)"`
}

func (w *WechatUser)Insert()(int64,error){
	o:=orm.NewOrm()
	return o.Insert(w)
}