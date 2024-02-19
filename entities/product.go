package entities

import "time"

type Product struct {
	Id        uint
	Name      string
	Category  Category
	Stock     int
	Price     int
	Desc      string
	ImagePath string
	CreatedAt time.Time
	UpdatedAt time.Time
}
