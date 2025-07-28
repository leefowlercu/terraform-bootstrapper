package selectprocess

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/leefowlercu/terraform-bootstrapper/internal/stage"
)

// Force implementation of Stage Interface
var _ stage.Stage = (*model)(nil)

// Define the Model for this Stage
type model struct {
	message string
}

// Creates and returns an initial Model for this Stage
func New() stage.Stage {
	return model{
		message: "Hello, World!",
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (stage.Stage, tea.Cmd) {
	return m, nil
}

func (m model) View() string {
	return m.message
}
