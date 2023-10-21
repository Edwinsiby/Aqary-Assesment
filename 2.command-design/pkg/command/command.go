package command

import "fmt"

type Command interface {
	Execute()
}

type Light struct {
	IsOn bool
}

func (l *Light) On() {
	l.IsOn = true
	fmt.Println("Light is on")
}

func (l *Light) Off() {
	l.IsOn = false
	fmt.Println("Light is off")
}

type LightOnCommand struct {
	light *Light
}

func NewLightOnCommand(light *Light) *LightOnCommand {
	return &LightOnCommand{light}
}

func (c *LightOnCommand) Execute() {
	c.light.On()
}

type LightOffCommand struct {
	light *Light
}

func NewLightOffCommand(light *Light) *LightOffCommand {
	return &LightOffCommand{light}
}

func (c *LightOffCommand) Execute() {
	c.light.Off()
}

type RemoteControl struct {
	command Command
}

func (r *RemoteControl) SetCommand(command Command) {
	r.command = command
}

func (r *RemoteControl) PressButton() {
	r.command.Execute()
}
