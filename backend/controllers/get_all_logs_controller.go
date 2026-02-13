package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)


func (lg *LogController) GetAllLogs(context *gin.Context) {
	logs, err := lg.repo.GetAllLogs()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch logs"})
		log.Printf("error: %v", err)
		return
	}
	context.JSON(http.StatusOK, logs)
}

