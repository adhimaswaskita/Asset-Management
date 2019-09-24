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
	Increase int    `json:"increment"`
	Decrease int    `json:"decrement"`
}

//Decrement represent needed data to statistics dashboard
type Decrement struct {
	Month    string `json:"month"`
	Decrease int    `json:"decrease"`
}

//RecentInsertion represent needed data to recent insertion
type RecentInsertion struct {
	Category string `json:"category"`
	Name     string `json:"name"`
	Code     string `json:"code"`
	Time     string `json:"time"`
}
