// Package footballdataapi fetching football data
package footballdataapi

import (
	"encoding/json"
	"fmt"
	"io"
)

type CompetitionCode string

const (
	PL CompetitionCode = "PL"
)

var competitions = map[CompetitionCode]struct{}{
	PL: {},
}

type Competition struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Code   string `json:"code"`
	Type   string `json:"type"`
	Emblem string `json:"emblem"`
}

type Standing struct {
	Stage string  `json:"stage"`
	Type  string  `json:"type"`
	Table []Table `json:"table"`
}

type Table struct {
	Position       int `json:"position"`
	Team           Team
	PlayedGames    int    `json:"playedGames"`
	Form           string `json:"form"`
	Won            int    `json:"won"`
	Draw           int    `json:"draw"`
	Lost           int    `json:"lost"`
	Points         int    `json:"points"`
	GoalsFor       int    `json:"goalsFor"`
	GoalsAgainst   int    `json:"goalsAgainst"`
	GoalDifference int    `json:"goalDifference"`
}

type Standings struct {
	client *Client
}

func newStandings(client *Client) *Standings {
	return &Standings{client: client}
}

type StandingsResponse struct {
	Competition Competition `json:"competition"`
	Season      Season      `json:"season"`
	Standings   []Standing  `json:"standings"`
}

func (s *Standings) Fetch(competitionCode CompetitionCode, season int) (*StandingsResponse, error) {
	if _, ok := competitions[competitionCode]; !ok {
		return nil, ErrNotSupportedCompetition
	}

	res, err := s.client.Do(fmt.Sprintf("%s/competitions/%s/standings?season=%d", url, competitionCode, season))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	resp := StandingsResponse{}
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
