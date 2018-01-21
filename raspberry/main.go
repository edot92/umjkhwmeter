package main

import (
	"log"
	"time"

	"github.com/astaxie/beego"
	"github.com/edot92/umjkhwmeter/raspberry/helper"
	_ "github.com/edot92/umjkhwmeter/raspberry/routers"
	_ "github.com/fatih/color"
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

	// run arduino
	go func() {
		isArduinoEnable := beego.AppConfig.String("enableArduino")
		if isArduinoEnable == "true" {

		}
	}()
	// run modbus
	go func() {
		time.Sleep(1 * time.Second)
		for {
			isPMEnable := beego.AppConfig.String("enablePM")
			// color.Red(isPMEnable)
			if isPMEnable == "true" {
				err := helper.RunModbusDebug()
				if err != nil {
					log.Fatal(err)
				}
			}
		}

	}()
	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.SetStaticPath("/", "dist")
	beego.BConfig.CopyRequestBody = true
	beego.SetStaticPath("/static/js", "dist/static/js")
	beego.SetStaticPath("/static/css", "dist/static/css")
	beego.SetStaticPath("/static/fonts", "dist/static/fonts")

	beego.BConfig.Log.AccessLogs = false
	beego.Run()
}
