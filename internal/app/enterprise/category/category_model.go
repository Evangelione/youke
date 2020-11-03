package category

import "yk/internal/pkg/infra"

// swagger:parameters CategoryWhereParams
type WhereParams struct {
	infra.Pagination `structs:"-"`
	// 商户名称
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
}
