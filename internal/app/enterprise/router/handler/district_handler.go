package handler

import (
	"yk/internal/app/enterprise/district"
	"yk/internal/pkg/infra"

	"github.com/gin-gonic/gin"
)

type DistrictHandler struct {
	infra.BaseHandler
}

// swagger:route GET /districts 地区 DistrictWhereParams
//
// 查询省份列表.
//
// Responses:
// 200: ServerSuccess
// 400: ClientError
// 401: ClientError
// 500: ServerError
func (d DistrictHandler) Index(c *gin.Context) {
	// 1.参数校验
	var where district.WhereParams
	err := d.BindParams(c, &where)
	if err != nil {
		d.InvalidParameter(c, err.Error())
		return
	}

	// 2.调用service
	all, err := district.FindAll(where)
	if err != nil {
		d.HandleError(c, err)
		return
	}

	// 3.返回结果
	d.Success(c, all)
}

// swagger:route GET /districts/{id} 地区 DistrictPathParams
//
// 查询单个地区信息.
//
// Responses:
// 200: ServerSuccess
// 400: ClientError
// 401: ClientError
// 500: ServerError
func (d DistrictHandler) Show(c *gin.Context) {
	// 1.参数校验
	id, err := d.GetParam(c, "id")
	if err != nil {
		d.InvalidParameter(c, err.Error())
		return
	}

	// 2.调用service
	one, err := district.FindOne(id)
	if err != nil {
		d.HandleError(c, err)
		return
	}

	// 3.返回结果
	d.Success(c, one)
}

func (d DistrictHandler) Create(c *gin.Context) {
	d.Success(c, "Create")
}

func (d DistrictHandler) Update(c *gin.Context) {
	d.Success(c, "Update")
}

func (d DistrictHandler) Destroy(c *gin.Context) {
	d.Success(c, "Destroy")
}
