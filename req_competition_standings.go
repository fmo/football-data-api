// Package footballdataapi fetching football data
package footballdataapi

import (
	"encoding/json"
	"fmt"
	"io"
)

type ReqCompStandings struct {
	client *Client
}

func NewReqCompStandings(client *Client) *ReqCompStandings {
	return &ReqCompStandings{client}
}

type RespCompStandings struct {
	Competition Competition `json:"competition"`
	Season      Season      `json:"season"`
	Standings   []Standings `json:"standings"`
}

func (r *ReqCompStandings) Do(competitionCode CompetitionCode, season int) (*RespCompStandings, error) {
	if _, ok := competitions[competitionCode]; !ok {
		return nil, ErrNotSupportedCompetition
	}

	res, err := r.client.Do(fmt.Sprintf("%s/competitions/%s/standings?season=%d", url, competitionCode, season))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	resp := RespCompStandings{}
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
