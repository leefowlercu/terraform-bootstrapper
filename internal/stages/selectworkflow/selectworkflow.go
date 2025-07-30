package selectworkflow

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/leefowlercu/terraform-bootstrapper/internal/messages"
	"github.com/leefowlercu/terraform-bootstrapper/internal/stages"
	"github.com/leefowlercu/terraform-bootstrapper/internal/stages/executeworkflow"
	"github.com/leefowlercu/terraform-bootstrapper/internal/workflows"
	"github.com/leefowlercu/terraform-bootstrapper/internal/workflows/createcontrolworkspace"
)

// Define the Model for this Stage
type model struct {
	workflowList list.Model
	keys         selectWorkflowKeyMap
}

// Compile-time validation of implementation of the Stage interface
var _ stages.Stage = (*model)(nil)

// Creates and returns an initial Model for this Stage
func New() *model {
	createControlWorkspaceWorkflow := createcontrolworkspace.New()

	itemList := []list.Item{createControlWorkspaceWorkflow}

	// Initialize with zero values, the AvailableSizeMsg will set the correct size.
	workflowList := list.New(itemList, list.NewDefaultDelegate(), 0, 0)
	workflowList.Title = "Select a Bootstrapping Workflow..."
	workflowList.SetShowTitle(true)
	workflowList.SetFilteringEnabled(true)
	workflowList.SetShowHelp(false)

	return &model{
		workflowList: workflowList,
		keys:         defaultSelectWorkflowKeyMap,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (stages.Stage, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case messages.AvailableSizeMsg:
		m.workflowList.SetSize(msg.Width, msg.Height)
		return m, nil
	case tea.WindowSizeMsg:
		m.workflowList, cmd = m.workflowList.Update(msg)
		return m, cmd
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Select):
			selectedWorkflow := m.workflowList.SelectedItem().(workflows.Workflow)

			// Return a tea.Cmd that will emit the messages.ChangeStageMsg
			return m, func() tea.Msg {
				return messages.ChangeStageMsg{
					Stage: executeworkflow.New(selectedWorkflow),
				}
			}
		}
	}

	m.workflowList, cmd = m.workflowList.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return m.workflowList.View()
}

func (m model) KeyMaps() (help.KeyMap, help.KeyMap) {
	return m.keys, nil
}
