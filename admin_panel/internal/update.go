package internal

import (
	"time"
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "c":
			m.notifications = []Notification{}
			return m, nil

		case "r":
			if !m.connected {
				m.status = "Reconnecting..."
				return m, ConnectWebSocket
			}
			return m, nil
		}
	
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil
	
	case ConnectedMsg:
		m.connected = true
		m.status = "Connected ✓"
		return m, nil
	
	case DisconnectedMsg:
		m.connected = false
		m.status = "Disconnected ✗ (press 'r' to reconnect)"
		return m, tea.Tick(time.Second*3, func(t time.Time) tea.Msg {
			return reconnectMsg{}
		})
	
	case reconnectMsg:
		m.status = "Reconnecting..."
		return m, ConnectWebSocket
	
	case NotificationMsg:
		m.notifications = append([]Notification{msg.Notification}, m.notifications...)
		
		if len(m.notifications) > 50 {
			m.notifications = m.notifications[:50]
		}
		return m, nil
	
	case ErrorMsg:
		m.status = "Error: " + msg.Err.Error()
		return m, nil
	}
	
	return m, nil
}

