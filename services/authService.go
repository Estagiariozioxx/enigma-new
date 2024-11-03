package services

import (
    "golang.org/x/crypto/bcrypt"
    "log"
)

func Authenticate(passwordInput, password string) bool {
    if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(passwordInput)); err != nil {
        log.Println("false")
        return false
    }
    return true
}

func HashPassword(password string) (passwordCrypt string,error error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "",err
    }
    Password := string(hashedPassword)
    return Password,nil
}

