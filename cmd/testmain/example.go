package main

import (
	"fmt"
	"log"
	"net/http"

	footballdataapi "github.com/fmo/football-data-api"
)

func main() {
	client := footballdataapi.NewMatches(&footballdataapi.Client{HTTPClient: &http.Client{}})
	resp, err := client.Do(footballdataapi.PL, 2025, 26)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)
}
