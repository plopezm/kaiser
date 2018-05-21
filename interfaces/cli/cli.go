package cli

import (
	ishell "gopkg.in/abiosoft/ishell.v2"
)

// StartShell Starts a cli interface with the commands
func StartShell(commands ...*ishell.Cmd) {
	// create new shell.
	// by default, new shell includes 'exit', 'help' and 'clear' commands.
	shell := ishell.New()

	// display welcome info.
	shell.Println("Sample Interactive Shell")

	for _, command := range commands {
		shell.AddCmd(command)
	}

	// register a function for "greet" command.
	shell.AddCmd(&ishell.Cmd{
		Name: "version",
		Help: "Shows program version",
		Func: func(c *ishell.Context) {
			c.Println("Kaiser - v1.0.0")
		},
	})

	// run shell
	shell.Run()
}
