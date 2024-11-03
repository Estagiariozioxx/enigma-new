package config

import (
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "os"
	"enigma-new/models"
	"enigma-new/database"

)


func InitDB() {
	var err error

	database.DBOpen()
	
    if err = database.DB.AutoMigrate(&models.User{}, &models.Key{}).Error; err != nil {
        panic("Erro ao criar as tabelas do banco de dados: " + err.Error())
    }

	models.CreateMasterUser(database.DB)

	database.DB.Close()

}

var JWTSecret = os.Getenv("JWT_SECRET")
