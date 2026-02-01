package footballdataapi

import (
	"encoding/json"
	"io"
)

type RequestArea struct {
	client *Client
}

type Response struct {
	Areas []*Area
}

func (r *RequestArea) Do() ([]*Area, error) {
	res, err := r.client.Do("https://api.football-data.org/v4/areas")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	resp := Response{}
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}

	return resp.Areas, nil
}
