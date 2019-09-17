package models

//StatusCount represent needed data to dashboard
type StatusCount struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

//AssetSummary represent all asset data name and it's quantity to dashboard
type AssetSummary struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}
