package controllers

import (
	"github.com/astaxie/beego"
	"github.com/edot92/umjkhwmeter/raspberry/helper"
)

// Operations about object
type AlatController struct {
	beego.Controller
}

// /
// func (c *AlatController) Updatenilaialat() {
// 	/*
// 	   	{
// 	   	"mikro":{
// 	   		"teggangan":"0.1",
// 	   		"arus":"0.2"
// 	   	}
// 	   }
// 	*/
// 	type paramBody struct {
// 		Mikro struct {
// 			Tegangan string `json:"tegangan"`
// 			Arus     string `json:"arus"`
// 		} `json:"mikro"`
// 		PowerMeter struct {
// 			Frekuensi string `json:"frekuensi"`
// 			Cospi     string `json:"cospi"`
// 			Tegangan  string `json:"tegangan"`
// 			Arus      string `json:"arus"`
// 		} `json:"power_meter"`
// 	}
// 	var paramJSON paramBody
// 	err := json.Unmarshal(c.Ctx.Input.RequestBody, &paramJSON)
// 	if err != nil {
// 		c.Data["json"] = map[string]interface{}{
// 			"error":   true,
// 			"message": err.Error(),
// 		}
// 		c.ServeJSON()
// 		return
// 	}
// 	date := fmtdate.Format("YYYY-MM-DD hh:mm:ss", time.Now())
// 	date = date + ".000000000+07:00"
// 	err = helper.DBKon.Create(
// 		&helper.DataAlat{
// 			TeganganMikro: paramJSON.Mikro.Tegangan,
// 			ArusMikro:     paramJSON.Mikro.Arus,
// 			PMFrekuensi:   paramJSON.PowerMeter.Frekuensi,
// 			PMCospi:       paramJSON.PowerMeter.Cospi,
// 			PMTegangan:    paramJSON.PowerMeter.Tegangan,
// 			PMArus:        paramJSON.PowerMeter.Arus,
// 			Waktu:         date,
// 		},
// 	).Error
// 	if err != nil {
// 		c.Data["json"] = map[string]interface{}{
// 			"error":   true,
// 			"message": err.Error(),
// 		}
// 		c.ServeJSON()
// 		return
// 	}
// 	c.Data["json"] = map[string]interface{}{
// 		"error":   false,
// 		"message": "berhasil menyimpan",
// 	}
// 	c.ServeJSON()
// 	return

// }

// @router /getnilaipm [get]
func (c *AlatController) Getnilaipm() {

	c.Data["json"] = map[string]interface{}{
		"error": false,
		"payload": map[string]interface{}{
			"data": helper.DataPMValue,
		},
		"message": "berhasil ambil data pm",
	}
	c.ServeJSON()
	return

}
