package test

import (
	"command/pkg/command"
	"testing"
)

func TestLightOnCommand_Execute(t *testing.T) {
	light := &command.Light{}
	command := command.NewLightOnCommand(light)

	command.Execute()

	if !light.IsOn {
		t.Error("Expected light to be on")
	}
}
