package commands

import (
	"context"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/hashicorp/go-tfe"
)

type CreateWorkspaceResultMsg struct {
	WorkspaceID string
	Err         error
}

func CreateWorkspace(token, organizationName, projectID, workspaceName string) tea.Cmd {
	return func() tea.Msg {
		config := &tfe.Config{
			Token: token,
		}

		client, err := tfe.NewClient(config)
		if err != nil {
			return CreateWorkspaceResultMsg{Err: err}
		}

		w, err := client.Workspaces.Create(context.Background(), organizationName, tfe.WorkspaceCreateOptions{
			Name:       tfe.String(workspaceName),
			SourceName: tfe.String("terraform-bootstrapper"),
			Project:    &tfe.Project{ID: projectID},
		})
		if err != nil {
			return CreateWorkspaceResultMsg{Err: err}
		}

		time.Sleep(2 * time.Second) // Introduce a delay to render fancy UI elements (spinners and such)
		return CreateWorkspaceResultMsg{WorkspaceID: w.ID}
	}
}
