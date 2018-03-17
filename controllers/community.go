package controllers

//技术咨讯
//@router /customer/community/techNews
func (customer *CustomerController)TechNews()  {
	//暂时无需逻辑
	customer.Data["json"]="techNews"
	customer.ServeJSON()
}

//校园地图
//@router /customer/community/map
func (customer *CustomerController)Map()  {
	//暂时无需逻辑
	customer.Data["json"]="map"
	customer.ServeJSON()
}

//天气
//@router /customer/community/weather
func (customer *CustomerController)Weather(){
	customer.Data["json"]="weather"
	customer.ServeJSON()
}

//失物招领
//@router /customer/community/lostAndFound
func (customer *CustomerController)LostAndFound(){
	customer.Data["json"]="lostAndFound"
	customer.ServeJSON()
}

//投票
//@router /customer/community/vote
func (customer *CustomerController)Vote(){
	//暂时不做
	customer.Data["json"]="vote"
	customer.ServeJSON()
}

//秘密墙
//@router /customer/community/secret
func (customer *CustomerController)Secret(){
	customer.Data["json"]="secret"
	customer.ServeJSON()
}

//近日话题
//@router /customer/community/topic
func (customer *CustomerController)Topic(){
	customer.Data["json"]="topic"
	customer.ServeJSON()
}

//技术问答
//@router /customer/community/techQuestion
func (customer *CustomerController)TechQuestion(){
	customer.Data["json"]="techQuestion"
	customer.ServeJSON()
}


//通知公告
//@router /customer/community/notice
func (customer *CustomerController)Notice(){
	customer.Data["json"]="notice"
	customer.ServeJSON()
}


//随拍
//@router /customer/community/photograph
func (customer *CustomerController)Photograph(){
	customer.Data["json"]="notice"
	customer.ServeJSON()
}