package main

import (
	"command/pkg/command"
)

func main() {
	light := &command.Light{}
	remote := &command.RemoteControl{}

	onCommand := command.NewLightOnCommand(light)
	offCommand := command.NewLightOffCommand(light)

	remote.SetCommand(onCommand)
	remote.PressButton()

	remote.SetCommand(offCommand)
	remote.PressButton()
}
