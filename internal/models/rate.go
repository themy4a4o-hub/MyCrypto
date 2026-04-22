package models

import "time"

type Rates struct {
	Id             int       `db:"id"`
	Cryptocurrency string    `db:"cryptocurrency"`
	Price          float64   `db:"price"`
	Timestamp      time.Time `db:"timestamp"`
}
