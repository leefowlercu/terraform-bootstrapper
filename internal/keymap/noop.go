package keymap

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
)

// NoopKeyMap is an empty KeyMap that implements the help.KeyMap interface.
type NoopKeyMap struct{}

var _ help.KeyMap = (*NoopKeyMap)(nil)

func (NoopKeyMap) ShortHelp() []key.Binding {
	return nil
}

func (NoopKeyMap) FullHelp() [][]key.Binding {
	return nil
}
