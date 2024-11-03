package controllers

import (
	"enigma-new/models"
	"enigma-new/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type currentKey struct {
    CurrentKey models.Key `json:"current_key"`
}


// @Summary Obtém a chave atual
// @Description Retorna a chave atualmente ativa no sistema
// @Tags keys
// @Param Authorization header string true "Token JWT"
// @Produce json
// @Success 200 {object} currentKey
// @Failure 500 {object} utils.TypeErrorResponse "Erro interno"
// @Router /api/current-key [get]
func GetCurrentKey(c *gin.Context) {
    key, err := models.GetCurrentKey()
    if err != nil {
        c.JSON(http.StatusInternalServerError, utils.TypeErrorResponse{
            Error: err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, currentKey{
        CurrentKey: key,
    })
}

// @Summary Lista as chaves
// @Description Retorna uma lista de chaves armazenadas
// @Tags keys
// @Produce json
// @Param Authorization header string true "Token JWT"
// @Param page query int false "Número da página"
// @Success 200 {object} []models.Key
// @Failure 400 {object} utils.TypeErrorResponse
// @Router /api/keys [get]
func ListKeys(c *gin.Context) {
    pageStr := c.Query("page")

    if pageStr == "" {
        c.JSON(http.StatusBadRequest, utils.TypeErrorResponse{
            Error: "Parâmetros de paginação 'page' é obrigatório",
        })
        return
    }

    page, err := strconv.Atoi(pageStr)
    if err != nil || page < 1 {
        c.JSON(http.StatusBadRequest, utils.TypeErrorResponse{
            Error: "O parâmetro 'page' deve ser um número inteiro positivo",
        })
        return
    }
    result, keys := models.ListKeys(page)
    if  keys!= nil {
        c.JSON(http.StatusOK, gin.H{"keys":keys})
        return
    }

    c.JSON(http.StatusBadRequest, utils.TypeErrorResponse{
        Error: result.Error,
    })
}
