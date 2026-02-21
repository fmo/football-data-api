package main

import (
	"fmt"
	"log"

	footballdataapi "github.com/fmo/football-data-api"
)

func main() {
	client := footballdataapi.NewClient()
	resp, err := client.TeamMatches.Fetch(759)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)
}
