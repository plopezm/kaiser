package utils

import (
	"crypto/sha512"
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

// ReadInverseFileContent reads the content of a file a returns a pointer to it
func ReadInverseFileContent(filepath string, maxLines int) (*string, error) {
	file, _ := os.OpenFile(filepath, os.O_RDONLY, 0666)
	defer file.Close()
	fstat, _ := file.Stat()

	result := ""
	newline := ""
	linesFound := 0
	buf := make([]byte, 1)
	for i := fstat.Size() - 1; i >= 0; i-- {
		file.ReadAt(buf, i)
		if buf[0] == '\n' {
			if len(newline) == 0 {
				continue
			}
			linesFound++
			result = result + newline + "\n"
			if linesFound == maxLines {
				return &result, nil
			}
			newline = ""
			continue
		}
		newline = string(buf) + newline
	}
	return &result, nil
}
