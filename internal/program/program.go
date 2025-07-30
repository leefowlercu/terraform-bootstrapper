package program

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/leefowlercu/terraform-bootstrapper/internal/keymap"
	"github.com/leefowlercu/terraform-bootstrapper/internal/messages"
	"github.com/leefowlercu/terraform-bootstrapper/internal/stages"
	"github.com/leefowlercu/terraform-bootstrapper/internal/stages/selectworkflow"
	"github.com/leefowlercu/terraform-bootstrapper/internal/styles"
)

// Define the Model for the Program
type model struct {
	keys         programKeyMap
	currentStage stages.Stage
	help         help.Model
	viewWidth    int
	viewHeight   int
}

// Creates and returns an initial Model for the Program
func New() *model {
	return &model{
		keys:         defaultProgramKeyMap,
		currentStage: selectworkflow.New(),
		help:         help.New(),
	}
}

func (m model) Init() tea.Cmd {
	return m.currentStage.Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.viewWidth = msg.Width
		m.viewHeight = msg.Height
		m.help.Width = msg.Width
		m.currentStage, cmd = m.currentStage.Update(msg)
		cmds = append(cmds, cmd, m.sendAvailableSizeCmd(msg.Width, msg.Height))

	case messages.ChangeStageMsg:
		// Switch the current Stage based on the message
		m.currentStage = msg.Stage

		// Initialize the new Stage and send a tea.Cmd to update the available size
		cmds = append(cmds, m.currentStage.Init(), m.sendAvailableSizeCmd(m.viewWidth, m.viewHeight))

	case tea.KeyMsg:
		switch {
		// Check for global keybindings
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
			cmds = append(cmds, m.sendAvailableSizeCmd(m.viewWidth, m.viewHeight))
		default:
			m.currentStage, cmd = m.currentStage.Update(msg)
			cmds = append(cmds, cmd)
		}

	default:
		m.currentStage, cmd = m.currentStage.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	stageView := m.currentStage.View()
	helpView := m.help.View(m.getCombinedKeyMap())

	return styles.Program.Render(lipgloss.JoinVertical(lipgloss.Left, stageView, helpView))
}

func (m model) sendAvailableSizeCmd(width, height int) tea.Cmd {
	return func() tea.Msg {
		helpView := m.help.View(m.getCombinedKeyMap())
		helpHeight := lipgloss.Height(helpView)

		availableHeight := height - helpHeight - 2 // Subtract 2 for the vertical padding

		return messages.AvailableSizeMsg{
			Width:  width,
			Height: availableHeight,
		}
	}
}

func (m model) getCombinedKeyMap() help.KeyMap {
	// Dynamically update help text based on toggled state
	if m.help.ShowAll {
		m.keys.Help.SetHelp("?", "less")
	} else {
		m.keys.Help.SetHelp("?", "more")
	}

	// Grab the KeyMaps from the current Stage and, if applicable, the Workflow
	stageKeys, workflowKeys := m.currentStage.KeyMaps()

	// Combine KeyMaps to show both Global and Stage-specific keys
	return keymap.CombinedKeyMap{
		Global:   m.keys,
		Stage:    stageKeys,
		Workflow: workflowKeys,
	}
}
