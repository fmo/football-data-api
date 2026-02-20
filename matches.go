package footballdataapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"time"
)

type Status string

const (
	Finished  Status = "FINISHED"
	Scheduled Status = "SCHEDULED"
	Live      Status = "LIVE"
	InPlay    Status = "IN_PLAY"
)

var ErrNotSupportedCompetition = errors.New("competition is not supported")
var ErrNotSupportedStatus = errors.New("not supported status")

type Match struct {
	HomeTeam Team      `json:"homeTeam"`
	AwayTeam Team      `json:"awayTeam"`
	Status   string    `json:"status"`
	UtcDate  time.Time `json:"utcDate"`
	Score    struct {
		FullTime struct {
			Home int `json:"home"`
			Away int `json:"away"`
		} `json:"fullTime"`
	} `json:"score"`
}

type Matches struct {
	client *Client
}

func newMatches(client *Client) *Matches {
	return &Matches{client}
}

type MatchesResponse struct {
	Competition Competition
	Matches     []Match `json:"matches"`
}

func (m *Matches) Do(competitionCode CompetitionCode, season, matchday int) (*MatchesResponse, error) {
	if _, ok := competitions[competitionCode]; !ok {
		return nil, ErrNotSupportedCompetition
	}

	res, err := m.client.Do(fmt.Sprintf("%s/competitions/%s/matches?season=%d&matchday=%d", url, competitionCode, season, matchday))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	resp := &MatchesResponse{}
	if err := json.Unmarshal(body, resp); err != nil {
		return nil, err
	}

	return resp, nil
}
