package main

import (
	"fmt"
	"os"

	"github.com/sansmoraxz/toi-go/pkg/ui"
)

func main() {
	// Initialize the game and UI
	ui := ui.NewUI()

	// Start the UI
	err := ui.Start()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to start UI: %v\n", err)
		os.Exit(1)
	}
}
