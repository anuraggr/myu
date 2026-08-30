package main

import "fmt"

func pickResult(results []SearchResult, query string) (int, error) {
	fmt.Printf("Results for '%s':\n\n", query)

	for i, r := range results {
		dur := "?"
		if r.Duration != nil {
			dur = fmt.Sprintf("%d:%02d", *r.Duration/60, *r.Duration%60)
		}
		fmt.Printf("  %2d) [%s] %s\n", i+1, dur, r.Title)
	}

	fmt.Printf("\nPlay which? [1-%d, 0=quit]: ", len(results))

	var choice int
	if _, err := fmt.Scanf("%d", &choice); err != nil {
		return 0, err
	}
	if choice == 0 {
		return 0, fmt.Errorf("cancelled")
	}
	if choice < 1 || choice > len(results) {
		return 0, fmt.Errorf("invalid choice")
	}
	return choice - 1, nil
}
