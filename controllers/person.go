package controllers


//个人详情
//@router /customer/person/personMsg
func (customer *CustomerController)PersonMsg()  {
	customer.Data["json"]="personMsg"
	customer.ServeJSON()
}

//打卡详情
//@router /customer/person/punchTheClock
func (customer *CustomerController)UnchTheClock()  {
	customer.Data["json"]="punchTheClock"
	customer.ServeJSON()
}

//tag
//@router /customer/person/tag
func (customer *CustomerController)Tag()  {
	customer.Data["json"]="tag"
	customer.ServeJSON()
}

//成就
//@router /customer/person/achieve
func (customer *CustomerController)Achieve()  {
	customer.Data["json"]="achieve"
	customer.ServeJSON()
}

//奖品记录
//@router /customer/person/winningLog
func (customer *CustomerController)WinningLog()  {
	customer.Data["json"]="winningLog"
	customer.ServeJSON()
}

//反馈
//@router /customer/person/feedback
func (customer *CustomerController)Feedback()  {
	customer.Data["json"]="feedback"
	customer.ServeJSON()
}

//作品情况
//@router /customer/person/myPhoto
func (customer *CustomerController)MyPhoto()  {
	customer.Data["json"]="myPhoto"
	customer.ServeJSON()
}