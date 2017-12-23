package helper

import (
	"github.com/jinzhu/gorm"
)

func init() {

}

var DBKon *gorm.DB

type DataAlat struct {
	gorm.Model
	TeganganMikro string `json:"teganganMikro" gorm:"column:teganganMikro"`
	ArusMikro     string `json:"arusMikro" gorm:"column:arusMikro"`
	Waktu         string `json:"waktu" gorm:"column:waktu"`
	PMFrekuensi   string `json:"pmFrekuensi" gorm:"pmFrekuensi"`
	PMCospi       string `json:"pmCospi" gorm:"pmCospi"`
	PMTegangan    string `json:"pmTegangan" gorm:"pmTegangan"`
	PMArus        string `json:"pmArus" gorm:"pmArus"`
}

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "tes.db")
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&DataAlat{})
	DBKon = db
	return db, nil
}
