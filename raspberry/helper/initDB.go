package helper

import (
	"github.com/jinzhu/gorm"
)

func init() {

}

var DBKon *gorm.DB

// type DataAlat struct {
// 	gorm.Model
// 	TeganganMikro string `json:"teganganMikro" gorm:"column:teganganMikro"`
// 	ArusMikro     string `json:"arusMikro" gorm:"column:arusMikro"`
// 	Waktu         string `json:"waktu" gorm:"column:waktu"`
// 	PMFrekuensi   string `json:"pmFrekuensi" gorm:"pmFrekuensi"`
// 	PMCospi       string `json:"pmCospi" gorm:"pmCospi"`
// 	PMTegangan    string `json:"pmTegangan" gorm:"pmTegangan"`
// 	PMArus        string `json:"pmArus" gorm:"pmArus"`
// }
type DataPM struct {
	gorm.Model
	PMFrekuensi string `json:"pmFrekuensi" gorm:"pmFrekuensi"`
	PMCospi     string `json:"pmCospi" gorm:"pmCospi"`
	PMTegangan  string `json:"pmTegangan" gorm:"pmTegangan"`
	PMArus      string `json:"pmArus" gorm:"pmArus"`

	MikroTegangan string `json:"mikroTegangan" gorm:"mikroTegangan"`
	MikroArus1    string `json:"mikroArus1" gorm:"mikroArus1"`
	MikroArus2    string `json:"mikroArus2" gorm:"mikroArus2"`
	MikroArus3    string `json:"mikroArus3" gorm:"mikroArus3"`
	Waktu         string `json:"waktu" gorm:"column:waktu"`
}

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "tes1.db")
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&DataPM{})
	DBKon = db
	return db, nil
}
