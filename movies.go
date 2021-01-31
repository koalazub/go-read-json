package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movies struct {
	Title string `json:"title,omitempty"`
	Year  int    `json:"year,omitempty"`
}

func getMovies(movies []byte) []Movies {
	var movieItems []Movies

	unMarshalledReviews := json.Unmarshal(movies, &movieItems)
	if unMarshalledReviews != nil {
		log.Fatal("Check the reads for the file. Unable to read")
	}

	for movie := range movieItems {
		movieField := movieItems[movie]
		fmt.Printf("Movie title(%+v): %+v", movieField.Year, movieField.Title)
	}
	return movieItems
}
