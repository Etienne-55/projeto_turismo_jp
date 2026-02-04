package controllers

import (
	"log"
	"net/http"
	"projeto_turismo_jp/models"
	"projeto_turismo_jp/utils"

	"github.com/gin-gonic/gin"
)


// Signup godoc
// @Summary      Login in a tourist account
// @Tags         Authentication
// @Accept       json
// @Param        request body models.Tourist true "Signup credentials"
// @Success      201 {object} map[string]string "User loged in"
// @Failure      400 {object} map[string]string "Invalid request"
// @Failure      500 {object} map[string]string "Internal server error"
// @Router       /login [post]
func (tc *TouristController) Login(context *gin.Context) {
	var tourist models.Tourist

	err := context.ShouldBindJSON(&tourist)
	if err != nil {
		log.Printf("error: %v", err)
		return	
	}

	err = tc.repo.ValidateCredentials(&tourist)
	if err != nil {
		log.Printf("error: %v", err)
		return
	}

	token, err := utils.GenerateToken(tourist.Email, tourist.ID, tourist.Role)
	if err != nil {
		log.Printf("token generation failed for user %d: %v", tourist.ID, err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not authenticate user", "error": err} )
		return 
	}

	context.JSON(http.StatusOK, gin.H{"message": "login successfull", "token": token })
}

