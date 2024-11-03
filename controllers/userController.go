package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "strconv"
    "enigma-new/models"
    "enigma-new/utils"
)

type Sucess struct {
    Message string `json:"message"`
}

type Create struct {
    Username string `json:"username"`
    Password string `json:"password"`
}
var user models.User

// @Summary Lista todos os usuários
// @Description Retorna uma lista de todos os usuários cadastrados
// @Tags users
// @Accept json
// @Param Authorization header string true "Token JWT"
// @Produce json
// @Success 200 {array} models.User
// @Router /api/users [get]
func ListUsers(c *gin.Context) {
    c.JSON(http.StatusOK, models.ListUsers())
}

// @Summary Cria um novo usuário
// @Description Cria um novo usuário com os dados fornecidos
// @Tags users
// @Accept json
// @Produce json
// @Param user body Create true "Dados do usuário"
// @Success 201 {object} models.User
// @Failure 400 {object} utils.TypeErrorResponse
// @Failure 500 {object} utils.TypeErrorResponse
// @Router /api/users [post]
func CreateUser(c *gin.Context) {
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, utils.TypeErrorResponse{
            Error: "Dados inválidos",
        })
        return
    }

    createdUser, err := models.CreateUser(user)
    if err != nil {
        c.JSON(http.StatusInternalServerError, utils.TypeErrorResponse{
            Error: "Erro ao criar o usuário",
        })
        return
    }

    c.JSON(http.StatusCreated, createdUser)
}

// @Summary Atualiza um usuário existente
// @Description Atualiza os dados de um usuário com base no ID fornecido
// @Tags users
// @Accept json
// @Produce json
// @Param Authorization header string true "Token JWT"
// @Param id path int true "ID do usuário"
// @Param user body Create true "Dados do usuário"
// @Success 200 {object} models.User
// @Failure 400 {object} utils.TypeErrorResponse
// @Failure 404 {object} utils.TypeErrorResponse
// @Failure 403 {object} utils.TypeErrorResponse
// @Router /api/users/{id} [put]
func UpdateUser(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)

    if err != nil {
        c.JSON(http.StatusBadRequest, utils.TypeErrorResponse{
            Error: "ID inválido",
        })
        return
    }

    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, utils.TypeErrorResponse{
            Error: "Dados inválidos",
        })
        return
    }

    user, err = models.UpdateUser(id, user)
    if err != nil {
        c.JSON(http.StatusForbidden, utils.TypeErrorResponse{
            Error: err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, user)
}

// @Summary Exclui um usuário
// @Description Exclui um usuário com base no ID fornecido
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "ID do usuário"
// @Param Authorization header string true "Token JWT"
// @Success 200 {object} Sucess
// @Failure 403 {object} utils.TypeErrorResponse
// @Router /api/users/{id} [delete]
func DeleteUser(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    err := models.DeleteUser(id)

    if err != nil {
        c.JSON(http.StatusForbidden, utils.TypeErrorResponse{
            Error: err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, Sucess{
        Message: "Usuário excluído com sucesso",
    })
}
