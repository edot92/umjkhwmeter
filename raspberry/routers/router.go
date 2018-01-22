// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/edot92/umjkhwmeter/raspberry/controllers"
	"github.com/edot92/umjkhwmeter/raspberry/helper"
)

var FilterRouter = func(ctx *context.Context) {
	if strings.HasPrefix(ctx.Input.URL(), "/api") {
		return
	}
	if !strings.Contains(ctx.Input.URL(), "#") {
		ctx.Redirect(302, "/#/"+ctx.Input.URL())
	}
	// if !ok {
	//     ctx.Redirect(302, "/login")
	// }
	// _, ok := ctx.Input.Session("uid").(int)
	// if !ok {
	//     ctx.Redirect(302, "/login")
	// }
}

func init() {
	helper.InitSocketio()
	beego.InsertFilter("/*", beego.BeforeRouter, FilterRouter)

	ns := beego.NewNamespace("/api",
		beego.NSNamespace("/alat",
			beego.NSInclude(
				&controllers.AlatController{},
			),
		),
	)
	beego.AddNamespace(ns)
	// enable cross origigin / request dari beda domain untuk http
	enableCors()
}
func enableCors() {
	//		// CORS for https://foo.* origins, allowing:
	//		// - PUT and PATCH methods
	//		// - Origin header
	//		// - Credentials share
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: []string{
			"http://localhost:*",
			"http://*:*",
			"http://*",
			"http://kumail8.pagekite.me/",
		},
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
}
