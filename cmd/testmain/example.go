package main

import (
	"fmt"
	"log"
	"net/http"

	footballdataapi "github.com/fmo/football-data-api"
)

func main() {
	client := footballdataapi.NewReqCompStandings(&footballdataapi.Client{HTTPClient: &http.Client{}})
	resp, err := client.Do(footballdataapi.PL, 2025)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.Season)
}
