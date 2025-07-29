package selectworkflow

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/leefowlercu/terraform-bootstrapper/internal/process"
	"github.com/leefowlercu/terraform-bootstrapper/internal/stage"
)

// Define the Model for this Stage
type model struct {
	workflowList list.Model
	stageKeys    selectWorkflowKeyMap
}

// Force implementation of Stage Interface
var _ stage.Stage = (*model)(nil)

// Creates and returns an initial Model for this Stage
func New() *model {
	createControlWorkspaceItem := process.New(
		"create-control-workspace",
		"Create Control Workspace",
		"Creates a Control Workspace and associated Project",
	)

	itemList := []list.Item{createControlWorkspaceItem}
	workflowList := list.New(itemList, list.NewDefaultDelegate(), 80, 50)
	workflowList.Title = "Select a Bootstrapping Workflow..."
	workflowList.SetShowTitle(true)
	workflowList.SetFilteringEnabled(true)
	workflowList.SetShowHelp(false)

	return &model{
		workflowList: workflowList,
		stageKeys:    defaultSelectWorkflowKeyMap,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (stage.Stage, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.workflowList.SetSize(msg.Width, msg.Height-2)
	}

	var cmd tea.Cmd
	m.workflowList, cmd = m.workflowList.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return m.workflowList.View()
}

func (m model) KeyMap() help.KeyMap {
	return m.stageKeys
}
