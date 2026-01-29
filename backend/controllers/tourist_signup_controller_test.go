package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"projeto_turismo_jp/models"
	"projeto_turismo_jp/server"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)


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

func TestSignup(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockRepo := &MockTouristRepository{
		SaveFunc: func(tourist *models.Tourist) error {
			tourist.ID = 123
			return nil

		},
	}

	controller := NewTouristController(mockRepo)
	server := server.SetupServer()
	server.POST("/signup", controller.Signup)

	exampleTourist := models.Tourist{
		Email: "test@gmail.com",
		Password: "testpassword1234",
	}
	
	userJson, _ := json.Marshal(exampleTourist)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/signup", bytes.NewBuffer(userJson))
	req.Header.Set("Content-Type", "application.json")

	server.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "user created successfully")
}

