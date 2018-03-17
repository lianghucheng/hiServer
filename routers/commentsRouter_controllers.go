package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["hiChangda/serve/controllers:CustomerController"] = append(beego.GlobalControllerRouter["hiChangda/serve/controllers:CustomerController"],
		beego.ControllerComments{
			Method: "Achieve",
			Router: `/customer/community/achieve`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiChangda/serve/controllers:CustomerController"] = append(beego.GlobalControllerRouter["hiChangda/serve/controllers:CustomerController"],
		beego.ControllerComments{
			Method: "AutoPush",
			Router: `/customer/community/autoPush`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiChangda/serve/controllers:CustomerController"] = append(beego.GlobalControllerRouter["hiChangda/serve/controllers:CustomerController"],
		beego.ControllerComments{
			Method: "CampusCard",
			Router: `/customer/community/campusCard`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiChangda/serve/controllers:CustomerController"] = append(beego.GlobalControllerRouter["hiChangda/serve/controllers:CustomerController"],
		beego.ControllerComments{
			Method: "Course",
			Router: `/customer/community/course`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiChangda/serve/controllers:CustomerController"] = append(beego.GlobalControllerRouter["hiChangda/serve/controllers:CustomerController"],
		beego.ControllerComments{
			Method: "Feedback",
			Router: `/customer/community/feedback`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiChangda/serve/controllers:CustomerController"] = append(beego.GlobalControllerRouter["hiChangda/serve/controllers:CustomerController"],
		beego.ControllerComments{
			Method: "LostAndFound",
			Router: `/customer/community/lostAndFound`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiChangda/serve/controllers:CustomerController"] = append(beego.GlobalControllerRouter["hiChangda/serve/controllers:CustomerController"],
		beego.ControllerComments{
			Method: "Map",
			Router: `/customer/community/map`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiChangda/serve/controllers:CustomerController"] = append(beego.GlobalControllerRouter["hiChangda/serve/controllers:CustomerController"],
		beego.ControllerComments{
			Method: "MyPhoto",
			Router: `/customer/community/myPhoto`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiChangda/serve/controllers:CustomerController"] = append(beego.GlobalControllerRouter["hiChangda/serve/controllers:CustomerController"],
		beego.ControllerComments{
			Method: "Notice",
			Router: `/customer/community/notice`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiChangda/serve/controllers:CustomerController"] = append(beego.GlobalControllerRouter["hiChangda/serve/controllers:CustomerController"],
		beego.ControllerComments{
			Method: "PersonMsg",
			Router: `/customer/community/personMsg`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiChangda/serve/controllers:CustomerController"] = append(beego.GlobalControllerRouter["hiChangda/serve/controllers:CustomerController"],
		beego.ControllerComments{
			Method: "Photograph",
			Router: `/customer/community/photograph`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiChangda/serve/controllers:CustomerController"] = append(beego.GlobalControllerRouter["hiChangda/serve/controllers:CustomerController"],
		beego.ControllerComments{
			Method: "UnchTheClock",
			Router: `/customer/community/punchTheClock`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiChangda/serve/controllers:CustomerController"] = append(beego.GlobalControllerRouter["hiChangda/serve/controllers:CustomerController"],
		beego.ControllerComments{
			Method: "Score",
			Router: `/customer/community/score`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiChangda/serve/controllers:CustomerController"] = append(beego.GlobalControllerRouter["hiChangda/serve/controllers:CustomerController"],
		beego.ControllerComments{
			Method: "Secret",
			Router: `/customer/community/secret`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiChangda/serve/controllers:CustomerController"] = append(beego.GlobalControllerRouter["hiChangda/serve/controllers:CustomerController"],
		beego.ControllerComments{
			Method: "SelfStudyRoom",
			Router: `/customer/community/selfStudyRoom`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiChangda/serve/controllers:CustomerController"] = append(beego.GlobalControllerRouter["hiChangda/serve/controllers:CustomerController"],
		beego.ControllerComments{
			Method: "Tag",
			Router: `/customer/community/tag`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiChangda/serve/controllers:CustomerController"] = append(beego.GlobalControllerRouter["hiChangda/serve/controllers:CustomerController"],
		beego.ControllerComments{
			Method: "TeacherMsg",
			Router: `/customer/community/teacherMsg`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiChangda/serve/controllers:CustomerController"] = append(beego.GlobalControllerRouter["hiChangda/serve/controllers:CustomerController"],
		beego.ControllerComments{
			Method: "TechNews",
			Router: `/customer/community/techNews`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiChangda/serve/controllers:CustomerController"] = append(beego.GlobalControllerRouter["hiChangda/serve/controllers:CustomerController"],
		beego.ControllerComments{
			Method: "TechQuestion",
			Router: `/customer/community/techQuestion`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiChangda/serve/controllers:CustomerController"] = append(beego.GlobalControllerRouter["hiChangda/serve/controllers:CustomerController"],
		beego.ControllerComments{
			Method: "Topic",
			Router: `/customer/community/topic`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiChangda/serve/controllers:CustomerController"] = append(beego.GlobalControllerRouter["hiChangda/serve/controllers:CustomerController"],
		beego.ControllerComments{
			Method: "Vote",
			Router: `/customer/community/vote`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiChangda/serve/controllers:CustomerController"] = append(beego.GlobalControllerRouter["hiChangda/serve/controllers:CustomerController"],
		beego.ControllerComments{
			Method: "Weather",
			Router: `/customer/community/weather`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiChangda/serve/controllers:CustomerController"] = append(beego.GlobalControllerRouter["hiChangda/serve/controllers:CustomerController"],
		beego.ControllerComments{
			Method: "WinningLog",
			Router: `/customer/community/winningLog`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiChangda/serve/controllers:ObjectController"] = append(beego.GlobalControllerRouter["hiChangda/serve/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiChangda/serve/controllers:ObjectController"] = append(beego.GlobalControllerRouter["hiChangda/serve/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiChangda/serve/controllers:ObjectController"] = append(beego.GlobalControllerRouter["hiChangda/serve/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiChangda/serve/controllers:ObjectController"] = append(beego.GlobalControllerRouter["hiChangda/serve/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiChangda/serve/controllers:ObjectController"] = append(beego.GlobalControllerRouter["hiChangda/serve/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiChangda/serve/controllers:UserController"] = append(beego.GlobalControllerRouter["hiChangda/serve/controllers:UserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiChangda/serve/controllers:UserController"] = append(beego.GlobalControllerRouter["hiChangda/serve/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiChangda/serve/controllers:UserController"] = append(beego.GlobalControllerRouter["hiChangda/serve/controllers:UserController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiChangda/serve/controllers:UserController"] = append(beego.GlobalControllerRouter["hiChangda/serve/controllers:UserController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiChangda/serve/controllers:UserController"] = append(beego.GlobalControllerRouter["hiChangda/serve/controllers:UserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiChangda/serve/controllers:UserController"] = append(beego.GlobalControllerRouter["hiChangda/serve/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hiChangda/serve/controllers:UserController"] = append(beego.GlobalControllerRouter["hiChangda/serve/controllers:UserController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
