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

func getMovies(movies []byte) {
	var movieDescription []Movies

	unMarshalledReviews := json.Unmarshal(movies, &movieDescription)
	if unMarshalledReviews != nil {
		log.Fatal("Check the reads for the file. Unable to read")
	}

	for movie := range movieDescription {
		movieField := movieDescription[movie]
		fmt.Printf("Movie title(%+v): %+v", movieField.Year, movieField.Title)
	}
}
