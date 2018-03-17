package controllers


//课表
//@router /customer/serve/course
func (customer *UserController)Course()  {
	customer.Data["json"]="course"
	customer.ServeJSON()
}

//成绩
//@router /customer/serve/score
func (customer *UserController)Score()  {
	customer.Data["json"]="score"
	customer.ServeJSON()
}

//校园卡
//@router /user/serve/campusCard
func (customer *UserController)CampusCard()  {
	customer.Data["json"]="score"
	customer.ServeJSON()
}

//空自习室
//@router /user/serve/selfStudyRoom
func (customer *UserController)SelfStudyRoom()  {
	customer.Data["json"]="selfStudyRoom"
	customer.ServeJSON()
}

//教师查询
//@router /user/serve/teacherMsg
func (customer *UserController)TeacherMsg()  {
	customer.Data["json"]="teacherMsg"
	customer.ServeJSON()
}

//普通话、四六级、计算机


//开启推送
//@router /user/serve/autoPush
func (customer *UserController)AutoPush()  {
	customer.Data["json"]="autoPush"
	customer.ServeJSON()
}