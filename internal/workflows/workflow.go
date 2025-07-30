package workflows

import (
	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
)

// Workflow is an interface that represents a single 'workflow' or process
// that the application can execute.
type Workflow interface {
	// Init is the first function that will be called. It returns a command to be executed.
	Init() tea.Cmd

	// Update is called when a message is received. It returns the next Workflow to be rendered
	// and a command. A Workflow can return itself to continue being displayed.
	Update(msg tea.Msg) (Workflow, tea.Cmd)

	// View renders the Workflow's UI.
	View() string

	// KeyMap returns any Key Bindings for this Workflow.
	KeyMap() help.KeyMap

	// Identifier returns a unique identifier for the Workflow.
	Identifier() string

	// Title returns the human-readable title of the Workflow.
	Title() string

	// Description returns a brief description of the Workflow.
	Description() string

	// LongDescription returns a detailed description of the Workflow.
	LongDescription() string

	// FilterValue returns a string used for filtering in lists.
	FilterValue() string
}
