package iot

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/goinggo/tracelog"
	ini "gopkg.in/ini.v1"

	serial "go.bug.st/serial.v1"
)

type DataSerialMikro struct {
	Tegangan string `json:"tegangan"`
	Arus     string `json:"arus"`
}

// {"Temperature":"0",
// "Humidity":0,
// "Busvoltage":"12.35",
// "Shuntvoltage":"0.81",
// "Loadvoltage":"12.35",
// "Current":"7.90"}

// ser serial

// var ComPilihan = "COM4"
var ComPilihan = "/dev/ttyACM0"

func init() {
	tracelog.Start(tracelog.LevelTrace)

}

// ReadKonfigurasi baca konfigurasi file setting.ini
func ReadKonfigurasi() (*ini.File, error) {
	cfg, err := ini.InsensitiveLoad("setting.ini")
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
func readKonfigurasiFile() (string, error) {
	cfg, err := ini.InsensitiveLoad("setting.ini")
	if err != nil {
		return "", err
	}
	key, err := cfg.Section("").GetKey("PORTUSB")
	if err != nil {
		return "", err
	}
	return key.String(), nil
}
func BukaPort() (serial.Port, string, error) {

	var (
		// address  string
		baudrate int
		// databits int
	)
	ComPilihan, err := readKonfigurasiFile()
	if err != nil {
		// fmt.Println(err)
		return nil, ComPilihan, err
	}
	ports, err := serial.GetPortsList()
	if err != nil {
		// fmt.Println(err)
		// return false, ComPilihan, err
	}
	if len(ports) == 0 {
		tracelog.Warning("serial_engine", "BukaPort", "No serial ports found!")
		// tracelog.Stop()
		// fmt.Println("No serial ports found!")
		// return false
		return nil, ComPilihan, nil
	}

	baudrate = 9600
	// databits = 8
	// _ = databits
	config := &serial.Mode{
		BaudRate: baudrate,
		// DataBits: databits,
		Parity:   serial.NoParity,
		StopBits: serial.OneStopBit,
	}

	// var serMode serial.Mode
	// serMode.
	for _, port := range ports {
		fmt.Printf("Found port on: %v\n", port)
		// ComPilihan = port
	}

	portOpen, err := serial.Open(ComPilihan, config)
	if err != nil {
		fmt.Println("open error")
		fmt.Println(err)
		return nil, ComPilihan, err
	}
	return portOpen, ComPilihan, nil
}

// ReadAndParse ...
func ReadAndParse(portnya serial.Port) (bool, DataSerialMikro, error) {
	var dataSerial DataSerialMikro

	if portnya == nil {
		return false, dataSerial, errors.New("port is nil")
	}

	// defer portOpen.Close()
	var buf [100]byte
	var pesan string
bacalagi:
	n, err := portnya.Read(buf[:])
	if err != nil {
		return false, dataSerial, err
	}
	pesan = pesan + string(buf[:n])
	if strings.Contains(pesan, "#") == false {
		goto bacalagi
	}
	pesan = strings.Replace(pesan, "#", "", -1)
	pesan = strings.Replace(pesan, "\r", "", -1)
	pesan = strings.Replace(pesan, "\n", "", -1)
	err = json.Unmarshal([]byte(pesan), &dataSerial)
	if err != nil {
		fmt.Println(err)
		fmt.Println(pesan)

		return false, dataSerial, errors.New("failed format type ")
	}
	fmt.Println("pesan:", pesan)
	// fmt.Println(dataSerial)
	// err = SimpanData(dataSerial)
	// fmt.Println(err)
	return true, dataSerial, nil

}
