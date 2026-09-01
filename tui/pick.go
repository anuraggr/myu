package tui

import (
    "fmt"

    tea "github.com/charmbracelet/bubbletea"

    "github.com/anuraggr/myu/youtube"
)

func PickResult(results []youtube.SearchResult, query string) (int, error) {
    p := tea.NewProgram(newListModel(results, query))
    final, err := p.Run()
    if err != nil {
        return 0, err
    }

    m, ok := final.(listModel)
    if !ok || m.chosen < 0 {
        return 0, fmt.Errorf("cancelled")
    }
    return m.chosen, nil
}
