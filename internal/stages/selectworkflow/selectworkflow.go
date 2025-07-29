package selectworkflow

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/leefowlercu/terraform-bootstrapper/internal/messages"
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

	// Initialize with zero values, the AvailableSizeMsg will set the correct size.
	workflowList := list.New(itemList, list.NewDefaultDelegate(), 0, 0)
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
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case messages.AvailableSizeMsg:
		m.workflowList.SetSize(msg.Width, msg.Height)
		return m, nil
	case tea.WindowSizeMsg:
		m.workflowList, cmd = m.workflowList.Update(msg)
		return m, cmd
	}

	m.workflowList, cmd = m.workflowList.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return m.workflowList.View()
}

func (m model) KeyMap() help.KeyMap {
	return m.stageKeys
}
