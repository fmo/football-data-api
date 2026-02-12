package footballdataapi

import (
	"context"
	"net/http"
	"os"
)

const url = "https://api.football-data.org/v4"

type Client struct {
	HTTPClient *http.Client
}

func NewClient(httpClient *http.Client) *Client {
	return &Client{httpClient}
}

func (c *Client) Do(url string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("X-Auth-Token", os.Getenv("API_KEY"))

	return c.HTTPClient.Do(req)
}
