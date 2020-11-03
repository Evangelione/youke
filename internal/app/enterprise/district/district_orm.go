package district

import "gorm.io/gorm"

type District struct {
	DB           *gorm.DB `json:"-" gorm:"-"`
	ID           int      `json:"id" gorm:"column:area_id;primaryKey"`
	PID          int      `json:"pid" gorm:"column:area_pid"`
	Deep         int      `json:"deep" gorm:"column:area_type"`
	Name         string   `json:"name" gorm:"column:area_name"`
	PinyinPrefix string   `json:"pinyin_prefix" gorm:"column:first_pinyin"`
	Pinyin       string   `json:"pinyin" gorm:"column:area_url"`       // 没有，暂时占位
	ExtID        int      `json:"ext_id" gorm:"column:area_id"`        // 没有，暂时占位
	ExtName      string   `json:"ext_name" gorm:"column:area_in_desc"` // 没有，暂时占位
}

func (District) TableName() string {
	return "maycms_area"
}
