package models

type RateStats struct {
	Cryptocurrency    string   `json:"cryptocurrency"`
	CurrentPrice      float64  `json:"current_price"`
	MinPriceToday     float64  `json:"min_price"`
	MaxPriceToday     float64  `json:"max_price"`
	ChangePercentHour *float64 `json:"change_percent_hour"`
}
