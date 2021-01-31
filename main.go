package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

type Description struct {
	Title  string `json:"title"`
	Review string `json:"review"`
	Score  int    `json:"score"`
	Year   int    `json:"year"`
}

func main() {
	getReview(openFileToData("./reviews.json"))
	getMovies(openFileToData("./movies.json"))

}

func openFileToData(filename string) []byte {
	file, err := os.Open(filename)
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("Unable to read from file %+v - check the directory", err)
	}
	return data
}
