package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/edot92/umjkhwmeter/raspberry/controllers:AlatController"] = append(beego.GlobalControllerRouter["github.com/edot92/umjkhwmeter/raspberry/controllers:AlatController"],
		beego.ControllerComments{
			Method: "Getnilaimikro",
			Router: `/getnilaimikro`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/edot92/umjkhwmeter/raspberry/controllers:AlatController"] = append(beego.GlobalControllerRouter["github.com/edot92/umjkhwmeter/raspberry/controllers:AlatController"],
		beego.ControllerComments{
			Method: "Getnilaipm",
			Router: `/getnilaipm`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/edot92/umjkhwmeter/raspberry/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/edot92/umjkhwmeter/raspberry/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/edot92/umjkhwmeter/raspberry/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/edot92/umjkhwmeter/raspberry/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/edot92/umjkhwmeter/raspberry/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/edot92/umjkhwmeter/raspberry/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/edot92/umjkhwmeter/raspberry/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/edot92/umjkhwmeter/raspberry/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/edot92/umjkhwmeter/raspberry/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/edot92/umjkhwmeter/raspberry/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/edot92/umjkhwmeter/raspberry/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/edot92/umjkhwmeter/raspberry/controllers:UserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/edot92/umjkhwmeter/raspberry/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/edot92/umjkhwmeter/raspberry/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/edot92/umjkhwmeter/raspberry/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/edot92/umjkhwmeter/raspberry/controllers:UserController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/edot92/umjkhwmeter/raspberry/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/edot92/umjkhwmeter/raspberry/controllers:UserController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/edot92/umjkhwmeter/raspberry/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/edot92/umjkhwmeter/raspberry/controllers:UserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/edot92/umjkhwmeter/raspberry/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/edot92/umjkhwmeter/raspberry/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/edot92/umjkhwmeter/raspberry/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/edot92/umjkhwmeter/raspberry/controllers:UserController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
