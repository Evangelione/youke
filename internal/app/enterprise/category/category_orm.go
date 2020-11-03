package category

import "gorm.io/gorm"

type Category struct {
	DB    *gorm.DB `json:"-" gorm:"-"`
	ID    int      `json:"id" gorm:"column:cat_id;primaryKey"`
	PID   int      `json:"pid" gorm:"column:cat_fid"`
	Name  string   `json:"name" gorm:"column:cat_name"`
	IsHot int      `json:"is_hot"`
}

func (Category) TableName() string {
	return "maycms_merchant_category"
}
