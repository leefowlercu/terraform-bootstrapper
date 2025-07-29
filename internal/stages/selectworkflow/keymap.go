package selectworkflow

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
)

type selectWorkflowKeyMap struct {
	Filter    key.Binding
	GoToStart key.Binding
	GoToEnd   key.Binding
}

var _ help.KeyMap = (*selectWorkflowKeyMap)(nil)

var defaultSelectWorkflowKeyMap = selectWorkflowKeyMap{
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

func (k selectWorkflowKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Filter}
}

func (k selectWorkflowKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Filter, k.GoToStart, k.GoToEnd},
	}
}
