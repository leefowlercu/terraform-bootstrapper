package executeworkflow

import (
	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/leefowlercu/terraform-bootstrapper/internal/messages"
	"github.com/leefowlercu/terraform-bootstrapper/internal/stages"
	"github.com/leefowlercu/terraform-bootstrapper/internal/styles"
	"github.com/leefowlercu/terraform-bootstrapper/internal/workflows"
)

const maxDescriptionWidth = 100

// Define the Model for this Stage
type model struct {
	workflow   workflows.Workflow
	keys       help.KeyMap
	viewWidth  int
	viewHeight int
}

// Compile-time validation of implementation of the Stage interface
var _ stages.Stage = (*model)(nil)

func New(workflow workflows.Workflow) *model {
	return &model{
		workflow: workflow,
		keys:     defaultExecuteWorkflowKeyMap,
	}
}

func (m model) Init() tea.Cmd {
	return m.workflow.Init()
}

func (m model) Update(msg tea.Msg) (stages.Stage, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case messages.AvailableSizeMsg:
		m.viewWidth = msg.Width
		m.viewHeight = msg.Height
	}

	m.workflow, cmd = m.workflow.Update(msg)

	return m, cmd
}

func (m model) View() string {
	title := styles.WorkflowTitle.Render(m.workflow.Title()+" Workflow...") + "\n"
	description := styles.WorkflowDescription.Width(maxDescriptionWidth).Render(m.workflow.LongDescription()) + "\n"
	workflowView := m.workflow.View()

	header := lipgloss.JoinVertical(lipgloss.Left, title, description)
	headerHeight := lipgloss.Height(header)

	remainingHeight := max(m.viewHeight-headerHeight, 0)

	// Use a container to manage the height of the workflow view
	workflowContainer := lipgloss.NewStyle().Height(remainingHeight).Render(workflowView)

	return lipgloss.JoinVertical(lipgloss.Left, header, workflowContainer)
}

func (m model) KeyMaps() (help.KeyMap, help.KeyMap) {
	return m.keys, m.workflow.KeyMap()
}
