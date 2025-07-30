package stages

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
	// and a command. A Stage can return itself to continue being displayed.
	Update(msg tea.Msg) (Stage, tea.Cmd)

	// View renders the Stage's UI.
	View() string

	// KeyMaps returns any Key Bindings for this Stage and, if applicable
	// the KeyMap for the currently executing Workflow.
	KeyMaps() (help.KeyMap, help.KeyMap)
}
