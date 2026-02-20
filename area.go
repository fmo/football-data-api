package footballdataapi

import (
	"encoding/json"
	"io"
)

type Area struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	CountryCode  string `json:"countryCode"`
	Flag         string `json:"flag"`
	ParentAreaID int    `json:"parentAreaId"`
	ParentArea   string `json:"parentArea"`
}

type Areas struct {
	client *Client
}

func newAreas(client *Client) *Areas {
	return &Areas{
		client: client,
	}
}

type AreasResponse struct {
	Areas []*Area
}

func (a *Areas) Do() (*AreasResponse, error) {
	res, err := a.client.Do("https://api.football-data.org/v4/areas")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	resp := AreasResponse{}
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
