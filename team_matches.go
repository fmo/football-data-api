package footballdataapi

import (
	"encoding/json"
	"fmt"
	"io"
)

type TeamMatches struct {
	client *Client
}

func newTeamMatches(c *Client) *TeamMatches {
	return &TeamMatches{c}
}

type TeamMatchesResponse struct {
	Matches []Match `json:"matches"`
}

func (tm *TeamMatches) Do(teamID int) (*TeamMatchesResponse, error) {
	resp, err := tm.client.Do(fmt.Sprintf("%s/teams/%d/matches", url, teamID))
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	r := TeamMatchesResponse{}

	if err := json.Unmarshal(body, &r); err != nil {
		return nil, err
	}

	return &r, nil
}
