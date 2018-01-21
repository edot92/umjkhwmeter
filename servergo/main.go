package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"time"

	"github.com/goburrow/modbus"
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

func main() {
	// Modbus RTU/ASCII
	handler := modbus.NewRTUClientHandler("/dev/ttyUSB0")
	handler.BaudRate = 9600
	handler.DataBits = 8
	handler.Parity = "N"
	handler.StopBits = 1
	handler.SlaveId = 1
	handler.Timeout = 2 * time.Second

	err := handler.Connect()
	if err != nil {
		log.Fatal(err)
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
				log.Fatal(err.Error())
				handler.Close()
				// goto scanSerialPort
			} else {
				x3 := []byte{resModbusByte[2], resModbusByte[3], resModbusByte[0], resModbusByte[1]}
				c3 := binary.BigEndian.Uint32(x3)
				res := math.Float32frombits(c3)
				if index == 0 {
					dataEnergyMeter.Frekuensi = fmt.Sprintf("%2f", res)
				} else if index == 1 {
					dataEnergyMeter.Cospi = fmt.Sprintf("%2f", res)
				} else if index == 2 {
					dataEnergyMeter.Tegangan = fmt.Sprintf("%2f", res)

				} else if index == 3 {
					dataEnergyMeter.Arus = fmt.Sprintf("%2f", res)
				}
			}
			fmt.Println(dataEnergyMeter.Arus)
		}
		time.Sleep(2 * time.Second)

		handler.Close()
	}

}
