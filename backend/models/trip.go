package models

import(
	"time"
)

type Trip struct {
	ID 								int64	 		`json:"id"`
	LodgingLocation 	string 		`json:"lodging_location" binding:"required"`
	TripDescription 	string 		`json:"trip_description binding:required"`
	ArrivalDate 			string		`json:"arrival_date" binding:"required"`
	DepartureDate 		string 		`json:"departure_date" binding:"required"`
	TripReview				string		`json:"trip_review"`
  Status          	string    `json:"status"`
  TouristID       	int       `json:"tourist_id"`
	CreatedAt 				time.Time 
	UpdatedAt 				time.Time 
}

