package internal

import (
	"fmt"
	"strings"
	"time"
	
	"github.com/charmbracelet/lipgloss"
)

func (m model) View() string {
	var s strings.Builder
	
	header := headerStyle.Render("Admin Dashboard")
	s.WriteString(header)
	s.WriteString("\n\n")
	
	statusText := m.status
	var statusStyled string
	if m.connected {
		statusStyled = connectedStyle.Render(statusText)
	} else {
		statusStyled = disconnectedStyle.Render(statusText)
	}
	s.WriteString("Status: " + statusStyled)
	s.WriteString("\n")
	
	stats := fmt.Sprintf("Total Notifications: %d", len(m.notifications))
	s.WriteString(stats)
	s.WriteString("\n")

	s.WriteString(strings.Repeat("─", m.width))
	s.WriteString("\n")
	
	if len(m.notifications) == 0 {
		s.WriteString("\n")
		s.WriteString(lipgloss.NewStyle().
			Foreground(mutedColor).
			Italic(true).
			Render("  No notifications yet. Waiting for activity..."))
		s.WriteString("\n")
	} else {
		maxDisplay := 10
		if len(m.notifications) < maxDisplay {
			maxDisplay = len(m.notifications)
		}
		
		for i := 0; i < maxDisplay; i++ {
			notif := m.notifications[i]
			s.WriteString(renderNotification(notif))
		}
	}
	
	s.WriteString("\n")
	s.WriteString(strings.Repeat("─", m.width))
	s.WriteString("\n")
	// help := helpStyle.Render("Press 'c' to clear | 'q' to quit")
	// s.WriteString(help)
	s.WriteString(renderHelp(m.connected))
	
	return s.String()
}

func renderHelp(connected bool) string {
	if connected {
		return helpStyle.Render("Press 'c' to clear | 'q' to quit")
	} else {
		return helpStyle.Render("Press 'r' to reconnect | 'q' to quit")
	}
}

func renderNotification(n Notification) string {
	color := getNotificationColor(n.Type)
	icon := getNotificationIcon(n.Type)
	
	box := notificationBoxStyle.
		BorderForeground(color)
	
	typeStyled := notificationTypeStyle.
		Foreground(color).
		Render(icon + " " + formatType(n.Type))
	
	messageStyled := notificationMessageStyle.Render(n.Message)
	
	timeAgo := formatTimeAgo(n.Timestamp)
	timeStyled := notificationTimeStyle.Render(timeAgo)
	
	content := fmt.Sprintf("%s\n%s\n%s", typeStyled, messageStyled, timeStyled)
	
	return box.Render(content)
}

func formatType(t string) string {
	types := map[string]string{
		"trip_created": "Trip Created",
		"trip_updated": "Trip Updated",
		"trip_deleted": "Trip Deleted",
		"user_signup":  "New User",
		"stats_update": "Statistics",
	}
	if formatted, ok := types[t]; ok {
		return formatted
	}
	return strings.ReplaceAll(t, "_", " ")
}

func formatTimeAgo(t time.Time) string {
	diff := time.Since(t)
	
	if diff.Seconds() < 60 {
		return fmt.Sprintf("%d seconds ago", int(diff.Seconds()))
	} else if diff.Minutes() < 60 {
		mins := int(diff.Minutes())
		if mins == 1 {
			return "1 minute ago"
		}
		return fmt.Sprintf("%d minutes ago", mins)
	} else if diff.Hours() < 24 {
		hours := int(diff.Hours())
		if hours == 1 {
			return "1 hour ago"
		}
		return fmt.Sprintf("%d hours ago", hours)
	} else {
		return t.Format("Jan 02, 15:04")
	}
}

