package tui

import (
    "fmt"

    "github.com/anuraggr/myu/youtube"
)

type listItem struct {
    result youtube.SearchResult
}

func (i listItem) Title() string { return i.result.Title }

func (i listItem) Description() string {
    dur := "?"
    if i.result.Duration != nil {
        dur = fmt.Sprintf("%d:%02d", *i.result.Duration/60, *i.result.Duration%60)
    }
    return fmt.Sprintf("%s · %s", i.result.Channel, dur)
}

func (i listItem) FilterValue() string { return i.result.Title }
