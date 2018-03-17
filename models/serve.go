package models
//服务模块所需数据库

//课表
type Course struct{
	Id int64
	customer *Customer
}

//成绩
type Score struct{
	Id int64
	customer *Customer
}

//校园卡
type StuCard struct{
	Id int64
	customer *Customer
}

//空教室
type SelfStudyRoom struct{
	Id int64
}

//教师信息
type Teacher struct{
	Id int64
}

