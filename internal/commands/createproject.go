package commands

import (
	"context"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/hashicorp/go-tfe"
)

type CreateProjectResultMsg struct {
	ProjectID string
	Err       error
}

func CreateProject(token, organizationName, projectName string) tea.Cmd {
	return func() tea.Msg {
		config := &tfe.Config{
			Token: token,
		}

		client, err := tfe.NewClient(config)
		if err != nil {
			return CreateProjectResultMsg{Err: err}
		}

		p, err := client.Projects.Create(context.Background(), organizationName, tfe.ProjectCreateOptions{
			Name: projectName,
		})
		if err != nil {
			return CreateProjectResultMsg{Err: err}
		}

		time.Sleep(2 * time.Second) // Introduce a delay to render fancy UI elements (spinners and such)
		return CreateProjectResultMsg{ProjectID: p.ID}
	}
}
