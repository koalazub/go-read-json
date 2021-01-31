package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)
import "testing"

func TestOpenFile(t *testing.T) {
	openedFile, err := os.Open("../tests.json")
	if openedFile != nil {
		t.Errorf("Expected read file, but got %v", err)
	}
}

func TestReadFile(t *testing.T) {
	var content map[string]interface{}
	openData, err := os.Open("./tests.json")
	readData, err2 := ioutil.ReadAll(openData)
	if err2 != nil {
		t.Errorf("could not read opened file. Check that it is aligned with path and file type")
	}
	if err != nil {
		t.Errorf("Could not read file. Check path for potential failures")
	}
	unmarshalledData := json.Unmarshal(readData, &content)
	if unmarshalledData != nil {
		t.Errorf("Could not unmarshal JSON data")
	}
}

func TestReturnLimitedTweetCount(t *testing.T) {

}

func TestReturnRating(t *testing.T) {

}
