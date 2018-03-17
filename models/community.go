package models
//广场模块所需数据库

//失物招领
type LostAndFound struct{
	Id int64
	customer *Customer
}

//活动
type Active struct{
	Id int64
}

//活动详情
type ActiveDetail struct{
	Id int64
	active *Active
}

//秘密墙
type Secret struct{
	Id int64
	customer *Customer
}

//秘密评论
type SecretComment struct{
	Id int64
	customer *Customer
	secret *Secret
}

//秘密点赞
type SecretPraise struct {
	Id int64
	customer *Customer
	photograph *Photograph
}

//近日话题
type Topic struct{
	Id int64
	customer *Customer
}

//近日话题评论
type TopicComment struct{
	Id int64
	customer *Customer
	Topic *Topic
}

//近日话题点赞
type TopicPraise struct {
	Id int64
	customer *Customer
	photograph *Photograph
}

//技术问答
type TechQuestion struct{
	Id int64
}

//技术问答评论
type TechQuestionComment struct{
	Id int64
	customer *Customer
	techQuestion *TechQuestion
}

//技术问答点赞
type TechQuestionPraise struct {
	Id int64
	customer *Customer
	photograph *Photograph
}

//随拍作品
type Photograph struct{
	Id int64
	customer *Customer
}

//随拍作品评论
type PhotographComment struct{
	Id int64
	customer *Customer
	photograph *Photograph
}


//随拍作品点赞
type PhotographPraise struct {
	Id int64
	customer *Customer
	photograph *Photograph
}


//(置顶图文推送的)技术咨询
type TechNews struct{
	Id int64
}

