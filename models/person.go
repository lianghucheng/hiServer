package models

import (
	"github.com/astaxie/beego/orm"
)

//个人模块所需数据库

//个人详情
type PersonMsg struct{
	Id int64
	StuName string `orm:"null"`
	Nick string `orm:"null"`
	Sex int `orm:"null"`
	Phone string `orm:"null"`
	Age string `orm:"null"`
	Constellation string `orm:"null"`
	Address string `orm:"null"`
	PersonalitySignature string `orm:"null"`
	Birth string `orm:"null"`
	Email string `orm:"null"`
	PunchTheClockNum string `orm:"null"`
	Active int `orm:"null"`
	Date string `orm:"null"`
	User *User `orm:"rel(one)"`
}

//打卡记录
type PunchTheClock struct{
	Id int64
	Date string
	User *User `orm:"rel(fk)"`
}

//个人标签
type Tag struct{
	Id int64
	content string
	Date string
	User *User `orm:"rel(fk)"`
}

//个人成就详情
type AchieveMsg struct{
	Id int64
	Content string `orm:"null"`
	Date string
	User *User `orm:"rel(fk)"`
}

//消息通知
type NoticeMsg struct{
	Id int
	Content string `orm:"type(text)"`
	User *User `orm:"rel(fk)"`
}


//反馈信息
type Feedback struct{
	Id int64
	Content string `orm:"type(text)"`
	Date string
	User *User `orm:"rel(fk)"`
}

func (p *PersonMsg)PerosonMsgInsert()(int64,error){
	o:=orm.NewOrm()
	return o.Insert(p)
}