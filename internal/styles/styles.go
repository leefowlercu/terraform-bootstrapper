package styles

import "github.com/charmbracelet/lipgloss"

var (
	AppStyle = lipgloss.NewStyle().Padding(1, 2)

	FocusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("202")).Bold(true)
	BlurredStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	EnteredStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("255")).Bold(true)
	WarningStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("196"))
	SuccessStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("46"))
	FailureStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("196"))
	SpinnerStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
)
