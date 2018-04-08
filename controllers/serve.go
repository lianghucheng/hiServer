package controllers


//课表
//@router /user/hiserver/serve/course
func (user *UserController)Course()  {
	user.Data["json"]="course"
	user.ServeJSON()
}

//成绩
//@router /user/hiserver/serve/score
func (user *UserController)Score()  {
	user.Data["json"]="score"
	user.ServeJSON()
}

//校园卡
//@router /user/hiserver/serve/campusCard
func (user *UserController)CampusCard()  {
	user.Data["json"]="score"
	user.ServeJSON()
}

//空自习室
//@router /user/hiserver/serve/selfStudyRoom
func (user *UserController)SelfStudyRoom()  {
	user.Data["json"]="selfStudyRoom"
	user.ServeJSON()
}

//教师查询
//@router /user/hiserver/serve/teacherMsg
func (user *UserController)TeacherMsg()  {
	user.Data["json"]="teacherMsg"
	user.ServeJSON()
}

//普通话、四六级、计算机


//开启推送
//@router /user/hiserver/serve/autoPush
func (user *UserController)AutoPush()  {
	user.Data["json"]="autoPush"
	user.ServeJSON()
}