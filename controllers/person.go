package controllers


//个人详情
//@router /user/person/personMsg
func (user *UserController)PersonMsg()  {
	user.Data["json"]="personMsg"
	user.ServeJSON()
}

//打卡详情
//@router /user/person/punchTheClock
func (user *UserController)UnchTheClock()  {
	user.Data["json"]="punchTheClock"
	user.ServeJSON()
}

//tag
//@router /user/person/tag
func (user *UserController)Tag()  {
	user.Data["json"]="tag"
	user.ServeJSON()
}

//成就
//@router /user/person/achieve
func (user *UserController)Achieve()  {
	user.Data["json"]="achieve"
	user.ServeJSON()
}

//奖品记录
//@router /user/person/winningLog
func (user *UserController)WinningLog()  {
	user.Data["json"]="winningLog"
	user.ServeJSON()
}

//反馈
//@router /user/person/feedback
func (user *UserController)Feedback()  {
	user.Data["json"]="feedback"
	user.ServeJSON()
}

//作品情况
//@router /user/person/myPhoto
func (user *UserController)MyPhoto()  {
	user.Data["json"]="myPhoto"
	user.ServeJSON()
}