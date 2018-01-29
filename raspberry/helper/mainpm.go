package helper

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/jimlawless/whereami"

	"github.com/astaxie/beego"

	"github.com/goburrow/modbus"
	"github.com/metakeule/fmtdate"
)

type StructEnergyMeter struct {
	Tegangan  string `json:"tegangan"`
	Arus      string `json:"arus"`
	Cospi     string `json:"cospi"`
	Frekuensi string `json:"frekuensi"`
}

var register [15]uint16
var errModbus error
var dataEnergyMeter StructEnergyMeter
var DataPMValue DataPM

func RunModbus() {
cobaKonek:
	PORT_PM := beego.AppConfig.String("PORT_PM")
	handler := modbus.NewRTUClientHandler(PORT_PM)
	handler.BaudRate = 9600
	handler.DataBits = 8
	handler.Parity = "N"
	handler.StopBits = 1
	handler.SlaveId = 1
	handler.Timeout = 2 * time.Second

	err := handler.Connect()
	if err != nil {
		fmt.Println(err)
		goto cobaKonek
	}
	defer handler.Close()

	client := modbus.NewClient(handler)
	for {
		/*
		   A1 Current phase 1 					3929
		   V1 Voltage phase 1 to neutral 		3927
		   PF1 Power factor, phase 1			3923
		   F Frequency, Hz						3915
		*/
		indexRegister := []uint16{3915 - 1, 3923 - 1, 3927 - 1, 3929 - 1}
		for index := 0; index < 4; index++ {
			registerReq := indexRegister[index]
			resModbusByte, err := client.ReadHoldingRegisters(registerReq, 4)
			if err != nil {
				errModbus = err
				handler.Close()
				goto cobaKonek

				// goto scanSerialPort
			} else {
				x3 := []byte{resModbusByte[2], resModbusByte[3], resModbusByte[0], resModbusByte[1]}
				c3 := binary.BigEndian.Uint32(x3)
				res := math.Float32frombits(c3)
				if index == 0 {
					dataEnergyMeter.Frekuensi = fmt.Sprintf("%4f", res)
				} else if index == 1 {
					if res == 0.00 {
						res = 1.0
					}
					res1 := math.Abs(float64(res))
					dataEnergyMeter.Cospi = fmt.Sprintf("%4f", res1)
				} else if index == 2 {
					dataEnergyMeter.Tegangan = fmt.Sprintf("%4f", res)

				} else if index == 3 {
					res=res/2
					dataEnergyMeter.Arus = fmt.Sprintf("%4f", res)
				}
			}
		}
		date := fmtdate.Format("YYYY-MM-DD hh:mm:ss", time.Now())
		date = date + ".000000000+07:00"
		DataPMValue.PMTegangan = dataEnergyMeter.Tegangan
		DataPMValue.PMArus = dataEnergyMeter.Arus
		DataPMValue.PMFrekuensi = dataEnergyMeter.Frekuensi
		DataPMValue.PMCospi = (dataEnergyMeter.Cospi)
		DataPMValue.Waktu = date
		fmt.Print("PM ARUs=")
		fmt.Println(dataEnergyMeter.Arus)
		time.Sleep(2 * time.Second)
	}
}
func RunModbusDebug() error {
	time.Sleep(5 * time.Second)
	var err error
	dataEnergyMeter.Tegangan = fmt.Sprintf("%4f", rand.Float64())
	dataEnergyMeter.Arus = fmt.Sprintf("%4f", rand.Float64())
	dataEnergyMeter.Cospi = fmt.Sprintf("%4f", rand.Float64())
	dataEnergyMeter.Frekuensi = fmt.Sprintf("%4f", rand.Float64())

	err = UpdatenilaialatPM(dataEnergyMeter)
	if err != nil {
		fmt.Println(err)
		return errors.New(err.Error() + whereami.WhereAmI())
	} else {
		// fmt.Println("sukses save")
	}
	return nil

}
func UpdatenilaialatPM(paramJSON StructEnergyMeter) error {
	/*
	   	{
	   	"mikro":{
	   		"teggangan":"0.1",
	   		"arus":"0.2"
	   	}
	   }
	*/
	var err error
	date := fmtdate.Format("YYYY-MM-DD hh:mm:ss", time.Now())
	date = date + ".000000000+07:00"
	err = DBKon.Create(
		&DataPM{
			PMTegangan:  paramJSON.Tegangan,
			PMArus:      paramJSON.Arus,
			PMFrekuensi: paramJSON.Cospi,
			PMCospi:     paramJSON.Frekuensi,
			Waktu:       date,
		},
	).Error

	DataPMValue.PMTegangan = paramJSON.Tegangan
	DataPMValue.PMArus = paramJSON.Arus
	DataPMValue.PMFrekuensi = paramJSON.Cospi
	DataPMValue.PMCospi = paramJSON.Frekuensi
	DataPMValue.Waktu = date
	return err

}
