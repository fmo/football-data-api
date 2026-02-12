package footballdataapi

type Status string

const (
	Finished  Status = "FINISHED"
	Scheduled Status = "SCHEDULED"
)

var statuses = map[Status]struct{}{
	Finished:  {},
	Scheduled: {},
}

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
