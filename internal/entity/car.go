package entity

import "time"

type Car struct {
	Id          string    `db:"id" json:"id"`
	Slug        string    `db:"slug" json:"slug"`
	Brand       string    `db:"brand" json:"brand"`
	Model       string    `db:"model" json:"model"`
	Price       float32   `db:"price_per_day" json:"pricePerDay"`
	Description string    `db:"description" json:"description"`
	Thumbnail   string    `db:"thumbnail" json:"thumbnail"`
	Images      []string  `db:"images" json:"images"`
	CreatedAt   time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt   time.Time `db:"updated_at" json:"updatedAt"`
}
