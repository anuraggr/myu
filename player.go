package main

import (
	"fmt"
	"os"
	"os/exec"
)

func PlayByID(videoID string) error {
	if videoID == "" {
		return fmt.Errorf("empty video id")
	}
	cmd := exec.Command("mpv", "--no-video", "https://youtu.be/"+videoID)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func PlayByIDAsync(videoID string) (*exec.Cmd, error) {
	cmd := exec.Command("mpv", "--no-terminal", "https://youtu.be/"+videoID)
	if err := cmd.Start(); err != nil {
		return nil, err
	}
	return cmd, nil
}
