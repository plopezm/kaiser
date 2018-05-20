package task

import (
	"io/ioutil"
	"log"
	"os"
)

// JobArgs Represents the input arguments to the executor
type JobArgs struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}

// JobTask Represents a job task to be performed
type JobTask struct {
	Script     *string `json:"script"`
	ScriptFile *string `json:"scriptFile"`
	OnSuccess  string  `json:"onSuccess"`
	OnFailure  string  `json:"onFailure"`
}

func (task *JobTask) GetScript(folder string) string {
	if task.Script != nil {
		return *task.Script
	}
	if task.ScriptFile != nil {
		raw, err := ioutil.ReadFile(folder + *task.ScriptFile)
		if err != nil {
			log.Fatalln(err.Error())
			os.Exit(1)
		}
		return string(raw)
	}
	return "console.log('ERROR PARSING SCRIPT')"
}
