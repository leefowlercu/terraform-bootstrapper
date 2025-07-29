// Package stage defines the contract for a single application stage.
package stage

import (
	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
)

// Stage is an interface that represents a single view or state
// in the application.
type Stage interface {
	// Init is the first function that will be called. It returns a command to be executed.
	Init() tea.Cmd

	// Update is called when a message is received. It returns the next Stage to be rendered
	// and a command. A stage can return itself to continue being displayed.
	Update(msg tea.Msg) (Stage, tea.Cmd)

	// View renders the Stage's UI.
	View() string

	// KeyMap returns the Key Bindings for this Stage.
	KeyMap() help.KeyMap
}
