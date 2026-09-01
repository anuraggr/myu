package tui

import (
    tea "github.com/charmbracelet/bubbletea"
)

func (m listModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.WindowSizeMsg:
        m.list.SetSize(msg.Width-4, msg.Height-4)
        return m, nil

    case tea.KeyMsg:
        switch msg.String() {
        case "q", "ctrl+c":
            m.chosen = -1
            return m, tea.Quit
        case "enter":
            m.chosen = m.list.Index()
            return m, tea.Quit
        }
    }

    var cmd tea.Cmd
    m.list, cmd = m.list.Update(msg)
    return m, cmd
}
