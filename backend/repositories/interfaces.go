package repositories

import (
	"database/sql"
	"projeto_turismo_jp/models"
)


type TouristRepository interface {
	Save(tourist *models.Tourist) error
	ValidateCredentials(tourist *models.Tourist) error
}

type touristRepositoryImpl struct {
	db *sql.DB
}

func NewTouristRepository(db *sql.DB) TouristRepository {
	return &touristRepositoryImpl{
		db: db,
	}
}

