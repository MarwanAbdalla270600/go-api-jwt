package entity

import "time"

type Rental struct {
	Id         string
	UserID     string
	CarID      string
	StartDate  time.Time
	EndDate    time.Time
	TotalPrice float32
	Status     string // ACTIVE | UPCOMING | COMPLETED | CANCELLED
}
 