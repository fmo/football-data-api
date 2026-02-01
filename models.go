package footballdataapi

type Area struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	CountryCode  string `json:"countryCode"`
	Flag         string `json:"flag"`
	ParentAreaID int    `json:"parentAreaId"`
	ParentArea   string `json:"parentArea"`
}

type Competition struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Code   string `json:"code"`
	Type   string `json:"type"`
	Emblem string `json:"emblem"`
}

type Season struct {
	ID              int      `json:"id"`
	StartDate       string   `json:"startDate"`
	EndDate         string   `json:"endDate"`
	CurrentMatchDay int      `json:"currentMatchDay"`
	Stages          []string `json:"stages"`
}

type Team struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	ShortName string `json:"shortName"`
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
	Stage string `json:"stage"`
	Type  string `json:"total"`
	Table []Table
}
