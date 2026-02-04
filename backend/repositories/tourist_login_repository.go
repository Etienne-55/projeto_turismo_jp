package repositories

import (
	"errors"
	"projeto_turismo_jp/db"
	"projeto_turismo_jp/models"
	"projeto_turismo_jp/utils"
)


func (r *touristRepositoryImpl) ValidateCredentials(t *models.Tourist) error {
	query := "SELECT id, password, role FROM tourist WHERE email = ?"
	row := db.DB.QueryRow(query, t.Email)

	var retreivedPassword string
	err := row.Scan(&t.ID, &retreivedPassword, &t.Role)
	if err != nil {
		return errors.New("error retreiving user data")
	}

	passwordIsValid := utils.CheckPasswordHash(t.Password, retreivedPassword)
	if !passwordIsValid {
		return errors.New("passwords dont match")
	}

	return nil 
}

