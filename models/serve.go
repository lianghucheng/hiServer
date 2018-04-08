package models

import (
	"github.com/astaxie/beego/orm"
)

//服务模块所需数据库

//课程
type Course struct{
	Id int64
	Order string
	Weekday string
	Name string
	Class string
	Teacher string
	ExistWeek string
	Position string
	Date string
	User *User `orm:"rel(fk)"`
}

//成绩
type Score struct{
	Id int64
	CourseCharacter string  //课程性质
	CourseKind string  //课程种类
	CourseName string  //课程名称
	Credit string  //学分
	ConductTerm string  //开课时间
	ExamCharacter string  //考试性质
	StuName string  //学生姓名
	StuNum string  //学生学号
	CourseNum string  //课程序号
	Score string  //成绩
	Period string  //学时
	User *User `orm:"rel(fk)"`
}

//校园卡
type StuCard struct{
	Id int64
	customer *User
}

//空教室
type SelfStudyRoom struct{
	Id int64
	TimeSlot string //时间段
	Position string //地点
}

//教师信息
type Teacher struct{
	Id int64
	Name string
	Massage string
}

func (course *Course)Insert()(int64,error){
	o:=orm.NewOrm()
	return o.Insert(course)
}


func (course *Course)Updata()(int64,error){
	o:=orm.NewOrm()
	return o.Update(course)
}

func (course *Course)Delete()(int64,error){
	o:=orm.NewOrm()
	return o.Delete(course)
}