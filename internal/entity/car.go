package entity

type Car struct {
	Id          string   `db:"id" json:"id"`
	Brand       string   `db:"brand" json:"brand"`
	Model       string   `db:"model" json:"model"`
	Price       float32  `db:"price_per_day" json:"pricePerDay"`
	Description string   `db:"description" json:"description"`
	Thumbnail   string   `db:"thumbnail" json:"thumbnail"`
	Images      []string `db:"images" json:"images"`
}
