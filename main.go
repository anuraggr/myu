package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	searchMode := flag.Bool("s", false, "search-picker mode")
	interactiveMode := flag.Bool("i", false, "interactive mode (TUI)")
	help := flag.Bool("h", false, "show help")
	flag.Parse()

	if *help {
		fmt.Println("Usage: myu [options] [query]")
		flag.PrintDefaults()
		return
	}

	query := strings.Join(flag.Args(), " ")

	if *searchMode && query == "" {
		fmt.Fprintln(os.Stderr, "Error: -s requires a search query")
		os.Exit(1)
	}

	if *interactiveMode || query == "" {
		fmt.Fprintf(os.Stderr, "TUI not implemented yet (query: %s)\n", query)
		return
	}

	count := 1
	if *searchMode {
		count = 10
	}

	results, err := FetchSearchResults(query, count)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
	if len(results) == 0 {
		fmt.Fprintf(os.Stderr, "No results for '%s'\n", query)
		os.Exit(1)
	}

	choice := 0
	if *searchMode {
		choice, err = pickResult(results, query)
		if err != nil {
			os.Exit(0)
		}
	}

	fmt.Printf("Now playing: %s\n", results[choice].Title)
	if err := PlayByID(results[choice].ID); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}
