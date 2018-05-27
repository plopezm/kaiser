package cli

import ishell "gopkg.in/abiosoft/ishell.v2"

func ExportCommands() []*ishell.Cmd {
	var commands = make([]*ishell.Cmd, 1)
	commands[0] = &ishell.Cmd{
		Name: "joblist",
		Help: "Show current registered Jobs",
		Func: func(c *ishell.Context) {
			c.Println("=== Job List ===")
		},
	}
	return commands
}
