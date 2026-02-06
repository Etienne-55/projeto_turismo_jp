package websocket

import (
    "encoding/json"
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
}

