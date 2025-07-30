package styles

import "github.com/charmbracelet/lipgloss"

var (
	Program = lipgloss.NewStyle().Padding(1, 1, 0, 1)

	Focused = lipgloss.NewStyle().Foreground(lipgloss.Color("202")).Bold(true)
	Blurred = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	Entered = lipgloss.NewStyle().Foreground(lipgloss.Color("255")).Bold(true)
	Warning = lipgloss.NewStyle().Foreground(lipgloss.Color("196"))
	Success = lipgloss.NewStyle().Foreground(lipgloss.Color("46"))
	Failure = lipgloss.NewStyle().Foreground(lipgloss.Color("196"))
	Spinner = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	WorkflowTitle       = lipgloss.NewStyle().Background(lipgloss.Color("62")).MarginLeft(2).Padding(0, 1)
	WorkflowDescription = lipgloss.NewStyle().Foreground(lipgloss.Color("250")).MarginLeft(1).Padding(0, 1)
)
