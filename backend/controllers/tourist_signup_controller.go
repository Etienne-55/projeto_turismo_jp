package controllers

import (
	"log"
	"net/http"
	"projeto_turismo_jp/models"

	"github.com/gin-gonic/gin"
)

type SignupRequest struct {
	Email    string `json:"email" binding:"required,email" example:"user@example.com"`
	Password string `json:"password" binding:"required,min=8" example:"password123"`
}

type SignupResponse struct {
	Message string `json:"message" example:"user created successfully"`
}

type ErrorResponse struct {
	Message string `json:"message" example:"error description"`
}

// Signup godoc
// @Summary      Create a new tourist account
// @Description  Register a new tourist user with email and password
// @Tags         Authentication
// @Accept       json
// @Produce      json
// @Param        request body models.Tourist true "Signup credentials"
// @Success      201 {object} map[string]string "User created successfully"
// @Failure      400 {object} map[string]string "Invalid request"
// @Failure      500 {object} map[string]string "Internal server error"
// @Router       /signup [post]
func (tc *TouristController) Signup(context *gin.Context) {
	var tourist models.Tourist

	err := context.ShouldBindJSON(&tourist)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message":"error"})
		log.Printf("error: %v", err)
		return
	}

	err = tc.repo.Save(&tourist)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"error"})
		log.Printf("error: %v", err)
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message":"user created successfully"})

}

