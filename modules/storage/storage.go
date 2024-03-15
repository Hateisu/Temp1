package storage

import (
	"encoding/json"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type configuration struct {
	Is_inited bool
}

var DB *gorm.DB

func init() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB = db
	// Migrate the schema
	DB.AutoMigrate(&SavedConfig{})
	DB.AutoMigrate(&Role{})
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Country{})
	//DB.AutoMigrate(&Seller{})
	//	DB.AutoMigrate(&Product{})
	initial_create()
}

func initial_create() {
	var cnfs_temp []SavedConfig
	DB.Find(&cnfs_temp)
	if len(cnfs_temp) == 1 {
		cnf := configuration{}
		json.Unmarshal([]byte(cnfs_temp[0].Data), &cnf)
		log.Print(cnf.Is_inited)
	} else {
		DB.Create(&Role{RoleName: "ADMIN"})
		DB.Create(&Role{RoleName: "USER"})
		DB.Create(&Role{RoleName: "COMPANY"})
		data, err := json.Marshal(configuration{Is_inited: true})
		if err != nil {

		}
		DB.Create(&SavedConfig{Data: string(json.RawMessage(data))})

	}
}
