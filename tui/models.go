package tui

import (
    "github.com/charmbracelet/bubbles/list"
    tea "github.com/charmbracelet/bubbletea"

    "github.com/anuraggr/myu/youtube"
)

type listModel struct {
    list   list.Model
    chosen int // -1 = user quit without picking
}

func newListModel(results []youtube.SearchResult, query string) listModel {
    items := make([]list.Item, len(results))
    for i, r := range results {
        items[i] = listItem{result: r}
    }

    l := list.New(items, list.NewDefaultDelegate(), 0, 0)
    l.Title = "Results: " + query
    l.SetShowStatusBar(false)

    return listModel{list: l, chosen: -1}
}

func (m listModel) Init() tea.Cmd { return nil }
