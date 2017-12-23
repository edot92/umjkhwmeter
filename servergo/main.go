package main

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/edot92/umjkhwmeter/servergo/iot"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-ini/ini"
	"github.com/goburrow/modbus"
	"github.com/metakeule/fmtdate"
	"github.com/parnurzeal/gorequest"
	ini "gopkg.in/ini.v1"
)

type structEnergyMeter struct {
	Tegangan  string
	Arus      string
	Cospi     string
	Frekuensi string
}

var register [15]uint16
var errModbus error
var dataEnergyMeter structEnergyMeter

func getEnergyMeter(c *gin.Context) {
	if errModbus != nil {
		c.JSON(200, gin.H{
			"error": true,
			// "userID":  claims["id"],
			"message": errModbus.Error(),
		})
		return
	}
	// claims := jwt.ExtractClaims(c)
	c.JSON(200, gin.H{
		"error":   false,
		"message": "data energy",
		// "userID":  claims["id"],
		"payload": dataEnergyMeter,
	})
}

func setPortDariWEB(c *gin.Context) {
	portnya := c.Query("port")
	cfg, err := ini.InsensitiveLoad("setting.ini")
	if err != nil {
		c.JSON(200, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}
	cfg.Section("").Key("PORT").SetValue(portnya)
	err = cfg.SaveTo("setting.ini")
	if err != nil {
		c.JSON(200, gin.H{
			"error":   true,
			"message": err.Error() + ", waktu:" + fmtdate.Format("hh-mm-ss DD-MM-YYYY", time.Now()),
		})
		return
	}
	c.JSON(200, gin.H{
		"error":   false,
		"message": "berhasil update",
	})
}

type paramBodySendServer struct {
	Mikro struct {
		Tegangan string `json:"tegangan"`
		Arus     string `json:"arus"`
	} `json:"mikro"`
	PowerMeter struct {
		Frekuensi string `json:"frekuensi"`
		Cospi     string `json:"cospi"`
		Tegangan  string `json:"tegangan"`
		Arus      string `json:"arus"`
	} `json:"power_meter"`
}

func runReadmikroAndDM6200() {
	go func() {
		for {

		BukaPort:
			// cfgG, _ := ini.InsensitiveLoad("setting.ini")
			// isPMEnable, _ := cfgG.Section("").Key("ENABLE_PM").Bool()
			// isMikroEnable, _ := cfgG.Section("").Key("ENABLE_MIKRO").Bool()
			// if (isMikroEnable==true){

			// }
			portnya, _, err := iot.BukaPort()
			if err != nil {
				log.Println(err)
				time.Sleep(3 * time.Second)
				goto BukaPort
			}
		scanSerialPort:
			_, data, err := iot.ReadAndParse(portnya)
			if err != nil {
				log.Println(err)
				time.Sleep(3 * time.Second)
				goto BukaPort
			}

			fmt.Println(data)

			//jalankan program dm6200
			cfg, err := iot.ReadKonfigurasi()
			if err != nil {
				time.Sleep(3 * time.Second)
				errModbus = err
				goto scanSerialPort
			}
			key, err := cfg.Section("").GetKey("PORTUSBPOWERMETER")
			if err != nil {
				time.Sleep(3 * time.Second)
				errModbus = err
				goto scanSerialPort
			}
			COMPORT := key.String()
			// Modbus RTU/ASCII
			handler := modbus.NewRTUClientHandler(COMPORT)
			handler.BaudRate = 9600
			handler.DataBits = 8
			handler.Parity = "N"
			handler.StopBits = 1
			handler.SlaveId = 1
			handler.Timeout = 5 * time.Second
			err = handler.Connect()
			if err != nil {
				if strings.Contains(err.Error(), "no such file or directory") {
					err = errors.New("KONEKSI PORT :" + COMPORT + " TIDAK SESUAI , silakan ubah konfigurasi COM PORT")
				} else if strings.Contains(err.Error(), "timeout") {
					handler.Close()
					err = errors.New("Tidak ada respon dari alat , silakan cek konfigurasi COM PORT , dan koneksi ke power meter")
				}
				errModbus = err
				fmt.Println(errModbus)
				time.Sleep(3 * time.Second)
				goto scanSerialPort
			}

			client := modbus.NewClient(handler)
			/*
			   A1 Current phase 1 					3929
			   V1 Voltage phase 1 to neutral 		3927
			   PF1 Power factor, phase 1			3923
			   F Frequency, Hz						3915
			*/
			indexRegister := []uint16{3915 + 1, 3923 + 1, 3927 + 1, 3929 + 1}
			for index := 0; index < 4; index++ {
				registerReq := indexRegister[index]
				log.Println(registerReq)
				resModbusByte, err := client.ReadHoldingRegisters(registerReq, 4)
				if err != nil {
					errModbus = err
					log.Println(err.Error())
					handler.Close()
					goto scanSerialPort
				}
				if index == 0 {
					frekuesni := float64(binary.LittleEndian.Uint16(resModbusByte[0:4]))
					dataEnergyMeter.Frekuensi = fmt.Sprintf("%2f", frekuesni)
				} else if index == 1 {
					cospi := float64(binary.LittleEndian.Uint16(resModbusByte[0:4]))
					dataEnergyMeter.Cospi = fmt.Sprintf("%2f", cospi)
				} else if index == 2 {
					tegangan := float64(binary.LittleEndian.Uint16(resModbusByte[0:4]))
					dataEnergyMeter.Tegangan = fmt.Sprintf("%2f", tegangan)
				} else if index == 3 {
					arus := float64(binary.BigEndian.Uint16(resModbusByte[0:4]))
					dataEnergyMeter.Arus = fmt.Sprintf("%2f", arus)
				}
			}
			handler.Close()
			// jalankan request ke server
			var paramJSONSend paramBodySendServer
			paramJSONSend.Mikro.Arus = data.Arus
			paramJSONSend.Mikro.Tegangan = data.Tegangan
			paramJSONSend.PowerMeter.Frekuensi = dataEnergyMeter.Frekuensi
			paramJSONSend.PowerMeter.Cospi = dataEnergyMeter.Cospi
			paramJSONSend.PowerMeter.Tegangan = dataEnergyMeter.Tegangan
			paramJSONSend.PowerMeter.Arus = dataEnergyMeter.Arus
			request := gorequest.New()

			r, _, errs := request.Post("http://localhost:8080/api/alat/updatenilaialat").
				Send(paramJSONSend).
				End()
			if len(errs) > 0 {
				fmt.Println(errs)
				goto scanSerialPort
			}
			b, err := ioutil.ReadAll(r.Body)
			defer r.Body.Close()
			if err != nil {
				log.Println(err)
				goto scanSerialPort

			}
			fmt.Println(string(b))
			// // Unmarshal
			// var msg Message
			// err = json.Unmarshal(b, &msg)
			// if err != nil {
			// 	goto scanSerialPort
			// 	return
			// }
			goto scanSerialPort
		}
	}()
}

func runDM6200() {

}
func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// http handler
	gin.SetMode(gin.ReleaseMode)

	portWeb := os.Getenv("PORT")
	r := gin.New()
	// r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.Default())
	if portWeb == "" {
		portWeb = "9999"
	}

	runReadmikroAndDM6200()
	runDM6200()

	//* endless loop */
	for {
		time.Sleep(100 * time.Millisecond)
	}
}
