package commands

import (
	"context"
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/hashicorp/go-tfe"
)

type GetGitHubAppIDMsg struct {
	GitHubAppID string
	Err         error
}

func GetGitHubAppID(token, repositoryPath string) tea.Cmd {
	return func() tea.Msg {
		if idx := strings.Index(repositoryPath, "/"); idx == -1 {
			return GetGitHubAppIDMsg{Err: fmt.Errorf("invalid repository path: %s, please use the format ':org/:repo'", repositoryPath)}
		}

		vcsOrg := repositoryPath[:strings.Index(repositoryPath, "/")]

		config := &tfe.Config{
			Token: token,
		}

		client, err := tfe.NewClient(config)
		if err != nil {
			return GetGitHubAppIDMsg{Err: err}
		}

		ghail, err := client.GHAInstallations.List(context.Background(), &tfe.GHAInstallationListOptions{})
		if err != nil {
			return GetGitHubAppIDMsg{Err: err}
		}

		for _, ghai := range ghail.Items {
			if *ghai.Name == vcsOrg {
				time.Sleep(2 * time.Second) // Introduce a delay to render fancy UI elements (spinners and such)
				return GetGitHubAppIDMsg{GitHubAppID: *ghai.ID}
			}
		}

		// No matching GitHub App Installation found
		return GetGitHubAppIDMsg{Err: fmt.Errorf("no GitHub App Installation found for organization: %s", vcsOrg)}
	}
}
