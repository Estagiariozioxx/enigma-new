package controllers

import (
    "net/http"
    "enigma-new/models"
    "github.com/gin-gonic/gin"
    "enigma-new/utils"
)

var input models.InputDecrypt
// @Summary Descriptografa um documento
// @Description Recebe um documento e uma chave para descriptografá-lo
// @Tags decryption
// @Accept json
// @Produce json
// @Param Authorization header string true "Token JWT"
// @Param input body models.InputDecrypt true "Dados para decriptação"
// @Success 200 {object} models.DecryptResult
// @Failure 400 {object} utils.TypeErrorResponse "Dados inválidos"
// @Router /api/decrypt [post]
func DecryptDocument(c *gin.Context) {

    // Bind JSON input
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, utils.TypeErrorResponse{
            Error: "Dados inválidos",
        })
        return
    }
	result := models.DecryptDocument(input)

	c.JSON(http.StatusOK, result)


}

