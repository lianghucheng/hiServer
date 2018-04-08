package models
//用户的基本数据库（学生、管理员）

const(
	Normal=iota
	Limit
	Frozen
	Student
	Manager
	Yes
	No
	Boy
	Girl
	Unknown
)


type User struct{
	Id int64
	StuNum string
	Password string
	Power int
	Identity int
	WechatUser *WechatUser `orm:"rel(one)"`

	PersonMsg *PersonMsg `orm:"reverse(one)"`
	PunchTheClock []*PunchTheClock `orm:"reverse(many)"`
	Tag []*Tag `orm:"reverse(many)"`
	AchieveMsg []*AchieveMsg `orm:"reverse(many)"`
	NoticeMsg []*NoticeMsg `orm:"reverse(many)"`
	Feedback []*Feedback `orm:"reverse(many)"`


	Course []*Course `orm:"reverse(many)"`
	Score []*Score `orm:"reverse(many)"`

	Active []*Active `orm:"reverse(many)"`
	JoinMsg []*JoinMsg `orm:"reverse(many)"`
	Vote []*Vote `orm:"reverse(many)"`
	LuckDrawDetail []*LuckDrawDetail `orm:"reverse(many)"`
	PattedWorks []*PattedWorks `orm:"reverse(many)"`
	Secret []*Secret `orm:"reverse(many)"`
	Topic []*Topic `orm:"reverse(many)"`
	TechQuestion []*TechQuestion `orm:"reverse(many)"`

	Comment []*Comment `orm:"reverse(many)"`
	Praise []*Praise `orm:"reverse(many)"`
}

