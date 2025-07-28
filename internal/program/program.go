package program

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/leefowlercu/terraform-bootstrapper/internal/stage"
	"github.com/leefowlercu/terraform-bootstrapper/internal/stages/selectprocess"
)

// Define global functionality to bind Keys to
type globalKeyMap struct {
	Up     key.Binding
	Down   key.Binding
	Select key.Binding
	Exit   key.Binding
	Help   key.Binding
}

// Define the Model for the Program
type model struct {
	globalKeys   *globalKeyMap
	currentStage stage.Stage
	help         help.Model
}

// Assign Keys and Help Messages to Global functions
var defaultGlobalKeyMap = globalKeyMap{
	Up: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "down"),
	),
	Select: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "select"),
	),
	Exit: key.NewBinding(
		key.WithKeys("esc", "ctrl+c", "q"),
		key.WithHelp("esc/q", "exit"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
}

// Force implementation of help.KeyMap Interface
var _ help.KeyMap = (*globalKeyMap)(nil)

// Creates and returns an initial Model for the Program
func New() model {
	return model{
		globalKeys:   &defaultGlobalKeyMap,
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
		case key.Matches(msg, m.globalKeys.Exit):
			return m, tea.Quit
		case key.Matches(msg, m.globalKeys.Help):
			m.help.ShowAll = !m.help.ShowAll
			return m, nil
		}
	}

	var cmd tea.Cmd
	_, cmd = m.currentStage.Update(msg)

	return m, cmd
}

func (m model) View() string {
	stageView := m.currentStage.View()
	helpView := m.help.View(m.globalKeys)

	return lipgloss.JoinVertical(lipgloss.Left, stageView, helpView)
}

func (k globalKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Exit}
}

func (k globalKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down, k.Select}, // First Column
		{k.Help, k.Exit},         // Second Column
	}
}
