package models

import (
)


type Tourist struct {
	ID int64
	Email string `binding:"required"`
	Password string `binding:"required"`
	IsAdmin bool 
}

