package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"time"

	"github.com/goburrow/modbus"
)

func DecodeHiLo(data []byte) (int16, error) {
	var i int16
	buf := bytes.NewReader(data)
	err := binary.Read(buf, binary.BigEndian, &i)

	return i, err
}
func main() {
	// log.SetFlags(log.LstdFlags | log.Lshortfile)

scanSerialPort:
	time.Sleep(2 * time.Second)
	COMPORT := "/dev/ttyUSB0"
	handler := modbus.NewRTUClientHandler(COMPORT)
	handler.BaudRate = 9600
	handler.DataBits = 8
	handler.Parity = "N"
	handler.StopBits = 1
	handler.SlaveId = 1
	handler.Timeout = 5 * time.Second
	err := handler.Connect()
	if err != nil {
		log.Println(err)
		handler.Close()
		goto scanSerialPort
	}

	client := modbus.NewClient(handler)
	resModbusByte, err := client.ReadHoldingRegisters(3910, 2)
	if err != nil {
		log.Println(err)
		handler.Close()
		goto scanSerialPort
	}

	fmt.Println(resModbusByte)
	handler.Close()
	// data1, _ := DecodeHiLo(resModbusByte[0:2])
	// data2, _ := DecodeHiLo(resModbusByte[2:4])
	// // a2 := math.Float32frombits(a)
	// fmt.Println(data1)
	// fmt.Println(data2)
	var i uint32
	buf := bytes.NewReader(resModbusByte)
	err = binary.Read(buf, binary.LittleEndian, &i)
	fmt.Println(math.Float32frombits(i) / 1000000)
	// fmt.Println()
	// for i := 0; i < 20*2; i++ {
	// 	datak1 := float32(binary.LittleEndian.Uint16(resModbusByte[i : i+2]))
	// 	fmt.Print("i:" + strconv.Itoa(i) + "=")
	// 	fmt.Println(datak1)
	// }
	// fmt.Println()

	// i := (3040 - 3030)
	// i := 22
	// datak1 := float32(binary.LittleEndian.Uint16(resModbusByte[i:i+4])) * 0.01
	// fmt.Print("i:" + strconv.Itoa(i) + "=")
	// fmt.Println(datak1)
	// a := fmt.Sprintf("%2f", datak)
	// fmt.Println(a)
	goto scanSerialPort

}
