package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["hiServer/controllers:UserController"] = append(beego.GlobalControllerRouter["hiServer/controllers:UserController"],
		beego.ControllerComments{
			Method: "Achieve",
			Router: `/user/hiserver//achieve`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiServer/controllers:UserController"] = append(beego.GlobalControllerRouter["hiServer/controllers:UserController"],
		beego.ControllerComments{
			Method: "LostAndFound",
			Router: `/user/hiserver/community/lostAndFound`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiServer/controllers:UserController"] = append(beego.GlobalControllerRouter["hiServer/controllers:UserController"],
		beego.ControllerComments{
			Method: "Map",
			Router: `/user/hiserver/community/map`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiServer/controllers:UserController"] = append(beego.GlobalControllerRouter["hiServer/controllers:UserController"],
		beego.ControllerComments{
			Method: "Notice",
			Router: `/user/hiserver/community/notice`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiServer/controllers:UserController"] = append(beego.GlobalControllerRouter["hiServer/controllers:UserController"],
		beego.ControllerComments{
			Method: "Photograph",
			Router: `/user/hiserver/community/photograph`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiServer/controllers:UserController"] = append(beego.GlobalControllerRouter["hiServer/controllers:UserController"],
		beego.ControllerComments{
			Method: "Secret",
			Router: `/user/hiserver/community/secret`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiServer/controllers:UserController"] = append(beego.GlobalControllerRouter["hiServer/controllers:UserController"],
		beego.ControllerComments{
			Method: "TechNews",
			Router: `/user/hiserver/community/techNews`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiServer/controllers:UserController"] = append(beego.GlobalControllerRouter["hiServer/controllers:UserController"],
		beego.ControllerComments{
			Method: "TechQuestion",
			Router: `/user/hiserver/community/techQuestion`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiServer/controllers:UserController"] = append(beego.GlobalControllerRouter["hiServer/controllers:UserController"],
		beego.ControllerComments{
			Method: "Topic",
			Router: `/user/hiserver/community/topic`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiServer/controllers:UserController"] = append(beego.GlobalControllerRouter["hiServer/controllers:UserController"],
		beego.ControllerComments{
			Method: "Vote",
			Router: `/user/hiserver/community/vote`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiServer/controllers:UserController"] = append(beego.GlobalControllerRouter["hiServer/controllers:UserController"],
		beego.ControllerComments{
			Method: "Weather",
			Router: `/user/hiserver/community/weather`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiServer/controllers:UserController"] = append(beego.GlobalControllerRouter["hiServer/controllers:UserController"],
		beego.ControllerComments{
			Method: "EngineTest",
			Router: `/user/hiserver/engine/test`,
			AllowHTTPMethods: []string{"*"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiServer/controllers:UserController"] = append(beego.GlobalControllerRouter["hiServer/controllers:UserController"],
		beego.ControllerComments{
			Method: "Feedback",
			Router: `/user/hiserver/hiserver/person/feedback`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiServer/controllers:UserController"] = append(beego.GlobalControllerRouter["hiServer/controllers:UserController"],
		beego.ControllerComments{
			Method: "Bind",
			Router: `/user/hiserver/person/bind`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiServer/controllers:UserController"] = append(beego.GlobalControllerRouter["hiServer/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetOpenid",
			Router: `/user/hiserver/person/getopenid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiServer/controllers:UserController"] = append(beego.GlobalControllerRouter["hiServer/controllers:UserController"],
		beego.ControllerComments{
			Method: "MyPhoto",
			Router: `/user/hiserver/person/myPhoto`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiServer/controllers:UserController"] = append(beego.GlobalControllerRouter["hiServer/controllers:UserController"],
		beego.ControllerComments{
			Method: "PersonMsg",
			Router: `/user/hiserver/person/personMsg`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiServer/controllers:UserController"] = append(beego.GlobalControllerRouter["hiServer/controllers:UserController"],
		beego.ControllerComments{
			Method: "UnchTheClock",
			Router: `/user/hiserver/person/punchTheClock`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiServer/controllers:UserController"] = append(beego.GlobalControllerRouter["hiServer/controllers:UserController"],
		beego.ControllerComments{
			Method: "Tag",
			Router: `/user/hiserver/person/tag`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiServer/controllers:UserController"] = append(beego.GlobalControllerRouter["hiServer/controllers:UserController"],
		beego.ControllerComments{
			Method: "WinningLog",
			Router: `/user/hiserver/person/winningLog`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiServer/controllers:UserController"] = append(beego.GlobalControllerRouter["hiServer/controllers:UserController"],
		beego.ControllerComments{
			Method: "AutoPush",
			Router: `/user/hiserver/serve/autoPush`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiServer/controllers:UserController"] = append(beego.GlobalControllerRouter["hiServer/controllers:UserController"],
		beego.ControllerComments{
			Method: "CampusCard",
			Router: `/user/hiserver/serve/campusCard`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiServer/controllers:UserController"] = append(beego.GlobalControllerRouter["hiServer/controllers:UserController"],
		beego.ControllerComments{
			Method: "Course",
			Router: `/user/hiserver/serve/course`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiServer/controllers:UserController"] = append(beego.GlobalControllerRouter["hiServer/controllers:UserController"],
		beego.ControllerComments{
			Method: "Score",
			Router: `/user/hiserver/serve/score`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiServer/controllers:UserController"] = append(beego.GlobalControllerRouter["hiServer/controllers:UserController"],
		beego.ControllerComments{
			Method: "SelfStudyRoom",
			Router: `/user/hiserver/serve/selfStudyRoom`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiServer/controllers:UserController"] = append(beego.GlobalControllerRouter["hiServer/controllers:UserController"],
		beego.ControllerComments{
			Method: "TeacherMsg",
			Router: `/user/hiserver/serve/teacherMsg`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
