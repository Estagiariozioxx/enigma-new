package controllers

import (
    "net/http"
    "enigma-new/models"
    "enigma-new/utils"
    "github.com/gin-gonic/gin"
    "golang.org/x/time/rate"
    "sync"
    "time"
)

type Login struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

type Loginsucess struct {
    Token string `json:"token"`
}


var limiterStore = struct {
    sync.RWMutex
    m map[string]*rate.Limiter
}{m: make(map[string]*rate.Limiter)}



func getLimiter(ip string) *rate.Limiter {
    limiterStore.Lock()
    defer limiterStore.Unlock()

    if limiter, exists := limiterStore.m[ip]; exists {
        return limiter
    }

    limiter := rate.NewLimiter(1, 5)
    limiterStore.m[ip] = limiter

    go func() {
        time.Sleep(10 * time.Minute)
        limiterStore.Lock()
        delete(limiterStore.m, ip)
        limiterStore.Unlock()
    }()

    return limiter
}



// @Summary Login do usuário
// @Description Realiza o login do usuário e retorna um token JWT
// @Tags auth
// @Accept json
// @Produce json
// @Param login body Login true "Dados de login"
// @Success 200 {object} Loginsucess
// @Failure 400 {object} utils.TypeErrorResponse
// @Failure 401 {object} utils.TypeErrorResponse
// @Failure 429 {object} utils.TypeErrorResponse "Muitas tentativas de login"
// @Router /api/login [post]
func Loginn(c *gin.Context) {

    clientIP := c.ClientIP()
    limiter := getLimiter(clientIP)

    if !limiter.Allow() {
        c.JSON(http.StatusTooManyRequests, utils.TypeErrorResponse{
            Error: "Muitas tentativas de login. Tente novamente mais tarde.",
        })
        return
    }

    var loginData Login

    if err := c.ShouldBindJSON(&loginData); err != nil {
        c.JSON(http.StatusBadRequest, utils.TypeErrorResponse{
            Error: "Dados inválidos",
        })
        return
    }

    user, err := models.VerifyUser(loginData.Username, loginData.Password)
    if err != nil {
        c.JSON(http.StatusUnauthorized, utils.TypeErrorResponse{
            Error: "Credenciais inválidas",
        })
        return
    }
    if user == nil {
        c.JSON(http.StatusInternalServerError, utils.TypeErrorResponse{
            Error: "Usuário não encontrado",
        })
        return
    }

    token, err := utils.GenerateJWT(user.ID)
    if err != nil {

        c.JSON(http.StatusInternalServerError, utils.TypeErrorResponse{
            Error: "Falha ao gerar o token",
        })
        return
    }

    c.JSON(http.StatusOK, Loginsucess{
        Token: token,
    })

}
