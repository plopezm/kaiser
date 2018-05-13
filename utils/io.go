package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// GetJSONObjectFromFile Fills the object with the content of a file
func GetJSONObjectFromFile(filepath string, object interface{}) error {
	raw, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatalln(err.Error())
		os.Exit(1)
	}
	err = json.Unmarshal(raw, &object)
	return err
}
