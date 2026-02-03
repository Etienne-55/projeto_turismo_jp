package controllers

import (
	"log"
	"net/http"
	"projeto_turismo_jp/models"

	"github.com/gin-gonic/gin"
)

type CreateTripRequest struct {
	LodgingLocation 	string 		`json:"lodging_location" binding:"required"`
	TripDescription 	string 		`json:"trip_description" binding:required"`
	ArrivalDate 			string		`json:"arrival_date" binding:"required"`
	DepartureDate 		string 		`json:"departure_date" binding:"required"`
}

type CreateTripResponse struct {
	Message string `json:"message" example:"user created successfully"`
}

// CreateTrip godoc
// @Summary      Create a new trip from logged user
// @Description  Registered tourist creates their upcoming trip
// @Tags         Trip
// @Accept       json
// @Produce      json
// @Param        request body controllers.CreateTripRequest true "Trip details"
// @Success      201 {object} controllers.CreateTripResponse "Trip created successfully"
// @Failure      400 {object} map[string]string "Invalid request"
// @Failure      500 {object} map[string]string "Internal server error"
// @Security     Bearer
// @Router       /trip [post]
func (tc *TripController) CreateTrip(context *gin.Context) {
	var trip models.Trip
	err := context.ShouldBindJSON(&trip)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "error"})
		log.Printf("error: %v", err)
		return
	}

	touristID := context.GetInt64("touristID")
	log.Printf("touristID from context: %d", touristID)

	trip.TouristID = int(touristID)
	tc.repo.SaveTrip(&trip)
	context.JSON(http.StatusOK, gin.H{"message": "trip saved"})
}

