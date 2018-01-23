package main

import (
	"fmt"
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
	} else {
		// time.Sleep()
	}
	db, err := helper.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// run arduino
	go func() {
		time.Sleep(1 * time.Second)
		isArduinoEnable := beego.AppConfig.String("enableArduino")
		isSImulasi, _ := beego.AppConfig.Bool("SIMULASI_ARDUINO")
		if isArduinoEnable == "true" {
			if isSImulasi {
				fmt.Println("SIMULASI_ARDUINO")

				err := helper.RunArduinoDebug()
				if err != nil {
					log.Fatal(err)
				}
			} else {
				helper.RunArduino()
			}
		} else {
			fmt.Println("ARDUINO no enable")

		}
	}()
	// run modbus
	go func() {
		time.Sleep(1 * time.Second)
		for {
			time.Sleep(1 * time.Second)
			isPMEnable := beego.AppConfig.String("enablePM")
			if isPMEnable == "true" {
				isSImulasi, _ := beego.AppConfig.Bool("SIMULASI_PM")
				if isSImulasi {
					fmt.Println("SIMULASI_PM")
					err := helper.RunModbusDebug()
					if err != nil {
						log.Fatal(err)
					}
				} else {
					helper.RunModbus()
				}
			} else {
				fmt.Println("PM no enable")
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
