package commands

import (
	"context"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/hashicorp/go-tfe"
)

type UpdateWorkspaceVCSMsg struct {
	Err error
}

func UpdateWorkspaceVCS(token, workspaceID, repositoryPath, gitHubAppID string) tea.Cmd {
	return func() tea.Msg {
		config := &tfe.Config{
			Token: token,
		}

		client, err := tfe.NewClient(config)
		if err != nil {
			return UpdateWorkspaceVCSMsg{Err: err}
		}

		_, err = client.Workspaces.UpdateByID(context.Background(), workspaceID, tfe.WorkspaceUpdateOptions{
			VCSRepo: &tfe.VCSRepoOptions{
				Identifier:        tfe.String(repositoryPath),
				GHAInstallationID: tfe.String(gitHubAppID),
			},
		})
		if err != nil {
			return UpdateWorkspaceVCSMsg{Err: err}
		}

		time.Sleep(2 * time.Second) // Introduce a delay to render fancy UI elements (spinners and such)
		return UpdateWorkspaceVCSMsg{Err: nil}
	}
}
