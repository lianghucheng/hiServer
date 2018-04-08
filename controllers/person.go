package controllers

import (
	"github.com/astaxie/beego"
	mpoauth2 "gopkg.in/chanxuehong/wechat.v2/mp/oauth2"
	"gopkg.in/chanxuehong/wechat.v2/oauth2"
	"github.com/chanxuehong/rand"
	"github.com/astaxie/beego/orm"
	"hiServer/models"
)

const (
	wxAppId           = "wxebc1d9dbfbd550d4"                           // 填上自己的参数
	wxAppSecret       = "8153e67bdcc17c325f628d0a029b6608"                       // 填上自己的参数
	oauth2RedirectURI = "https://www.lianghucheng.top/user/hiserver/person/getopenid" // 填上自己的参数
	oauth2Scope       = "snsapi_userinfo"                 // 填上自己的参数
)

var (
	oauth2Endpoint oauth2.Endpoint = mpoauth2.NewEndpoint(wxAppId, wxAppSecret)
)

//获取openid
//@router /user/hiserver/person/getopenid
func (user *UserController)GetOpenid(){
	var token *oauth2.Token

	c:=user.Ctx

	state := string(rand.NewHex())

	oauth2Client := oauth2.Client{
		Endpoint: oauth2Endpoint,
	}
	code:=c.Input.Query("code")
	beego.Debug(code)
	if code==""{
		AuthCodeURL := mpoauth2.AuthCodeURL(wxAppId, oauth2RedirectURI, oauth2Scope, state)
		beego.Debug(AuthCodeURL)
		c.Redirect(302,AuthCodeURL)
	}else{
		token_, err := oauth2Client.ExchangeToken(code)
		if err != nil {
			beego.Error(err)
			return
		}
		beego.Debug(token_)
		token=token_
		beego.Debug(token.OpenId)
		o:=orm.NewOrm()
		if i,_:=o.QueryTable("wechat_user").Filter("open_id",token.OpenId).Count();i==0{
			wechatUser:=models.WechatUser{
				OpenId:token.OpenId,
			}
			beego.Debug(wechatUser.Insert())
		}
		i,_:=o.QueryTable("WechatUser").Filter("open_id",token.OpenId).Count()
		beego.Debug(i)
		user.Data["json"]="success"
		user.ServeJSON()
	}
}

//绑定学号
//@router /user/hiserver/person/bind [get]
func (user *UserController)Bind(){
	var token *oauth2.Token

	c:=user.Ctx

	state := string(rand.NewHex())

	oauth2Client := oauth2.Client{
		Endpoint: oauth2Endpoint,
	}
	code:=c.Input.Query("code")
	beego.Debug(code)
	if code==""{
		AuthCodeURL := mpoauth2.AuthCodeURL(wxAppId, oauth2RedirectURI, oauth2Scope, state)
		beego.Debug(AuthCodeURL)
		c.Redirect(302,AuthCodeURL)
	}else{
		token_, err := oauth2Client.ExchangeToken(code)
		if err != nil {
			beego.Error(err)
			return
		}
		beego.Debug(token_)
		token=token_
	}

	beego.Debug(token.OpenId)
	stuNum:=user.GetString("stuNum")
	password:=user.GetString("password")
	beego .Debug(stuNum,password)

	user.Data["json"]="bind"
	user.ServeJSON()
}

//个人详情
//@router /user/hiserver/person/personMsg [post]
func (user *UserController)PersonMsg()  {
	stuName:=user.GetString("stuName")
	nick:=user.GetString("phone")
	sex:=user.GetString("sex")
	phone:=user.GetString("phone")
	age:=user.GetString("age")
	constellation:=user.GetString("constellation")
	address:=user.GetString("address")
	personalitySignature:=user.GetString("personalitySignature")
	birth:=user.GetString("birth")
	email:=user.GetString("email")
	beego.Debug(stuName,nick,sex,phone,age,constellation,address,personalitySignature,birth,email)

	user.Data["json"]="personMsg"
	user.ServeJSON()
}
//打卡详情
//@router /user/hiserver/person/punchTheClock
func (user *UserController)UnchTheClock()  {
	user.Data["json"]="punchTheClock"
	user.ServeJSON()
}

//tag
//@router /user/hiserver/person/tag
func (user *UserController)Tag()  {
	user.Data["json"]="tag"
	user.ServeJSON()
}

//成就
//@router /user/hiserver//achieve
func (user *UserController)Achieve()  {
	user.Data["json"]="achieve"
	user.ServeJSON()
}

//奖品记录
//@router /user/hiserver/person/winningLog
func (user *UserController)WinningLog()  {
	user.Data["json"]="winningLog"
	user.ServeJSON()
}

//反馈
//@router /user/hiserver/hiserver/person/feedback
func (user *UserController)Feedback()  {
	user.Data["json"]="feedback"
	user.ServeJSON()
}

//作品情况
//@router /user/hiserver/person/myPhoto
func (user *UserController)MyPhoto()  {
	user.Data["json"]="myPhoto"
	user.ServeJSON()
}