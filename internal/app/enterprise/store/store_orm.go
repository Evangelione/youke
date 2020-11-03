package store

import (
	"yk/internal/pkg/infra"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Store struct {
	DB           *gorm.DB        `json:"-" gorm:"-"`
	ID           int             `json:"id"`
	MerID        int             `json:"mer_id"`
	Name         string          `json:"name"`
	Introduction *string         `json:"introduction"`
	Logo         string          `json:"logo"`
	Images       datatypes.JSON  `json:"images"`
	OpeningHours string          `json:"opening_hours"`
	Contact      string          `json:"contact"`
	Tel          string          `json:"tel"`
	Email        *string         `json:"email"`
	CategoryID   int             `json:"category_id"`
	Districts    infra.Districts `json:"districts" gorm:"embedded"`
	CircleID     int             `json:"circle_id"`
	MarketID     int             `json:"market_id"`
	Address      string          `json:"address"`
	Long         float64         `json:"long"`
	Lat          float64         `json:"lat"`
	Sort         int             `json:"sort"`
	Status       int             `json:"status"`
	CreatedAt    int64           `json:"created_at"`
	UpdatedAt    int64           `json:"updated_at"`
	DeletedAt    *int64          `json:"deleted_at"`
}
