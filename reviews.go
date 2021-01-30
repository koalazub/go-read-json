package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func getReview(reviews []byte) {
	var reviewDescription []Description

	unMarshalledReviews := json.Unmarshal(reviews, &reviewDescription)
	if unMarshalledReviews != nil {
		log.Fatal("Check the reads for the file. Unable to read")
	}

	for rev := range reviewDescription {
		reviewDescription := reviewDescription[rev]
		returnLimitedTweetCount(reviewDescription.Title, reviewDescription.Review, 140)
		returnRating(reviewDescription.Score)
	}
}

func returnLimitedTweetCount(title string, review string, maxTweetLength int) {
	fmt.Printf("Title: %+v\n", title)

	if len(review) > maxTweetLength {
		review = review[:maxTweetLength] + "..."
	}
	fmt.Printf("Review: %+v\n", review)
}
func returnRating(score int) {
	if score <= 25 {
		fmt.Printf("Rating: %+v\n\n", "*")
	}
	if score <= 50 && score > 25 {
		fmt.Printf("score: %+v\n\n", "**")
	}
	if score <= 99 && score > 50 {
		fmt.Printf("score: %+v\n\n", "****")
	}
	if score == 100 {
		fmt.Printf("Rating: %+v\n\n", "*****")
	}
}
