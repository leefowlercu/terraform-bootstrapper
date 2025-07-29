package styles

import "github.com/charmbracelet/lipgloss"

var (
	Focused = lipgloss.NewStyle().Foreground(lipgloss.Color("202")).Bold(true)
	Blurred = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	Entered = lipgloss.NewStyle().Foreground(lipgloss.Color("255")).Bold(true)
	Warning = lipgloss.NewStyle().Foreground(lipgloss.Color("196"))
	Success = lipgloss.NewStyle().Foreground(lipgloss.Color("46"))
	Failure = lipgloss.NewStyle().Foreground(lipgloss.Color("196"))
	Spinner = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	APadded = lipgloss.NewStyle().Padding(1)
	HPadded = lipgloss.NewStyle().Padding(0, 1)
	VPadded = lipgloss.NewStyle().Padding(1, 0)
	AMargin = lipgloss.NewStyle().Margin(1)
	HMargin = lipgloss.NewStyle().Margin(0, 1)
	VMargin = lipgloss.NewStyle().Margin(1, 0)
	Program = lipgloss.NewStyle().Padding(1, 0, 0, 1)
)
