package footballdataapi

type Status string

const (
	Finished  Status = "FINISHED"
	Scheduled Status = "SCHEDULED"
	Live      Status = "LIVE"
	InPlay    Status = "IN_PLAY"
)

type Match struct {
	HomeTeam Team   `json:"homeTeam"`
	AwayTeam Team   `json:"awayTeam"`
	Status   string `json:"status"`
	Score    struct {
		FullTime struct {
			Home int `json:"home"`
			Away int `json:"away"`
		} `json:"fullTime"`
	} `json:"score"`
}
