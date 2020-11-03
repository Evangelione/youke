package infra

// 分页入参
type Pagination struct {
	// 当前页（默认：1）
	Page int `form:"page" json:"page" validate:"omitempty,number,gte=0"`
	// 每页数量（默认：10）
	PerPage int `form:"per_page" json:"per_page" validate:"omitempty,number,gte=0,lte=100"`
}

// 分页出参
type CountRows struct {
	// 总数
	Count int64 `json:"count"`
	// 分页数据
	Rows interface{} `json:"rows"`
}

// orm共用地区参数
type Districts struct {
	ProvinceID int `json:"province_id"`
	CityID     int `json:"city_id"`
	AreaID     int `json:"area_id"`
}

// token参数
type TokenClaims struct {
	Iss  string
	Exp  string
	Sub  string
	Aud  string
	Nbf  string
	Iat  string
	Id   int
	Name string
}
