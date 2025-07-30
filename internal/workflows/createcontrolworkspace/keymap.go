package createcontrolworkspace

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
)

type createControlWorkspaceKeyMap struct {
	Up     key.Binding
	Down   key.Binding
	CycleF key.Binding
	CycleB key.Binding
	Submit key.Binding
}

var _ help.KeyMap = (*createControlWorkspaceKeyMap)(nil)

var defaultCreateControlWorkspaceKeyMap = createControlWorkspaceKeyMap{
	Up: key.NewBinding(
		key.WithKeys("up"),
		key.WithHelp("↑", "up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down"),
		key.WithHelp("↓", "down"),
	),
	CycleF: key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("tab", "cycle forward"),
	),
	CycleB: key.NewBinding(
		key.WithKeys("shift+tab"),
		key.WithHelp("shift-tab", "cycle backward"),
	),
	Submit: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "next/submit"),
	),
}

func (k createControlWorkspaceKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.CycleF, k.CycleB, k.Submit}
}

func (k createControlWorkspaceKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.CycleF, k.CycleB, k.Submit},
	}
}
