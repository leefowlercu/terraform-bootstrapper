package messages

import "github.com/leefowlercu/terraform-bootstrapper/internal/stages"

type AvailableSizeMsg struct {
	Width  int
	Height int
}

type ChangeStageMsg struct {
	Stage stages.Stage
}
