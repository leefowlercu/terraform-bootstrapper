package program

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/leefowlercu/terraform-bootstrapper/internal/keymap"
	"github.com/leefowlercu/terraform-bootstrapper/internal/stage"
	"github.com/leefowlercu/terraform-bootstrapper/internal/stages/selectprocess"
	"github.com/leefowlercu/terraform-bootstrapper/internal/styles"
)

// Define the Model for the Program
type model struct {
	globalKeys   keymap.GlobalKeyMap
	currentStage stage.Stage
	help         help.Model
}

// Creates and returns an initial Model for the Program
func New() model {
	return model{
		globalKeys:   keymap.DefaultGlobalKeyMap,
		currentStage: selectprocess.New(),
		help:         help.New(),
	}
}

func (m model) Init() tea.Cmd {
	return m.currentStage.Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.help.Width = msg.Width

	case tea.KeyMsg:
		switch {
		// Check for global keybindings
		case key.Matches(msg, m.globalKeys.Quit):
			return m, tea.Quit
		case key.Matches(msg, m.globalKeys.Help):
			m.help.ShowAll = !m.help.ShowAll
			return m, nil
		}
	}

	var cmd tea.Cmd
	m.currentStage, cmd = m.currentStage.Update(msg)

	return m, cmd
}

func (m model) View() string {
	// Combine KeyMaps to show both Global and Stage-specific keys
	combinedKeyMap := keymap.CombinedKeyMap{
		Global: m.globalKeys,
		Stage:  m.currentStage.KeyMap(),
	}

	stageView := m.currentStage.View()
	helpView := m.help.View(combinedKeyMap)

	return styles.Program.Render(lipgloss.JoinVertical(lipgloss.Left, stageView, helpView))
}
