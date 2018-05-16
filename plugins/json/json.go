package main

import (
	"encoding/json"
	"fmt"
)

// KaiserExports Used by Kaiser, returns new functionality for Kaiser
func KaiserExports() (functions map[string]interface{}) {
	functions = make(map[string]interface{})
	functions["stringify"] = Stringify
	functions["parse"] = Parse
	return functions
}

func Stringify(object interface{}) string {
	bytes, err := json.Marshal(object)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(bytes)
}

func Parse(jsonString string, object interface{}) {
	err := json.Unmarshal([]byte(jsonString), &object)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
}
