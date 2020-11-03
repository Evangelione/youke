package employee

import "gorm.io/gorm"

type Employee struct {
	DB        *gorm.DB `json:"-" gorm:"-"`
	ID        int      `json:"id"`
	MerID     int      `json:"mer_id"`
	StoreID   int      `json:"store_id"`
	Name      string   `json:"name"`
	Tel       string   `json:"tel"`
	Email     *string  `json:"email"`
	Status    int      `json:"status"`
	CreatedAt int64    `json:"created_at"`
	UpdatedAt int64    `json:"updated_at"`
	DeletedAt *int64   `json:"deleted_at"`
}
