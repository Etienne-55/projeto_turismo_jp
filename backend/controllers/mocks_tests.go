package controllers

import "projeto_turismo_jp/models"


type MockTouristRepository struct {
	SaveFunc func(tourist *models.Tourist) error
	ValidateCredentialsFunc func(tourist *models.Tourist) error
}

func (m *MockTouristRepository) Save(tourist *models.Tourist) error {
	if m.SaveFunc != nil {
		return m.SaveFunc(tourist)
	}
	tourist.ID = 1
	return nil
}

func (m *MockTouristRepository) ValidateCredentials(tourist *models.Tourist) error {
	// if m.SaveFunc != nil {
	// 	return m.SaveFunc(tourist)
	// }
	// tourist.ID = 1
	return nil
}
