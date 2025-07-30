package createcontrolworkspace

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/leefowlercu/terraform-bootstrapper/internal/commands"
	"github.com/leefowlercu/terraform-bootstrapper/internal/styles"
	"github.com/leefowlercu/terraform-bootstrapper/internal/workflows"
)

type model struct {
	identifier       string
	title            string
	shortDescription string
	longDescription  string
	keys             createControlWorkspaceKeyMap

	focused            int
	inputs             []textinput.Model
	spinner            spinner.Model
	step               step
	projectStatus      status
	workspaceStatus    status
	githubAppStatus    status
	workspaceVCSStatus status
	variableStatus     status
	validationMessage  string
	submitButtonText   string

	submittedOrgName           string
	submittedProjectName       string
	submittedWorkspaceName     string
	submittedRepositoryPath    string
	submittedUserToken         string
	submittedOrganizationToken string

	actionErr   error
	projectID   string
	workspaceID string
	gitHubAppID string
	variableID  string
}

type step int

const (
	stepInput step = iota
	stepProcessing
	stepComplete
)

type status int

const (
	statusPending status = iota
	statusRunning
	statusSuccess
	statusFailure
)

// Compile-time validation of implementation of the Workflow interface
var _ workflows.Workflow = (*model)(nil)

// Compile-time validation of implementation of the list.DefaultItem interface
var _ list.DefaultItem = (*model)(nil)

func New() *model {
	inputs := make([]textinput.Model, 6)

	for i := range inputs {
		inputs[i] = textinput.New()
		inputs[i].Width = 120 // Initialize with a default width
	}

	// Create TextInput Placeholder Values
	inputs[0].Placeholder = "Organization Name..."
	inputs[1].Placeholder = "Project Name..."
	inputs[2].Placeholder = "Workspace Name..."
	inputs[3].Placeholder = "Repository Path..."
	inputs[4].Placeholder = "User Token..."
	inputs[5].Placeholder = "Organization Token..."

	// Set EchoMode for sensitive inputs
	inputs[4].EchoMode = textinput.EchoPassword
	inputs[4].EchoCharacter = '•'
	inputs[5].EchoMode = textinput.EchoPassword
	inputs[5].EchoCharacter = '•'

	// Set Initial focus on the first input
	inputs[0].Focus()

	// Initialize the Spinner
	s := spinner.New()
	s.Spinner = spinner.Line
	s.Style = styles.Spinner

	// Return initial Model state
	return &model{
		identifier:       "create-control-workspace",
		title:            "Create Control Workspace",
		shortDescription: "Creates a Control Workspace and associated Project",
		longDescription:  "This workflow will create an HCP Terraform or Terraform Enterprise Control Workspace and an associated Project. It will also configure the Workspace with a VCS Connection (GitHub.com) and set up the necessary variables for Provider Authentication.",
		keys:             defaultCreateControlWorkspaceKeyMap,

		focused:            0,
		inputs:             inputs,
		spinner:            s,
		step:               stepInput,
		projectStatus:      statusPending,
		workspaceStatus:    statusPending,
		githubAppStatus:    statusPending,
		workspaceVCSStatus: statusPending,
		variableStatus:     statusPending,
		validationMessage:  "",
		submitButtonText:   "[ Submit ]",
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (workflows.Workflow, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.step == stepInput {
			switch {
			case key.Matches(msg, m.keys.Up):
				if m.focused > 0 {
					m.focused--
				}
			case key.Matches(msg, m.keys.Down):
				if m.focused < len(m.inputs) {
					m.focused++
				}
			case key.Matches(msg, m.keys.CycleF):
				m.focused = (m.focused + 1) % len(m.inputs)
			case key.Matches(msg, m.keys.CycleB):
				m.focused = (m.focused - 1 + len(m.inputs)) % len(m.inputs)
			case key.Matches(msg, m.keys.Submit):
				// Submit Button is focused
				if m.focused == len(m.inputs) {
					// Perform validation on all inputs
					for i := range m.inputs {
						if m.inputs[i].Value() == "" {
							m.focused = i
							m.inputs[i].Focus()
							m.validationMessage = fmt.Sprintf("Please enter a value for the required field: %s", strings.TrimSuffix(m.inputs[i].Placeholder, "..."))
							return m, nil
						}
					}

					// If all inputs are valid, gather submitted values into the model
					// and proceed to processing step
					m.validationMessage = ""
					m.submittedOrgName = m.inputs[0].Value()
					m.submittedProjectName = m.inputs[1].Value()
					m.submittedWorkspaceName = m.inputs[2].Value()
					m.submittedRepositoryPath = m.inputs[3].Value()
					m.submittedUserToken = m.inputs[4].Value()
					m.submittedOrganizationToken = m.inputs[5].Value()
					m.step = stepProcessing

					// Set the Create Project Status to Running and send the CreateProject Command
					m.projectStatus = statusRunning
					return m, tea.Sequence(
						m.spinner.Tick,
						commands.CreateProject(m.submittedUserToken, m.submittedOrgName, m.submittedProjectName),
					)
				} else {
					m.focused = (m.focused + 1) % (len(m.inputs) + 1)
				}
			}
		}

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd

	case commands.CreateProjectResultMsg:
		// An error occurred while creating the Project
		// set the Project Status to Failure and update the Model with the error
		if msg.Err != nil {
			m.projectStatus = statusFailure
			m.actionErr = msg.Err
			return m, nil
		}

		// Project created successfully, set the Create Project Status to Success
		m.projectStatus = statusSuccess
		m.projectID = msg.ProjectID

		// Set the Create Workspace Status to Running and send the CreateWorkspace Command
		m.workspaceStatus = statusRunning
		return m, tea.Sequence(
			m.spinner.Tick,
			commands.CreateWorkspace(m.submittedUserToken, m.submittedOrgName, m.projectID, m.submittedWorkspaceName),
		)

	case commands.CreateWorkspaceResultMsg:
		// An error occurred while creating the Workspace
		// set the Workspace Status to Failure and update the Model with the error
		if msg.Err != nil {
			m.workspaceStatus = statusFailure
			m.actionErr = msg.Err
			return m, nil
		}

		// Workspace created successfully, set the Create Workspace Status to Success
		m.workspaceStatus = statusSuccess
		m.workspaceID = msg.WorkspaceID

		// Set the Create GitHub App Status to Running and send the GetGitHubAppID Command
		m.githubAppStatus = statusRunning
		return m, tea.Sequence(
			m.spinner.Tick,
			commands.GetGitHubAppID(m.submittedUserToken, m.submittedRepositoryPath),
		)

	case commands.GetGitHubAppIDMsg:
		// An error occurred while getting the GitHub App ID
		// set the GitHub App Status to Failure and update the Model with the error
		if msg.Err != nil {
			m.githubAppStatus = statusFailure
			m.actionErr = msg.Err
			return m, nil
		}

		// GitHub App ID retrieved successfully, set the Create GitHub App Status to Success
		m.githubAppStatus = statusSuccess
		m.gitHubAppID = msg.GitHubAppID

		// Set the Update Workspace VCS Status to Running and send the UpdateWorkspaceVCS Command
		m.workspaceVCSStatus = statusRunning
		return m, tea.Sequence(
			m.spinner.Tick,
			commands.UpdateWorkspaceVCS(m.submittedUserToken, m.workspaceID, m.submittedRepositoryPath, m.gitHubAppID),
		)

	case commands.UpdateWorkspaceVCSMsg:
		// An error occurred while updating the Workspace VCS Connection
		// set the Workspace VCS Status to Failure and update the Model with the error
		if msg.Err != nil {
			m.workspaceVCSStatus = statusFailure
			m.actionErr = msg.Err
			return m, nil
		}

		// Workspace VCS Connection updated successfully, set the Update Workspace VCS Status to Success
		m.workspaceVCSStatus = statusSuccess

		// Set the Add Workspace Variable Status to Running and send the AddWorkspaceVariable Command
		m.variableStatus = statusRunning
		return m, tea.Sequence(
			m.spinner.Tick,
			commands.AddWorkspaceVariable(m.submittedUserToken, m.workspaceID, m.submittedOrganizationToken),
		)

	case commands.AddWorkspaceVariableResultMsg:
		// An error occurred while adding the Workspace Variable
		// set the Variable Status to Failure and update the Model with the error
		if msg.Err != nil {
			m.variableStatus = statusFailure
			m.actionErr = msg.Err
			return m, nil
		}

		// Workspace Variable added successfully, set the Add Workspace Variable Status to Success
		m.variableStatus = statusSuccess
		m.variableID = msg.VariableID

		// Set the Model step to complete
		m.step = stepComplete
		return m, nil
	}

	cmds := make([]tea.Cmd, len(m.inputs))
	// Set focus on the currently focused index, if focus index represents a textinput
	for i := range m.inputs {
		if i == m.focused {
			m.inputs[i].Focus()
		} else {
			m.inputs[i].Blur()
		}
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	var sb strings.Builder

	switch m.step {
	case stepInput:
		// Style and render the input fields appropriately
		for i := range m.inputs {
			if i == m.focused {
				m.inputs[i].TextStyle = styles.Focused
				m.inputs[i].PromptStyle = styles.Focused
				m.inputs[i].PlaceholderStyle = styles.Blurred.Bold(true)
			} else {
				if m.inputs[i].Value() != "" {
					m.inputs[i].TextStyle = styles.Entered
					m.inputs[i].PromptStyle = styles.Entered
					m.inputs[i].PlaceholderStyle = styles.Entered
				} else {
					m.inputs[i].TextStyle = styles.Blurred
					m.inputs[i].PromptStyle = styles.Blurred
					m.inputs[i].PlaceholderStyle = styles.Blurred
				}
			}
			sb.WriteString(m.inputs[i].View() + "\n")
		}

		// Style and render the Submit Button
		if m.focused == len(m.inputs) {
			sb.WriteString("\n" + styles.Focused.Render(m.submitButtonText) + "\n")
		} else {
			sb.WriteString("\n" + styles.Blurred.Render(m.submitButtonText) + "\n")
		}

		// Style and render the Validation Message, if any
		if m.validationMessage != "" {
			sb.WriteString("\n" + styles.Warning.Render(m.validationMessage))
		}
	case stepProcessing:
		sb.WriteString(styles.WorkflowDescription.Render("Creating Terraform Resources...") + "\n\n")

		// Render the Create Project Status components
		switch m.projectStatus {
		case statusRunning:
			sb.WriteString(fmt.Sprintf("%s Creating project '%s' in organization '%s'...\n",
				m.spinner.View(), m.submittedProjectName, m.submittedOrgName,
			))
		case statusSuccess:
			sb.WriteString(fmt.Sprintf("%s Project created successfully!\n",
				styles.Success.Render("✓"),
			))
		case statusFailure:
			sb.WriteString(fmt.Sprintf("%s Failed to create project: %v\n",
				styles.Failure.Render("✗"), m.actionErr,
			))
		}

		// Render the Create Workspace components if Project Creation succeeds
		if m.projectStatus == statusSuccess {
			switch m.workspaceStatus {
			case statusRunning:
				sb.WriteString(fmt.Sprintf("%s Creating workspace '%s' in project '%s'...\n",
					m.spinner.View(), m.submittedWorkspaceName, m.submittedProjectName,
				))
			case statusSuccess:
				sb.WriteString(fmt.Sprintf("%s Workspace created successfully!\n",
					styles.Success.Render("✓"),
				))
			case statusFailure:
				sb.WriteString(fmt.Sprintf("%s Failed to create workspace: %v\n",
					styles.Failure.Render("✗"), m.actionErr,
				))
			}
		}

		// Render the Get GitHub App ID components if Workspace Creation succeeds
		if m.workspaceStatus == statusSuccess {
			switch m.githubAppStatus {
			case statusRunning:
				sb.WriteString(fmt.Sprintf("%s Retrieving GitHub App ID for repository '%s'...\n",
					m.spinner.View(), m.submittedRepositoryPath,
				))
			case statusSuccess:
				sb.WriteString(fmt.Sprintf("%s GitHub App ID retrieved successfully!\n",
					styles.Success.Render("✓"),
				))
			case statusFailure:
				sb.WriteString(fmt.Sprintf("%s Failed to retrieve GitHub App ID: %v\n",
					styles.Failure.Render("✗"), m.actionErr,
				))
			}
		}

		// Render the Update Workspace VCS components if GitHub App ID retrieval succeeds
		if m.githubAppStatus == statusSuccess {
			switch m.workspaceVCSStatus {
			case statusRunning:
				sb.WriteString(fmt.Sprintf("%s Updating workspace '%s' with VCS connection to Repository '%s'...\n",
					m.spinner.View(), m.submittedWorkspaceName, m.submittedRepositoryPath,
				))
			case statusSuccess:
				sb.WriteString(fmt.Sprintf("%s Workspace VCS connection updated successfully!\n",
					styles.Success.Render("✓"),
				))
			case statusFailure:
				sb.WriteString(fmt.Sprintf("%s Failed to update workspace VCS connection: %v\n",
					styles.Failure.Render("✗"), m.actionErr,
				))
			}
		}

		// Render the Add Workspace Variable components if Workspace VCS update succeeds
		if m.workspaceVCSStatus == statusSuccess {
			switch m.variableStatus {
			case statusRunning:
				sb.WriteString(fmt.Sprintf("%s Adding workspace variable for organization token...\n",
					m.spinner.View(),
				))
			case statusSuccess:
				sb.WriteString(fmt.Sprintf("%s Workspace variable added successfully!\n",
					styles.Success.Render("✓"),
				))
			case statusFailure:
				sb.WriteString(fmt.Sprintf("%s Failed to add workspace variable: %v\n",
					styles.Failure.Render("✗"), m.actionErr,
				))
			}
		}

	case stepComplete:
		sb.WriteString("\n" + styles.Success.Render(fmt.Sprintf("Control Workspace '%s' successfully bootstrapped!\n", m.submittedWorkspaceName)))
	}

	return sb.String()
}

func (m model) KeyMap() help.KeyMap {
	return m.keys
}

func (m model) Identifier() string {
	return m.identifier
}

func (m model) Title() string {
	return m.title
}

func (m model) Description() string {
	return m.shortDescription
}

func (m model) LongDescription() string {
	return m.longDescription
}

func (m model) FilterValue() string {
	return m.title
}
