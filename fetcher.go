package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
)

type SearchResult struct {
	Title    string `json:"title"`
	ID       string `json:"id"`
	Duration *int   `json:"duration"`
	Channel  string `json:"channel"`
}

func FetchSearchResults(query string, count int) ([]SearchResult, error) {
	if count < 1 {
		count = 1
	}
	if count > 50 {
		count = 50
	}

	cmd := exec.Command("yt-dlp",
		"--dump-single-json",
		"--flat-playlist",
		"--no-playlist",
		"--no-warnings",
		"--socket-timeout", "10",
		fmt.Sprintf("ytsearch%d:%s", count, query),
	)

	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = nil

	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("yt-dlp failed: %w", err)
	}

	var raw struct {
		Entries []SearchResult `json:"entries"`
	}
	if err := json.Unmarshal(stdout.Bytes(), &raw); err != nil {
		return nil, fmt.Errorf("parse json: %w", err)
	}
	return raw.Entries, nil
}
