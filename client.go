package footballdataapi

import (
	"context"
	"net/http"
	"os"
	"time"
)

const url = "https://api.football-data.org/v4"

type Client struct {
	http        *http.Client
	Standings   *Standings
	Matches     *Matches
	Areas       *Areas
	TeamMatches *TeamMatches
	Teams       *Teams
}

func NewClient() *Client {
	client := &Client{
		http: &http.Client{Timeout: 10 * time.Second},
	}

	client.Standings = newStandings(client)
	client.Matches = newMatches(client)
	client.Areas = newAreas(client)
	client.TeamMatches = newTeamMatches(client)
	client.Teams = newTeams(client)

	return client
}

func (c *Client) Do(url string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("X-Auth-Token", os.Getenv("API_KEY"))

	return c.http.Do(req)
}
