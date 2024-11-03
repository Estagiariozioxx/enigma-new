package utils

import (
    "errors"
    "github.com/golang-jwt/jwt/v4"
    "time"
)

var jwtKey = []byte("minha_chave_secreta")

type Claims struct {
    UserID uint `json:"user_id"`
    jwt.StandardClaims
}

func GenerateJWT(userID uint) (string, error) {
    expirationTime := time.Now().Add(24 * time.Hour)
    claims := &Claims{
        UserID: userID,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtKey)
}

func ValidateJWT(tokenString string) (*Claims, error) {
    claims := &Claims{}
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })
    if err != nil || !token.Valid {
        return nil, errors.New("token inválido")
    }
    return claims, nil
}
