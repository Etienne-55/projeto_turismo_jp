package repositories

import (
	"projeto_turismo_jp/db"
	"projeto_turismo_jp/models"
)


func (r *tripRepositoryImpl) SaveTrip(t *models.Trip) error {
	query := `
	INSERT INTO trip(lodging_location, trip_description, arrival_date, departure_date)
	VALUES (?, ?, ?, ?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	
	defer stmt.Close()
	result, err := stmt.Exec(t.LodgingLocation, t.TripDescription, t.ArrivalDate, t.DepartureDate)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	t.ID = id
	// events = append(events, *t)
	return err

}

