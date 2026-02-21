package footballdataapi

import (
	"encoding/json"
	"fmt"
	"io"
)

type Teams struct {
	client *Client
}

func newTeams(client *Client) *Teams {
	return &Teams{client}
}

type TeamsResponse struct {
	Teams []Team
}

func (t *Teams) Fetch(competitionCode CompetitionCode) (*TeamsResponse, error) {
	resp, err := t.client.Do(fmt.Sprintf("%s/competitions/%s/teams", url, string(competitionCode)))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	teamsResponse := TeamsResponse{}
	if err := json.Unmarshal(b, &teamsResponse); err != nil {
		return nil, err
	}

	return &teamsResponse, nil
}
