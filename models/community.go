package models
//广场模块所需数据库


//失物招领
type LostAndFound struct{
	Id int64
	customer *User
	Date string
}

//活动
type Active struct{
	Id int64
	Name string
	StartTime string
	EndTime string
	Initiator string
	Introduce string  `orm:"size(255)"`
	JoinMsg []*JoinMsg `orm:"reverse(many)"`
	Admin *User `orm:"rel(fk)"`
	LuckDrawEvent []*LuckDrawEvent `orm:"reverse(many)"`
	Date string
}

//参加人员信息		主动或被动
type JoinMsg struct{
	Id int64
	//阐述
	//照片
	IsCome int
	Active *Active `orm:"rel(fk)"`
	User *User `orm:"rel(fk)"`
	Vote []*Vote `orm:"reverse(many)"`
	Date string
}

//签到是一种逻辑


//投票
type Vote struct{
	Id int
	Max int  //最大可投票数
	Voter *User `orm:"rel(fk)"`
	Elector *JoinMsg `orm:"rel(fk)"`
	Date string
}

//抽奖事件
type LuckDrawEvent struct{
	Id int
	Name string
	Introduce string `orm:"size(255)"`
	Active *Active `orm:"rel(fk)"`
	LuckDrawItems []*LuckDrawItems `orm:"reverse(many)"`
	Date string
}

//奖项
type LuckDrawItems struct{
	Id int
	Name string
	LuckDrawEvent *LuckDrawEvent `orm:"rel(fk)"`
	Num int
	PictureUrl string
	Introduce string `orm:"size(255)"`
	LuckDrawDetail []*LuckDrawDetail `orm:"reverse(many)"`
	Date string
}

//中奖详情
type LuckDrawDetail struct{
	Id int
	LuckDrawItems *LuckDrawItems `orm:"rel(fk)"`
	Winners *User `orm:"rel(fk)"`
	Date string
}

//随拍
type PattedWorks struct {//随拍
	Id int
	WritingUrl string
	Author *User `orm:"rel(fk)"`
	Comment []*Comment `orm:"reverse(many)"`
	Praise []*Praise `orm:"reverse(many)"`
	PattedWorksDraw *PattedWorksDraw `orm:"reverse(one)"`
	Date string
}

//随拍奖励
type PattedWorksDraw struct {
	Id int
	PattedWorks *PattedWorks `orm:"rel(one)"`
	Item string
	Num int
	Date string
}


//秘密墙
type Secret struct{
	Id int64
	Anonymous int //匿名标识
	Content string `orm:"type(text)"`
	Author *User `orm:"rel(fk)"`
	Comment []*Comment `orm:"reverse(many)"`
	Praise []*Praise `orm:"reverse(many)"`
	Date string
}

//近日话题
type Topic struct{
	Id int64
	Content string `orm:"type(text)"`
	Author *User `orm:"rel(fk)"`
	Comment []*Comment `orm:"reverse(many)"`
	Praise []*Praise `orm:"reverse(many)"`
	Date string
}


//技术问答
type TechQuestion struct{
	Id int64
	Content string `orm:"type(text)"`
	Author *User `orm:"rel(fk)"`
	Comment []*Comment `orm:"reverse(many)"`
	Praise []*Praise `orm:"reverse(many)"`
	Date string
}

//通知公告，长大要闻，学术活动，学工动态
type SchoolOfficialWebMsg struct {
	Id int
	Content string `orm:"type(text)"`
	Date string
}

//(置顶图文推送的)技术咨询
type TechNews struct{
	Id int64
	Content string `orm:"type(text)"`
	Date string
}


type Comment struct{
	Id int
	Content string `orm:"type(text)"`
	Commentator *User `orm:"rel(fk)"`
	PattedWorks *PattedWorks `orm:"null;rel(fk)"`
	Secret *Secret `orm:"null;rel(fk)"`
	Topic *Topic `orm:"null;rel(fk)"`
	TechQuestion *TechQuestion `orm:"null;rel(fk)"`
	Date string
}

type Praise struct{
	Id int
	Description int  //点赞类型
	Praiser *User `orm:"rel(fk)"`
	PattedWorks *PattedWorks `orm:"null;rel(fk)"`
	Secret *Secret `orm:"null;rel(fk)"`
	Topic *Topic `orm:"null;rel(fk)"`
	TechQuestion *TechQuestion `orm:"null;rel(fk)"`
	Date string
}