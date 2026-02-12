package websocket

import (
	"encoding/json"
	"log"
	"projeto_turismo_jp/db"
	"time"
)

type Notification struct {
    Type      string    `json:"type"`      
    Message   string    `json:"message"`  
    Data      any       `json:"data"`     
    Timestamp time.Time `json:"timestamp"`
}

func (h *Hub) SendNotification(notifType, message string, data any) {
    notification := Notification{
        Type:      notifType,
        Message:   message,
        Data:      data,
        Timestamp: time.Now(),
    }
    
    jsonData, err := json.Marshal(notification)
    if err != nil {
        return
    }
    
    h.broadcast <- jsonData

		err = saveNotificationToLog(notification);
		if err != nil {
		log.Printf("error: %v", err)
		return
	}
	
	log.Printf("log saved to the database")
}

func saveNotificationToLog(n Notification) error {
	dataJSON, _ := json.Marshal(n.Data)

	query := `INSERT INTO notification_logs (type, message, data, timestamp)
						VALUES (?, ?, ?, ?)`

	db.DB.Exec(query, n.Type, n.Message, string(dataJSON), n.Timestamp)

	return nil
}

