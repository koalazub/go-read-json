package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type MapData struct {
	results []map[string]interface{}
}

type Review struct {
	Title  string `json:"title"`
	Review string `json:"review"`
	Score  int    `json:"score"`
}

func main() {
	readFile, err := os.Open("./reviews.json")
	if err != nil {
		log.Printf("Failed to open file!! %+v", err)
	}
	returnUnmarshalledData(convertItemsInFileToData(readFile))

}

func convertItemsInFileToData(file *os.File) []byte {
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("Unable to read from file %+v - check the directory. Or maybe you fucked up the read somehow", err)
	}
	return data
}

func returnUnmarshalledData(data []byte) {
	//var results MapData
	var review []Review
	unMarshalledReviews := json.Unmarshal(data, &review)

	if unMarshalledReviews != nil {
		log.Fatal("Check the reads for the file. Unable to read")
	}

	starRating := "*"
	for rev := range review {

		fmt.Printf("%+v\n", review[rev].Score)

		if review[rev].Score % 25 == 0 {
			starRating += starRating
			fmt.Printf(starRating + "\n")
		}
	}

}