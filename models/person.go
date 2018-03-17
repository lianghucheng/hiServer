package models
//个人模块所需数据库

//个人详情
type PersonMsg struct{
	Id int64
	customer *Customer
}

//打卡记录
type PunchTheClock struct{
	Id int64
	customer *Customer
}

//个人标签
type Tag struct{
	Id int64
	customer *Customer
}

//个人成就
type Achieve struct{
	Id int64
	customer *Customer
}

//奖品记录
type WinningLog struct{
	Id int64
	customer *Customer
}

//反馈信息
type Feedback struct{
	Id int64
	customer *Customer
}

//个人作品信息
type MyPhoto struct{
	Id int64
	customer *Customer
}