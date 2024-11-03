package models

import (
	"enigma-new/database"
	"enigma-new/services"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
    ID        uint   `gorm:"primary_key" json:"id"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
    Username  string `gorm:"unique;not null"`
    Password  string `gorm:"not null"`
    IsMaster  bool   `gorm:"default:false"`
}




func CreateMasterUser(DB *gorm.DB) {
	var user User

    
    if DB == nil {
        panic("Falha na conexao")
    }
	if err := database.DB.Where("is_master = ?", true).First(&user).Error; err == nil {
		return 
	}

    Password:= "senhamestre" 
    passwordCrypt, err := services.HashPassword(Password)
    if err != nil {
        panic("Falha ao gerar senha " + err.Error())
    }
    user.Password = passwordCrypt

    if err := DB.Where(User{IsMaster: true}).FirstOrCreate(&user, User{
        Username: "MESTRE",
        Password: user.Password,
        IsMaster: true,
    }).Error; err != nil {
        panic("Erro ao criar usuário mestre " + err.Error())
    }
}


func ListUsers() []User {
	database.DBOpen()

    var users []User

    database.DB.Find(&users)

	database.DBClose()

	return users
}


func CreateUser(user User) (User, error) {
    passwordCrypt, err := services.HashPassword(user.Password)
    if err != nil {
        return User{}, fmt.Errorf("erro ao hashear a senha: %w", err)
    }

    if err := database.DBOpen(); err != nil {
        return User{}, fmt.Errorf("falha ao abrir o banco de dados: %w", err)
    }
    defer database.DBClose() 

    user.Password = passwordCrypt

    if err := database.DB.Create(&user).Error; err != nil {
        return User{}, fmt.Errorf("erro ao criar o usuário: %w", err)
    }

    return user, nil
}


func UpdateUser(id int, updatedData User) (User, error) {
    var user User

    if err := database.DBOpen(); err != nil {
        return User{}, fmt.Errorf("falha ao abrir o banco de dados: %w", err)
    }
    defer database.DBClose()

    if database.DB.First(&user, id).RecordNotFound() {
        return User{}, fmt.Errorf("usuário não encontrado")
    }

    if user.IsMaster {
        return User{}, fmt.Errorf("o usuário MESTRE não pode ser modificado")
    }

    passwordCrypt, err := services.HashPassword(updatedData.Password)

    if err != nil {
        return User{}, fmt.Errorf("erro ao hashear a senha: %w", err)
    }

    user.Username = updatedData.Username
    user.Password = passwordCrypt

    if err := database.DB.Save(&user).Error; err != nil {
        return User{}, fmt.Errorf("erro ao salvar as alterações: %w", err)
    }

    return user, nil
}

func DeleteUser(id int) error {
	var user User

	if err := database.DBOpen(); err != nil {
        return fmt.Errorf("falha ao abrir o banco de dados: %w", err)
    }
    defer database.DBClose()

    if database.DB.First(&user, id).RecordNotFound() {
        return fmt.Errorf("Usuário não encontrado")
    }

    if user.IsMaster {
        return fmt.Errorf("Usuário MESTRE não pode ser excluído")
    }

    database.DB.Delete(&user)
    return nil
}


func VerifyUser(username, password string) (*User, error){

    var user User

	if err := database.DBOpen(); err != nil {
        log.Println("Erro ao conectar ao banco de dados:", err)
        return nil, errors.New("falha ao abrir banco de dados")
        
    }
	defer database.DBClose()
    database.DB.Where("username = ?", username).First(&user)
    if user.ID == 0 {
        log.Println("uário não encontrado:")
        return nil, errors.New("usuário não encontrado")
        
    }

    if(services.Authenticate(password,user.Password)){
        return &user, nil
    }
    log.Println("senha inválida")
    return nil, errors.New("senha inválida")

}



