package repositories

import (
	"projeto_turismo_jp/db"
	"projeto_turismo_jp/models"
	"projeto_turismo_jp/utils"
)


func Save(t *models.Tourist) error {
	query := "INSERT INTO tourist(email, password) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(t.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(t.Email, hashedPassword)
	if err != nil {
		return err
	}
	
	touristID, err := result.LastInsertId()

	t.ID = touristID
	return err

}
