package helper

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"

	serial "go.bug.st/serial.v1"

	"github.com/astaxie/beego"

	"github.com/metakeule/fmtdate"
)

type StructEnergyArduino struct {
	Tegangan string `json:"tegangan"`
	Arus1    string `json:"arus1"`
	Arus2    string `json:"arus2"`
	Arus3    string `json:"arus3"`
}

func RunArduino() {
cobaKonek:
	PORT_ARDUINO := beego.AppConfig.String("PORT_ARDUINO")
	baudrate := 9600
	config := &serial.Mode{
		BaudRate: baudrate,
		// DataBits: databits,
		Parity:   serial.NoParity,
		StopBits: serial.OneStopBit,
	}
	PortOpen, err := serial.Open(PORT_ARDUINO, config)
	if err != nil {
		goto cobaKonek
	}
	defer PortOpen.Close()
	for {
		var buf [100]byte
		var pesan string
	bacalagi:
		n, err := PortOpen.Read(buf[:])
		if err != nil {
			break
		}
		pesan = pesan + string(buf[:n])
		if strings.Contains(pesan, "#") == false {
			goto bacalagi
		}
		pesan = strings.Replace(pesan, "#", "", -1)
		pesan = strings.Replace(pesan, "\r", "", -1)
		pesan = strings.Replace(pesan, "\n", "", -1)
		var datak StructEnergyArduino
		err = json.Unmarshal([]byte(pesan), &datak)
		if err != nil {
			break
		}
		DataPMValue.MikroTegangan = datak.Tegangan
		DataPMValue.MikroArus1 = datak.Arus1
		DataPMValue.MikroArus2 = datak.Arus2
		DataPMValue.MikroArus3 = datak.Arus3
		fmt.Print("PM=")
		fmt.Println(DataPMValue)
		time.Sleep(1500 * time.Millisecond)
	}
	defer PortOpen.Close()

}
func RunArduinoDebug() error {
	time.Sleep(5 * time.Second)
	var err error

	DataPMValue.MikroTegangan = fmt.Sprintf("%2f", rand.Float64())
	DataPMValue.MikroArus1 = fmt.Sprintf("%2f", rand.Float64())
	DataPMValue.MikroArus2 = fmt.Sprintf("%2f", rand.Float64())
	DataPMValue.MikroArus3 = fmt.Sprintf("%2f", rand.Float64())
	date := fmtdate.Format("YYYY-MM-DD hh:mm:ss", time.Now())
	date = date + ".000000000+07:00"
	err = DBKon.Create(
		&DataPM{
			PMTegangan:  dataEnergyMeter.Tegangan,
			PMArus:      dataEnergyMeter.Arus,
			PMCospi:     dataEnergyMeter.Cospi,
			PMFrekuensi: dataEnergyMeter.Frekuensi,
			Waktu:       date,
		},
	).Error
	return err

}
func UpdatenilaialatArduino(paramJSON StructEnergyMeter) error {
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
