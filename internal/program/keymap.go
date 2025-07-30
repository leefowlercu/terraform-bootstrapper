package program

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
)

// programKeyMap defines a Program-level KeyMap for the application and
// implements the help.KeyMap interface.
type programKeyMap struct {
	Up   key.Binding
	Down key.Binding
	Quit key.Binding
	Help key.Binding
}

// Force implementation of help.KeyMap Interface
var _ help.KeyMap = (*programKeyMap)(nil)

// Assign Keys and Help Messages to Program Functionality
var defaultProgramKeyMap = programKeyMap{
	Up: key.NewBinding(
		key.WithKeys("up"),
		key.WithHelp("↑", "up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down"),
		key.WithHelp("↓", "down"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c"),
		key.WithHelp("ctrl+c", "quit"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "more"),
	),
}

func (k programKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Up, k.Down, k.Quit, k.Help}
}

func (k programKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down, k.Quit, k.Help},
	}
}
