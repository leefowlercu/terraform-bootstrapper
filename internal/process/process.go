package process

import (
	"github.com/charmbracelet/bubbles/list"
)

type process struct {
	identifier  string
	title       string
	description string
}

// Force implementation of the list.Item Interface
var _ list.Item = (*process)(nil)

// var _ list.ItemDelegate = (*process)(nil)

func New(identifier, title, description string) process {
	return process{
		identifier:  identifier,
		title:       title,
		description: description,
	}
}

func (p process) Identifier() string {
	return p.identifier
}

func (p process) Title() string {
	return p.title
}

func (p process) Description() string {
	return p.description
}

// Implement the list.Item interface's FilterValue method
func (p process) FilterValue() string {
	return p.title
}
