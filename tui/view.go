package tui

import (
    "github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

func (m listModel) View() string {
    return docStyle.Render(m.list.View())
}
