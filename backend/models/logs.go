package models

import "time"


type Notification struct {
		ID				int				`json:"id"`
    Type      string    `json:"type"`      
    Message   string    `json:"message"`  
    Data      any       `json:"data"`     
    Timestamp time.Time `json:"timestamp"`
}
