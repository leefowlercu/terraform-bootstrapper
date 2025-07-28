package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/leefowlercu/terraform-bootstrapper/internal/program"
)

func main() {
	program := tea.NewProgram(program.New(), tea.WithAltScreen())

	if _, err := program.Run(); err != nil {
		log.Fatal(err)
	}
}
