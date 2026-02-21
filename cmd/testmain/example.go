package main

import (
	"fmt"
	"log"

	footballdataapi "github.com/fmo/football-data-api"
)

func main() {
	client := footballdataapi.NewClient()
	resp, err := client.Teams.Fetch("PL")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)
}
