# Terraform Bootstrapper AI Coding Instructions

This document provides guidance for AI agents contributing to the `terraform-bootstrapper` codebase.

## Project Overview

This project is a Terminal User Interface (TUI) application built with Go and the [Bubble Tea](https://github.com/charmbracelet/bubbletea) framework. Its purpose is to guide users through bootstrapping a HashiCorp Cloud Platform (HCP) Terraform workspace.

The application follows The Elm Architecture (Model-View-Update), which is the core pattern of the Bubble Tea framework.

## Architecture

The application's logic is organized into a series of `stages`. The user progresses through these stages to complete a workflow.

- **`internal/program`**: This contains the main Bubble Tea program (`program.go`), which manages the overall application state and transitions between stages.
- **`internal/stages`**: Each subdirectory represents a distinct stage of the UI, such as `selectworkflow` or `executeworkflow`.
  - Each stage has its own `Update` function to handle messages and a `View` function to render the UI for that stage.
  - Stages are self-contained and communicate with the main program via `tea.Msg` objects.
- **`internal/messages`**: Defines the `tea.Msg` types used for communication between different parts of the application. When adding new functionality that requires state change notifications, define a new message struct here.
- **`internal/workflows`**: Contains the core business logic that is orchestrated by the TUI stages. For example, `createcontrolworkspace/createcontrolworkspace.go` contains the logic for creating a control workspace. These workflows are typically triggered from within a stage.
- **`internal/keymap`**: Defines the keybindings for different views. Each stage often has its own `keymap.go` file.
- **`internal/styles`**: Contains `lipgloss` styles for styling the TUI components.

### Data Flow Example: Creating a Project

1.  The user interacts with a view in a `stage`.
2.  The stage's `Update` function handles the user input and may call a function from a `workflow`.
3.  The workflow function executes the business logic (e.g., making an API call).
4.  The workflow function returns a `tea.Cmd` that, when run, will produce a `tea.Msg` (e.g., `ProjectCreateResultMsg` from `internal/commands/projectcreate.go`).
5.  The main program's `Update` function receives this message and updates the application's state, potentially transitioning to a new stage.

## Developer Workflow

### Build & Run

The project is a standard Go application.

- **To run the application:**
  ```sh
  go run ./cmd/terraform-bootstrapper
  ```
- **To build the binary:**
  ```sh
  go build -o terraform-bootstrapper ./cmd/terraform-bootstrapper
  ```

### Dependencies

Dependencies are managed using Go Modules. To add a new dependency, use `go get`.

## Coding Conventions

- **Styling**: Use the `lipgloss` library for all TUI styling. Define new styles in `internal/styles/styles.go`.
- **State Management**: Follow the Bubble Tea pattern. State should be managed in a model, and changes should be driven by messages. Avoid global state.
- **Error Handling**: Errors should be passed as `error` types within messages. For example, `ProjectCreateResultMsg` contains an `Err` field. The receiving `Update` function is responsible for handling the error and displaying it to the user.
