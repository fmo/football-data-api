package footballdataapi

import (
	"encoding/json"
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

func (r *ReqCompStandings) Do() (*RespCompStandings, error) {
	res, err := r.client.Do("https://api.football-data.org/v4/competitions/PL/standings?season=2025")
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
