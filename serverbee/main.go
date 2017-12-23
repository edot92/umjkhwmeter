package main

import (
	"log"

	"github.com/astaxie/beego"
	"github.com/edot92/umjkhwmeter/serverbee/helper"
	_ "github.com/edot92/umjkhwmeter/serverbee/routers"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	db, err := helper.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	beego.Run()
}
