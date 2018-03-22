package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["hiServer/controllers:UserController"] = append(beego.GlobalControllerRouter["hiServer/controllers:UserController"],
		beego.ControllerComments{
			Method: "EngineTest",
			Router: `/engine/test`,
			AllowHTTPMethods: []string{"*"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiServer/controllers:UserController"] = append(beego.GlobalControllerRouter["hiServer/controllers:UserController"],
		beego.ControllerComments{
			Method: "LostAndFound",
			Router: `/user/community/lostAndFound`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiServer/controllers:UserController"] = append(beego.GlobalControllerRouter["hiServer/controllers:UserController"],
		beego.ControllerComments{
			Method: "Map",
			Router: `/user/community/map`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiServer/controllers:UserController"] = append(beego.GlobalControllerRouter["hiServer/controllers:UserController"],
		beego.ControllerComments{
			Method: "Notice",
			Router: `/user/community/notice`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiServer/controllers:UserController"] = append(beego.GlobalControllerRouter["hiServer/controllers:UserController"],
		beego.ControllerComments{
			Method: "Photograph",
			Router: `/user/community/photograph`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiServer/controllers:UserController"] = append(beego.GlobalControllerRouter["hiServer/controllers:UserController"],
		beego.ControllerComments{
			Method: "Secret",
			Router: `/user/community/secret`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiServer/controllers:UserController"] = append(beego.GlobalControllerRouter["hiServer/controllers:UserController"],
		beego.ControllerComments{
			Method: "TechNews",
			Router: `/user/community/techNews`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiServer/controllers:UserController"] = append(beego.GlobalControllerRouter["hiServer/controllers:UserController"],
		beego.ControllerComments{
			Method: "TechQuestion",
			Router: `/user/community/techQuestion`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiServer/controllers:UserController"] = append(beego.GlobalControllerRouter["hiServer/controllers:UserController"],
		beego.ControllerComments{
			Method: "Topic",
			Router: `/user/community/topic`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiServer/controllers:UserController"] = append(beego.GlobalControllerRouter["hiServer/controllers:UserController"],
		beego.ControllerComments{
			Method: "Vote",
			Router: `/user/community/vote`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiServer/controllers:UserController"] = append(beego.GlobalControllerRouter["hiServer/controllers:UserController"],
		beego.ControllerComments{
			Method: "Weather",
			Router: `/user/community/weather`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiServer/controllers:UserController"] = append(beego.GlobalControllerRouter["hiServer/controllers:UserController"],
		beego.ControllerComments{
			Method: "Achieve",
			Router: `/user/person/achieve`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiServer/controllers:UserController"] = append(beego.GlobalControllerRouter["hiServer/controllers:UserController"],
		beego.ControllerComments{
			Method: "Feedback",
			Router: `/user/person/feedback`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiServer/controllers:UserController"] = append(beego.GlobalControllerRouter["hiServer/controllers:UserController"],
		beego.ControllerComments{
			Method: "MyPhoto",
			Router: `/user/person/myPhoto`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiServer/controllers:UserController"] = append(beego.GlobalControllerRouter["hiServer/controllers:UserController"],
		beego.ControllerComments{
			Method: "PersonMsg",
			Router: `/user/person/personMsg`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiServer/controllers:UserController"] = append(beego.GlobalControllerRouter["hiServer/controllers:UserController"],
		beego.ControllerComments{
			Method: "UnchTheClock",
			Router: `/user/person/punchTheClock`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiServer/controllers:UserController"] = append(beego.GlobalControllerRouter["hiServer/controllers:UserController"],
		beego.ControllerComments{
			Method: "Tag",
			Router: `/user/person/tag`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiServer/controllers:UserController"] = append(beego.GlobalControllerRouter["hiServer/controllers:UserController"],
		beego.ControllerComments{
			Method: "WinningLog",
			Router: `/user/person/winningLog`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiServer/controllers:UserController"] = append(beego.GlobalControllerRouter["hiServer/controllers:UserController"],
		beego.ControllerComments{
			Method: "AutoPush",
			Router: `/user/serve/autoPush`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiServer/controllers:UserController"] = append(beego.GlobalControllerRouter["hiServer/controllers:UserController"],
		beego.ControllerComments{
			Method: "CampusCard",
			Router: `/user/serve/campusCard`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiServer/controllers:UserController"] = append(beego.GlobalControllerRouter["hiServer/controllers:UserController"],
		beego.ControllerComments{
			Method: "Course",
			Router: `/user/serve/course`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiServer/controllers:UserController"] = append(beego.GlobalControllerRouter["hiServer/controllers:UserController"],
		beego.ControllerComments{
			Method: "Score",
			Router: `/user/serve/score`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiServer/controllers:UserController"] = append(beego.GlobalControllerRouter["hiServer/controllers:UserController"],
		beego.ControllerComments{
			Method: "SelfStudyRoom",
			Router: `/user/serve/selfStudyRoom`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiServer/controllers:UserController"] = append(beego.GlobalControllerRouter["hiServer/controllers:UserController"],
		beego.ControllerComments{
			Method: "TeacherMsg",
			Router: `/user/serve/teacherMsg`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
