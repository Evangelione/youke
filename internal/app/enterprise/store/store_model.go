package store

import "yk/internal/pkg/infra"

// swagger:parameters StoreWhereParams
type WhereParams struct {
	infra.Pagination `structs:"-"`
	// 商户id
	// in: path
	MerID int `structs:"-"`
	// 店铺名称
	Name string `form:"name" json:"name" structs:"name,omitempty"`
	// 联系人
	Contact string `form:"contact" json:"contact" structs:"contact,omitempty"`
	// 联系电话
	Tel string `form:"tel" json:"tel" structs:"tel,omitempty"`
	// 分类id
	CategoryID int `form:"category_id" json:"category_id" structs:"category_id,omitempty"`
	// 省份id
	ProvinceID int `form:"province_id" json:"province_id" structs:"province_id,omitempty"`
	// 城市id
	CityID int `form:"city_id" json:"city_id" structs:"city_id,omitempty"`
	// 区域id
	AreaID int `form:"area_id" json:"area_id" structs:"area_id,omitempty"`
	// 商圈id
	CircleID int `form:"circle_id" json:"circle_id" structs:"circle_id,omitempty"`
	// 商盟id
	MarketID int `form:"market_id" json:"market_id" structs:"market_id,omitempty"`
}

// swagger:parameters StorePathParams
type PathParams struct {
	// 商户id
	// in: path
	MerID int
	// 店铺id
	// in: path
	StoreID int
}

// swagger:parameters StoreCreateParams
type CreateParams struct {
	// 商户id
	// in: path
	MerID int
	// in: body
	Body struct {
		Name         string    `json:"name" validate:"required"`
		Introduction *string   `json:"introduction" validate:"omitempty"`
		Logo         string    `json:"logo" validate:"required,url"`
		Images       *[]string `json:"images"`
		OpeningHours string    `json:"opening_hours" validate:"required,ascii"`
		Contact      string    `json:"contact" validate:"required,alphanumunicode"`
		Tel          string    `json:"tel" validate:"required,number,len=11"`
		Email        *string   `json:"email" validate:"omitempty,email"`
		CategoryID   int       `json:"category_id" validate:"required,number"`
		ProvinceID   int       `json:"province_id" validate:"required,number"`
		CityID       int       `json:"city_id" validate:"required,number"`
		AreaID       int       `json:"area_id" validate:"required,number"`
		CircleID     int       `json:"circle_id" validate:"required,number"`
		MarketID     int       `json:"market_id" validate:"required,number"`
		Address      string    `json:"address" validate:"required"`
		Long         string    `json:"long" validate:"required,longitude"`
		Lat          string    `json:"lat" validate:"required,latitude"`
	}
}

// swagger:parameters StoreUpdateParams
type UpdateParams struct {
	// 商户id
	// in: path
	MerID int
	// 店铺id
	// in: path
	StoreID int
	// in: body
	Body struct {
		Name         string   `json:"name" validate:"omitempty" structs:"name,omitempty" `
		Introduction string   `json:"introduction" validate:"omitempty" structs:"introduction,omitempty"`
		Logo         string   `json:"logo" validate:"omitempty,url" structs:"logo,omitempty"`
		Images       []string `json:"images" structs:"images,omitempty"`
		OpeningHours string   `json:"opening_hours" validate:"omitempty,ascii" structs:"opening_hours,omitempty"`
		Contact      string   `json:"contact" validate:"omitempty,alphanumunicode" structs:"contact,omitempty"`
		Tel          string   `json:"tel" validate:"omitempty,number,len=11" structs:"tel,omitempty"`
		Email        string   `json:"email" validate:"omitempty,email" structs:"email,omitempty"`
		CategoryID   int      `json:"category_id" validate:"omitempty,number" structs:"category_id,omitempty"`
		ProvinceID   int      `json:"province_id" validate:"omitempty,number" structs:"province_id,omitempty"`
		CityID       int      `json:"city_id" validate:"omitempty,number" structs:"city_id,omitempty"`
		AreaID       int      `json:"area_id" validate:"omitempty,number" structs:"area_id,omitempty"`
		CircleID     int      `json:"circle_id" validate:"omitempty,number" structs:"circle_id,omitempty"`
		MarketID     int      `json:"market_id" validate:"omitempty,number" structs:"market_id,omitempty"`
		Address      string   `json:"address" validate:"omitempty" structs:"address,omitempty"`
		Long         string   `json:"long" validate:"omitempty,longitude" structs:"long,omitempty"`
		Lat          string   `json:"lat" validate:"omitempty,latitude" structs:"lat,omitempty"`
	}
}

// swagger:parameters StoreDeleteParams
type DeleteParams struct {
	// 商户id
	// in: path
	MerID int
	// 店铺id
	// in: path
	StoreID int
}
