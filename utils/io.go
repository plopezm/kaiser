package utils

import (
	"crypto/sha512"
	"encoding/json"
	"io/ioutil"
	"log"
)

// GetJSONObjectFromFile Fills the object with the content of a file
func GetJSONObjectFromFile(filepath string, object interface{}) error {
	raw, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	err = json.Unmarshal(raw, object)
	return err
}

// GetJSONObjectFromFileWithHash Fills the object with the content of a file
func GetJSONObjectFromFileWithHash(filepath string, object interface{}) ([]byte, error) {
	raw, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatalln(err.Error())
		return nil, err
	}
	err = json.Unmarshal(raw, object)

	hasher := sha512.New()
	hasher.Write(raw)
	return hasher.Sum(nil), err
}

// ReadFileContent reads the content of a file a returns a pointer to it
func ReadFileContent(filepath string) (*string, error) {
	raw, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatalln(err.Error())
		return nil, err
	}
	content := string(raw)
	return &content, nil
}
