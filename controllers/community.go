package controllers

//技术咨讯
//@router /user/hiserver/community/techNews
func (customer *UserController)TechNews()  {
	//暂时无需逻辑
	customer.Data["json"]="techNews"
	customer.ServeJSON()
}

//校园地图
//@router /user/hiserver/community/map
func (customer *UserController)Map()  {
	//暂时无需逻辑
	customer.Data["json"]="map"
	customer.ServeJSON()
}

//天气
//@router /user/hiserver/community/weather
func (customer *UserController)Weather(){
	customer.Data["json"]="weather"
	customer.ServeJSON()
}

//失物招领
//@router /user/hiserver/community/lostAndFound
func (customer *UserController)LostAndFound(){
	customer.Data["json"]="lostAndFound"
	customer.ServeJSON()
}

//投票
//@router /user/hiserver/community/vote
func (customer *UserController)Vote(){
	//暂时不做
	customer.Data["json"]="vote"
	customer.ServeJSON()
}

//秘密墙
//@router /user/hiserver/community/secret
func (customer *UserController)Secret(){
	customer.Data["json"]="secret"
	customer.ServeJSON()
}

//近日话题
//@router /user/hiserver/community/topic
func (customer *UserController)Topic(){
	customer.Data["json"]="topic"
	customer.ServeJSON()
}

//技术问答
//@router /user/hiserver/community/techQuestion
func (user *UserController)TechQuestion(){
	user.Data["json"]="techQuestion"
	user.ServeJSON()
}


//通知公告
//@router /user/hiserver/community/notice
func (user *UserController)Notice(){
	user.Data["json"]="notice"
	user.ServeJSON()
}


//随拍
//@router /user/hiserver/community/photograph
func (user *UserController)Photograph(){
	user.Data["json"]="notice"
	user.ServeJSON()
}