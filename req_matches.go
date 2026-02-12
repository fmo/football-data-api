package footballdataapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

var ErrNotSupportedCompetition = errors.New("competition is not supported")
var ErrNotSupportedStatus = errors.New("not supported status")

type ReqMatches struct {
	client *Client
}

func NewMatches(client *Client) *ReqMatches {
	return &ReqMatches{client}
}

type RespMatches struct {
	Competition Competition
	Matches     []Match `json:"matches"`
}

func (r *ReqMatches) Do(competitionCode CompetitionCode, status Status, season, matchday int) (*RespMatches, error) {
	if _, ok := competitions[competitionCode]; !ok {
		return nil, ErrNotSupportedCompetition
	}

	if _, ok := statuses[status]; !ok {
		return nil, ErrNotSupportedStatus
	}

	res, err := r.client.Do(fmt.Sprintf("%s/competitions/%s/matches?season=%d&matchday=%d&status=%s", url, competitionCode, season, matchday, status))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	resp := &RespMatches{}
	if err := json.Unmarshal(body, resp); err != nil {
		return nil, err
	}

	return resp, nil
}
