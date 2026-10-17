package main

import "fmt"

// Команда
type Command interface {
	Execute()
	Undo()
}

// Получатель
type Light struct{ isOn bool }

func (l *Light) On()  { l.isOn = true; fmt.Println("Свет включен") }
func (l *Light) Off() { l.isOn = false; fmt.Println("Свет выключен") }

// Конкретные команды
type LightOnCommand struct{ light *Light }

func (c *LightOnCommand) Execute() { c.light.On() }
func (c *LightOnCommand) Undo()    { c.light.Off() }

type LightOffCommand struct{ light *Light }

func (c *LightOffCommand) Execute() { c.light.Off() }
func (c *LightOffCommand) Undo()    { c.light.On() }

// Инициатор (Пульт)
type RemoteControl struct {
	history []Command
}

func (r *RemoteControl) Execute(cmd Command) {
	cmd.Execute()
	r.history = append(r.history, cmd)
}

func (r *RemoteControl) Undo() {
	if len(r.history) == 0 {
		return
	}
	lastCmd := r.history[len(r.history)-1]
	r.history = r.history[:len(r.history)-1]
	lastCmd.Undo()
}

func main() {
	light := &Light{}
	remote := &RemoteControl{}

	remote.Execute(&LightOnCommand{light})
	remote.Execute(&LightOffCommand{light})
	remote.Undo() // Включит свет обратно
}

//https://coddy.tech/learn/ru/go/object_oriented_programming/command_pattern
