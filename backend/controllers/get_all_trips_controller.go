package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)


// Signup godoc
// @Summary      Get all trips registered on the database
// @Tags         Trip
// @Accept       json
// @Success      201 {object} map[string]string "Database data"
// @Failure      400 {object} map[string]string "Invalid request"
// @Failure      500 {object} map[string]string "Internal server error"
// @Router       /get_all_trips [get]
func (tc *TripController) GetAllTrips(context *gin.Context) {
	trips, err := tc.repo.GetAllTrips()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch events"})
		log.Printf("error: %v", err)
		return
	}
	context.JSON(http.StatusOK, trips)
}

