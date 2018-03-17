package controllers


//课表
//@router /customer/serve/course
func (customer *CustomerController)Course()  {
	customer.Data["json"]="course"
	customer.ServeJSON()
}

//成绩
//@router /customer/serve/score
func (customer *CustomerController)Score()  {
	customer.Data["json"]="score"
	customer.ServeJSON()
}

//校园卡
//@router /customer/serve/campusCard
func (customer *CustomerController)CampusCard()  {
	customer.Data["json"]="score"
	customer.ServeJSON()
}

//空自习室
//@router /customer/serve/selfStudyRoom
func (customer *CustomerController)SelfStudyRoom()  {
	customer.Data["json"]="selfStudyRoom"
	customer.ServeJSON()
}

//教师查询
//@router /customer/serve/teacherMsg
func (customer *CustomerController)TeacherMsg()  {
	customer.Data["json"]="teacherMsg"
	customer.ServeJSON()
}

//普通话、四六级、计算机


//开启推送
//@router /customer/serve/autoPush
func (customer *CustomerController)AutoPush()  {
	customer.Data["json"]="autoPush"
	customer.ServeJSON()
}