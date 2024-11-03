package database

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "os"
    "log"
    "github.com/joho/godotenv"
)

var DB *gorm.DB


func DBOpen() error {
    if err := godotenv.Load(); err != nil {
        log.Println("Arquivo .env não encontrado, carregando variáveis de ambiente...")
    }

    var err error
    DB, err = gorm.Open("postgres", os.Getenv("DB_CONNECTION"))
    if err != nil {
        log.Println("Falha ao conectar ao banco de dados:", err)
        return err
    }

    if DB == nil {
        panic("Erro ao inicializar o banco de dados: ponteiro nulo")
    }

    return nil
}

func DBClose() {
    if DB != nil {
        DB.Close()
    }
}



