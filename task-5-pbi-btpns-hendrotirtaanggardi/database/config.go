package database

import (
	"pbi-hendrotirta-btpns/mod/app"
	"pbi-hendrotirta-btpns/mod/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func InitDB() {
	// Koneksi ke database
	db, err := gorm.Open("mysql", "root:password@/pbi-hendrotirta-btpn?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Failed to connect to database")
	}
	DB = db

	// Auto Migrate tabel
	DB.AutoMigrate(&app.User{}, &models.Photo{})
}
