package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["hiServer/controllers:CustomerController"] = append(beego.GlobalControllerRouter["hiServer/controllers:CustomerController"],
		beego.ControllerComments{
			Method: "LostAndFound",
			Router: `/customer/community/lostAndFound`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiServer/controllers:CustomerController"] = append(beego.GlobalControllerRouter["hiServer/controllers:CustomerController"],
		beego.ControllerComments{
			Method: "Map",
			Router: `/customer/community/map`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiServer/controllers:CustomerController"] = append(beego.GlobalControllerRouter["hiServer/controllers:CustomerController"],
		beego.ControllerComments{
			Method: "Notice",
			Router: `/customer/community/notice`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiServer/controllers:CustomerController"] = append(beego.GlobalControllerRouter["hiServer/controllers:CustomerController"],
		beego.ControllerComments{
			Method: "Photograph",
			Router: `/customer/community/photograph`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiServer/controllers:CustomerController"] = append(beego.GlobalControllerRouter["hiServer/controllers:CustomerController"],
		beego.ControllerComments{
			Method: "Secret",
			Router: `/customer/community/secret`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiServer/controllers:CustomerController"] = append(beego.GlobalControllerRouter["hiServer/controllers:CustomerController"],
		beego.ControllerComments{
			Method: "TechNews",
			Router: `/customer/community/techNews`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiServer/controllers:CustomerController"] = append(beego.GlobalControllerRouter["hiServer/controllers:CustomerController"],
		beego.ControllerComments{
			Method: "TechQuestion",
			Router: `/customer/community/techQuestion`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiServer/controllers:CustomerController"] = append(beego.GlobalControllerRouter["hiServer/controllers:CustomerController"],
		beego.ControllerComments{
			Method: "Topic",
			Router: `/customer/community/topic`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiServer/controllers:CustomerController"] = append(beego.GlobalControllerRouter["hiServer/controllers:CustomerController"],
		beego.ControllerComments{
			Method: "Vote",
			Router: `/customer/community/vote`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiServer/controllers:CustomerController"] = append(beego.GlobalControllerRouter["hiServer/controllers:CustomerController"],
		beego.ControllerComments{
			Method: "Weather",
			Router: `/customer/community/weather`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiServer/controllers:CustomerController"] = append(beego.GlobalControllerRouter["hiServer/controllers:CustomerController"],
		beego.ControllerComments{
			Method: "Achieve",
			Router: `/customer/person/achieve`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiServer/controllers:CustomerController"] = append(beego.GlobalControllerRouter["hiServer/controllers:CustomerController"],
		beego.ControllerComments{
			Method: "Feedback",
			Router: `/customer/person/feedback`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiServer/controllers:CustomerController"] = append(beego.GlobalControllerRouter["hiServer/controllers:CustomerController"],
		beego.ControllerComments{
			Method: "MyPhoto",
			Router: `/customer/person/myPhoto`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiServer/controllers:CustomerController"] = append(beego.GlobalControllerRouter["hiServer/controllers:CustomerController"],
		beego.ControllerComments{
			Method: "PersonMsg",
			Router: `/customer/person/personMsg`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiServer/controllers:CustomerController"] = append(beego.GlobalControllerRouter["hiServer/controllers:CustomerController"],
		beego.ControllerComments{
			Method: "UnchTheClock",
			Router: `/customer/person/punchTheClock`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiServer/controllers:CustomerController"] = append(beego.GlobalControllerRouter["hiServer/controllers:CustomerController"],
		beego.ControllerComments{
			Method: "Tag",
			Router: `/customer/person/tag`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiServer/controllers:CustomerController"] = append(beego.GlobalControllerRouter["hiServer/controllers:CustomerController"],
		beego.ControllerComments{
			Method: "WinningLog",
			Router: `/customer/person/winningLog`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiServer/controllers:CustomerController"] = append(beego.GlobalControllerRouter["hiServer/controllers:CustomerController"],
		beego.ControllerComments{
			Method: "AutoPush",
			Router: `/customer/serve/autoPush`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiServer/controllers:CustomerController"] = append(beego.GlobalControllerRouter["hiServer/controllers:CustomerController"],
		beego.ControllerComments{
			Method: "CampusCard",
			Router: `/customer/serve/campusCard`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiServer/controllers:CustomerController"] = append(beego.GlobalControllerRouter["hiServer/controllers:CustomerController"],
		beego.ControllerComments{
			Method: "Course",
			Router: `/customer/serve/course`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiServer/controllers:CustomerController"] = append(beego.GlobalControllerRouter["hiServer/controllers:CustomerController"],
		beego.ControllerComments{
			Method: "Score",
			Router: `/customer/serve/score`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiServer/controllers:CustomerController"] = append(beego.GlobalControllerRouter["hiServer/controllers:CustomerController"],
		beego.ControllerComments{
			Method: "SelfStudyRoom",
			Router: `/customer/serve/selfStudyRoom`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiServer/controllers:CustomerController"] = append(beego.GlobalControllerRouter["hiServer/controllers:CustomerController"],
		beego.ControllerComments{
			Method: "TeacherMsg",
			Router: `/customer/serve/teacherMsg`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
