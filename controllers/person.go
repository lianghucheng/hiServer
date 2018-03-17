package controllers


//个人详情
//@router /user/person/personMsg
func (customer *UserController)PersonMsg()  {
	customer.Data["json"]="personMsg"
	customer.ServeJSON()
}

//打卡详情
//@router /user/person/punchTheClock
func (customer *UserController)UnchTheClock()  {
	customer.Data["json"]="punchTheClock"
	customer.ServeJSON()
}

//tag
//@router /user/person/tag
func (customer *UserController)Tag()  {
	customer.Data["json"]="tag"
	customer.ServeJSON()
}

//成就
//@router /user/person/achieve
func (customer *UserController)Achieve()  {
	customer.Data["json"]="achieve"
	customer.ServeJSON()
}

//奖品记录
//@router /user/person/winningLog
func (customer *UserController)WinningLog()  {
	customer.Data["json"]="winningLog"
	customer.ServeJSON()
}

//反馈
//@router /user/person/feedback
func (customer *UserController)Feedback()  {
	customer.Data["json"]="feedback"
	customer.ServeJSON()
}

//作品情况
//@router /user/person/myPhoto
func (customer *UserController)MyPhoto()  {
	customer.Data["json"]="myPhoto"
	customer.ServeJSON()
}