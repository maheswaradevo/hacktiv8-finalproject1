package utils

import (
	"encoding/json"
	"io/fs"
	"io/ioutil"
	"log"
)

func WriteToJSON(data interface{}, filename string) error {
	dataByte, err := json.Marshal(data)
	if err != nil {
		log.Printf("[WriteToJSON] failed to marshalling the data: %v", err)
		return err
	}
	err = ioutil.WriteFile(filename, dataByte, fs.ModeAppend)
	if err != nil {
		log.Printf("[WriteToJSON] failed to write to JSON file: %v", err)
		return err
	}
	return nil
}
