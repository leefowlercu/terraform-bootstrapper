package keymap

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
)

type CombinedKeyMap struct {
	Global GlobalKeyMap
	Stage  help.KeyMap
}

var _ help.KeyMap = (*CombinedKeyMap)(nil)

func (c CombinedKeyMap) ShortHelp() []key.Binding {
	return append(c.Global.ShortHelp(), c.Stage.ShortHelp()...)
}

func (c CombinedKeyMap) FullHelp() [][]key.Binding {
	return append(c.Global.FullHelp(), c.Stage.FullHelp()...)
}
