package controllers

import (
	"log"
	"net/http"
	"projeto_turismo_jp/models"

	"github.com/gin-gonic/gin"
)


func (tc *TripController) CreateTrip( context *gin.Context) {
	var trip models.Trip
	err := context.ShouldBindJSON(&trip)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "error"})
		log.Printf("error: %v", err)
		return
	}

	touristID := context.GetInt64("touristID")
	trip.TouristID = int(touristID)

	tc.repo.SaveTrip(&trip)

	context.JSON(http.StatusOK, gin.H{"message": "event saved"})
}

