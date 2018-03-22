package models
//微信的基本信息

type WechatUser struct{
	Id int64
	OpenId string
	User *User `orm:"reverse(one)"`
}