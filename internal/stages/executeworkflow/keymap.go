package executeworkflow

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
)

type executeWorkflowKeyMap struct{}

var _ help.KeyMap = (*executeWorkflowKeyMap)(nil)

var defaultExecuteWorkflowKeyMap = executeWorkflowKeyMap{}

func (k executeWorkflowKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{}
}

func (k executeWorkflowKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{}
}
