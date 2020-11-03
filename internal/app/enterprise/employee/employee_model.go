package employee

import "yk/internal/pkg/infra"

// swagger:parameters EmployeeWhereParams
type WhereParams struct {
	infra.Pagination `structs:"-"`
	// 商户id
	// in:path
	MerID int
	// 店铺id
	// in: path
	StoreID int
	// 员工名称
	Name string `form:"name" json:"name" structs:"name,omitempty"`
}

// swagger:parameters EmployeePathParams
type PathParams struct {
	// 商户id
	// in: path
	MerID int
	// 店铺id
	// in: path
	StoreID int
}

// swagger:parameters EmployeeCreateParams
type CreateParams struct {
	// 商户id
	// in: path
	MerID int
	// 店铺id
	// in: path
	StoreID int
	// in: body
	Body struct {
		Name  string  `json:"name" validate:"required"`
		Tel   string  `json:"tel" validate:"required,number,len=11"`
		Email *string `json:"email" validate:"omitempty,email"`
	}
}

// swagger:parameters EmployeeUpdateParams
type UpdateParams struct {
	// 商户id
	// in: path
	MerID int
	// 店铺id
	// in: path
	StoreID int
	// 员工id
	// in: path
	EmpID int
	// in: body
	Body struct {
		Name  string `json:"name" validate:"omitempty" structs:"name,omitempty" `
		Tel   string `json:"tel" validate:"omitempty,number,len=11" structs:"tel,omitempty"`
		Email string `json:"email" validate:"omitempty,email" structs:"email,omitempty"`
	}
}

// swagger:parameters EmployeeDeleteParams
type DeleteParams struct {
	// 商户id
	// in: path
	MerID int
	// 店铺id
	// in: path
	StoreID int
	// 员工id
	// in: path
	EmpID int
}

// swagger:parameters EmployeeWhereParamsOfMerchant
type WhereParamsOfMerchant struct {
	infra.Pagination `structs:"-"`
	// 商户id
	// in: path
	MerID int
}
