package product

import "time"

type Product struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `json:"name" gorm:"not null;type:varchar(191)"`
	Brand     string `json:"brand" gorm:"not null;type:varchar(191)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
