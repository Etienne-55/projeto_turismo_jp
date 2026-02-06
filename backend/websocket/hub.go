package websocket

import (
    "sync"
)

type Hub struct {
    clients map[*Client]bool
    
    // Channel for messages to broadcast to everyone
    broadcast chan []byte
    
    // Channel to register new clients
    register chan *Client
    
    // Channel to unregister disconnected clients
    unregister chan *Client
    
    // Mutex to protect the clients map (prevent race conditions)
    mutex sync.Mutex
}

func NewHub() *Hub {
    return &Hub{
        clients:    make(map[*Client]bool),
        broadcast:  make(chan []byte),
        register:   make(chan *Client),
        unregister: make(chan *Client),
    }
}

func (h *Hub) Run() {
    for {
        select {
        case client := <-h.register:
            h.mutex.Lock()
            h.clients[client] = true
            h.mutex.Unlock()
            
        case client := <-h.unregister:
            h.mutex.Lock()
            if _, ok := h.clients[client]; ok {
                delete(h.clients, client)
                close(client.send)
            }
            h.mutex.Unlock()
            
        case message := <-h.broadcast:
            h.mutex.Lock()
            for client := range h.clients {
                select {
                case client.send <- message:
                default:
                    close(client.send)
                    delete(h.clients, client)
                }
            }
            h.mutex.Unlock()
        }
    }
}

