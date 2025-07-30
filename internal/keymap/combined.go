package keymap

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
)

type CombinedKeyMap struct {
	Global   help.KeyMap
	Stage    help.KeyMap
	Workflow help.KeyMap
}

var _ help.KeyMap = (*CombinedKeyMap)(nil)

func (c CombinedKeyMap) ShortHelp() []key.Binding {
	combinedBindings := c.Global.ShortHelp()
	if c.Stage != nil {
		combinedBindings = append(combinedBindings, c.Stage.ShortHelp()...)
	}
	if c.Workflow != nil {
		combinedBindings = append(combinedBindings, c.Workflow.ShortHelp()...)
	}
	return combinedBindings
}

func (c CombinedKeyMap) FullHelp() [][]key.Binding {
	combinedBindings := c.Global.FullHelp()
	if c.Stage != nil {
		combinedBindings = append(combinedBindings, c.Stage.FullHelp()...)
	}
	if c.Workflow != nil {
		combinedBindings = append(combinedBindings, c.Workflow.FullHelp()...)
	}
	return combinedBindings
}
