package merchant

import (
	"yk/internal/pkg/constants"
	"yk/internal/pkg/infra"
	"yk/internal/pkg/lib"

	"gorm.io/gorm"
)

type Merchant struct {
	DB           *gorm.DB        `json:"-" gorm:"-"`
	ID           int             `json:"id" gorm:"column:mer_id;primaryKey"`
	Name         string          `json:"name"`
	Introduction *string         `json:"introduction" gorm:"column:txt_info"`
	Logo         string          `json:"logo" gorm:"column:service_ico"`
	Images       string          `json:"images" gorm:"column:pic_info"`
	Tel          string          `json:"tel" gorm:"column:phone"`
	Email        *string         `json:"email"`
	CategoryID   int             `json:"category_id" gorm:"column:cat_id"`
	CategoryPID  int             `json:"category_pid" gorm:"column:cat_fid"`
	Districts    infra.Districts `json:"districts" gorm:"embedded"`
	CircleID     int             `json:"circle_id"`
	MarketID     int             `json:"market_id"`
	Address      string          `json:"address"`
	Long         float64         `json:"long"`
	Lat          float64         `json:"lat"`
	Account      string          `json:"account"`
	Password     *string         `json:"password" gorm:"column:pwd;<-;->:false"`
	Status       int             `json:"status"`
	CreatedAt    int64           `json:"created_at" gorm:"column:reg_time"`
	UpdatedAt    int64           `json:"updated_at" gorm:"column:last_time"`
	//Contact      string         `json:"contact"`
	//Sort         int            `json:"sort"`
	//DeletedAt    *int64         `json:"deleted_at"`

}

func (Merchant) TableName() string {
	return "maycms_merchant"
}

func (m *Merchant) BeforeCreate(tx *gorm.DB) error {
	crypt := lib.Md5Crypt(*m.Password, constants.MD5Salt)
	m.Password = &crypt
	return nil
}
