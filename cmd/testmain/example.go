package main

import (
	"fmt"
	"log"
	"net/http"

	footballdataapi "github.com/fmo/football-data-api"
)

func main() {
	client := footballdataapi.NewMatches(&footballdataapi.Client{HTTPClient: &http.Client{}})
	resp, err := client.Do(footballdataapi.PL, footballdataapi.Finished, 2025, 24)

	if err != nil {
		log.Fatal(err)
	}

	for _, m := range resp.Matches {
		fmt.Println(m.HomeTeam.Name, m.AwayTeam.Name, m.Score.FullTime)
	}
}
