package middlewares

import (
    "net/http"
    "enigma-new/utils"
    "github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Captura o token do cabeçalho
        tokenString := c.GetHeader("Authorization")

        if tokenString == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Token não fornecido"})
            c.Abort()
            return
        }

        // Valida o token JWT
        claims, err := utils.ValidateJWT(tokenString)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
            c.Abort()
            return
        }

        // Passa o contexto do usuário autenticado para as próximas requisições
        c.Set("user_id", claims.UserID)
        c.Next()
    }
}
