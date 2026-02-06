package repositories

import (
	"log"
	"projeto_turismo_jp/db"
	"projeto_turismo_jp/models"
)


func (r *tripRepositoryImpl) AddReview(t *models.Trip) error {
	query := "UPDATE trip SET trip_review = ? WHERE id = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		log.Printf("error: %v", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(t.TripReview, t.ID)

	return err
}

