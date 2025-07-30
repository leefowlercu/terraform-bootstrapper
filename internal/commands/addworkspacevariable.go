package commands

import (
	"context"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/hashicorp/go-tfe"
)

type AddWorkspaceVariableResultMsg struct {
	VariableID string
	Err        error
}

func AddWorkspaceVariable(token, workspaceID, organizationToken string) tea.Cmd {
	return func() tea.Msg {
		config := &tfe.Config{
			Token: token,
		}

		client, err := tfe.NewClient(config)
		if err != nil {
			return AddWorkspaceVariableResultMsg{Err: err}
		}

		v, err := client.Variables.Create(context.Background(), workspaceID, tfe.VariableCreateOptions{
			Key:       tfe.String("TFE_TOKEN"),
			Value:     tfe.String(organizationToken),
			Category:  tfe.Category(tfe.CategoryEnv),
			Sensitive: tfe.Bool(true),
		})
		if err != nil {
			return AddWorkspaceVariableResultMsg{Err: err}
		}

		time.Sleep(2 * time.Second) // Introduce a delay to render fancy UI elements (spinners and such)
		return AddWorkspaceVariableResultMsg{VariableID: v.ID}
	}
}
