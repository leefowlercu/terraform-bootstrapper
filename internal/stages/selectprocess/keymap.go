package selectprocess

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
)

type selectProcessKeyMap struct {
	Filter    key.Binding
	GoToStart key.Binding
	GoToEnd   key.Binding
}

var _ help.KeyMap = (*selectProcessKeyMap)(nil)

var defaultSelectProcessKeyMap = selectProcessKeyMap{
	Filter: key.NewBinding(
		key.WithKeys("/"),
		key.WithHelp("/", "filter"),
	),
	GoToStart: key.NewBinding(
		key.WithKeys("g"),
		key.WithHelp("g", "go to start"),
	),
	GoToEnd: key.NewBinding(
		key.WithKeys("G"),
		key.WithHelp("G", "go to end"),
	),
}

func (k selectProcessKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Filter}
}

func (k selectProcessKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Filter, k.GoToStart, k.GoToEnd},
	}
}
