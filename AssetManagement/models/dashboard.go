package models

//Dashboard represent needed data to dashboard
type Dashboard struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

//Statistics represent needed data to statistics dashboard
type Statistics struct {
	Month    string `json:"month"`
	Quantity int    `json:"quantity"`
}
