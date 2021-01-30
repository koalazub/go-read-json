package main

import (
	"encoding/json"
	"os"
)
import "testing"

type DummyData struct {
	data string `json:"[
	"Star Wars (1977): Great, this film was ★★★★,
    "Star Wars The Force Awake (2015): A long time ago in a galaxy far far away someone made the best sci-fi film of all time. Then some chap ★★½",
    "Metropolis: Another movie about a robot. Very strong futuristic look. But also very very old. Hard to understand what was happening becaus ★",
    "Dr. Strangelove or How I Learned to Stop Worrying and Love the Bomb (1964): Hilarious Kubrick satire ★★★★★",
    "Plan 9 from outer space (1959): So bad it's bad ½"
]"`
}

func TestReadFile(t *testing.T) {
	openedFile, err := os.Open("../reviews.json")
	if openedFile != nil {
		t.Errorf("Expected read file, but got %v", err)
	}
}

func TestGetReview(t *testing.T) {
	var dummyData DummyData
	var byteData []byte
	unmarshalledData := json.Unmarshal(byteData, &dummyData)
	if unmarshalledData != nil {
		t.Errorf("Could not unmarshal JSON data")
	}

}

func TestReturnLimitedTweetCount(t *testing.T) {

}

func TestReturnRating(t *testing.T) {

}
