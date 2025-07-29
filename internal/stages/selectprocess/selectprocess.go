package selectprocess

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/leefowlercu/terraform-bootstrapper/internal/process"
	"github.com/leefowlercu/terraform-bootstrapper/internal/stage"
)

// Define the Model for this Stage
type model struct {
	processList list.Model
	stageKeys   selectProcessKeyMap
}

// Force implementation of Stage Interface
var _ stage.Stage = (*model)(nil)

// Creates and returns an initial Model for this Stage
func New() stage.Stage {
	createControlWorkspaceItem := process.New(
		"create-control-workspace",
		"Create Control Workspace",
		"Creates a Control Workspace and associated Project",
	)

	itemList := []list.Item{createControlWorkspaceItem}
	processList := list.New(itemList, list.NewDefaultDelegate(), 80, 50)
	processList.Title = "Select a Bootstrapping Process..."
	processList.SetShowTitle(true)
	processList.SetFilteringEnabled(true)
	processList.SetShowHelp(false)

	return model{
		processList: processList,
		stageKeys:   defaultSelectProcessKeyMap,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (stage.Stage, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.processList.SetSize(msg.Width, msg.Height-2)
	}

	var cmd tea.Cmd
	m.processList, cmd = m.processList.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return m.processList.View()
}

func (m model) KeyMap() help.KeyMap {
	return m.stageKeys
}
