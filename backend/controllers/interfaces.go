package controllers

import "projeto_turismo_jp/repositories"



type TouristController struct {
	repo repositories.TouristRepository
}

func NewTouristController(repo repositories.TouristRepository) *TouristController {
	return &TouristController{
		repo: repo,
	}
}
