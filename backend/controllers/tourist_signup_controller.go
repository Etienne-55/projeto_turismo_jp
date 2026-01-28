package controllers

import (
	"net/http"
	"projeto_turismo_jp/models"
	"projeto_turismo_jp/repositories"

	"github.com/gin-gonic/gin"
)


func Signup(context *gin.Context) {
	var tourist models.Tourist
	err := context.ShouldBindJSON(&tourist)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message":"error"})
		return
	}

	err = repositories.Save(&tourist)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"error"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message":"user created successfully"})

}
