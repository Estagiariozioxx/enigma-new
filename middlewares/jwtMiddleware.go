package middlewares

import (
    "net/http"
    "enigma-new/utils"
    "github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")
        if tokenString == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Token não fornecido"})
            c.Abort()
            return
        }

        claims, err := utils.ValidateJWT(tokenString)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
            c.Abort()
            return
        }

        if claims == nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Token inválido"})
            c.Abort()
            return
        }

        c.Next()
    }
}
