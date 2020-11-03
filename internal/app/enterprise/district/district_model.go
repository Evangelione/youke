package district

// swagger:parameters DistrictWhereParams
type WhereParams struct {
	// 员工名称
	Level int `form:"level" json:"level" validate:"omitempty,min=1,max=2" structs:"level,omitempty"`
}

// swagger:parameters DistrictPathParams
type PathParams struct {
	// 地区id
	// in: path
	ID int
}
