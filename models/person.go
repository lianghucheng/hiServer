package models
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
	User *User `orm:"rel(one)"`
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

//成就奖品记录
type WinningLog struct{
	Id int64
	Item string
	Num int
	Date string
	User *User `orm:"rel(fk)"`
}

//反馈信息
type Feedback struct{
	Id int64
	Content string `orm:"type(text)"`
	Date string
	User *User `orm:"rel(fk)"`
}

