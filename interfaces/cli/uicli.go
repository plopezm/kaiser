package cli

import (
	"fmt"

	prompt "github.com/c-bata/go-prompt"
)

const (
	JOBS_COMMANDS = "jobs"
)

func executor(in string) {
	fmt.Println("Your input: " + in)
	switch in {
	case JOBS_COMMANDS:

	}
}

func completer(in prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "users", Description: "Store the username and age"},
		{Text: "articles", Description: "Store the article text posted by user"},
		{Text: "comments", Description: "Store the text commented to articles"},
		{Text: "groups", Description: "Combine users with specific rules"},
	}
	return prompt.FilterHasPrefix(s, in.GetWordBeforeCursor(), true)
}

func StartUICli() {
	p := prompt.New(executor, completer, prompt.OptionPrefix("> "), prompt.OptionTitle("Kaiser"))
	p.Run()
}
