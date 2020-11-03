package handler

import (
	"yk/internal/app/enterprise/merchant"
	"yk/internal/pkg/infra"

	"github.com/gin-gonic/gin"
)

type MerchantHandler struct {
	infra.BaseHandler
}

// swagger:route GET /merchants 商户 MerchantWhereParams
//
// 查询商户列表.
//
// 可传入对应筛选项.
//
// Responses:
// 200: ServerSuccess
// 400: ClientError
// 401: ClientError
// 500: ServerError
func (m MerchantHandler) Index(c *gin.Context) {
	// 1.参数校验
	page, size, err := m.GetPaginationParams(c)
	if err != nil {
		m.InvalidParameter(c, err.Error())
		return
	}

	var where merchant.WhereParams
	err = m.BindParams(c, &where)
	if err != nil {
		m.InvalidParameter(c, err.Error())
		return
	}

	// 2.调用service
	countRows, err := merchant.FindAll(page, size, where)
	if err != nil {
		m.HandleError(c, err)
		return
	}

	// 3.返回结果
	m.Success(c, countRows)
}

// swagger:route GET /merchants/{mer_id} 商户 MerchantPathParams
//
// 查询单个商户信息.
//
// Responses:
// 200: ServerSuccess
// 400: ClientError
// 401: ClientError
// 500: ServerError
func (m MerchantHandler) Show(c *gin.Context) {
	// 1.参数校验
	id, err := m.GetParam(c, "mer_id")
	if err != nil {
		m.InvalidParameter(c, err.Error())
		return
	}

	// 2.调用service
	one, err := merchant.FindOne(id)
	if err != nil {
		m.HandleError(c, err)
		return
	}

	// 3.返回结果
	m.Success(c, one)
}

// swagger:route POST /merchants 商户 MerchantCreateParams
//
// 创建商户.
//
// Responses:
// 200: ServerSuccess
// 400: ClientError
// 401: ClientError
// 500: ServerError
func (m MerchantHandler) Create(c *gin.Context) {
	// 1.参数校验
	var params merchant.CreateParams
	err := m.BindParams(c, &params.Body)
	if err != nil {
		m.InvalidParameter(c, err.Error())
		return
	}

	// 2.调用service
	create, err := merchant.Create(params)
	if err != nil {
		m.HandleError(c, err)
		return
	}

	// 3.返回结果
	m.Success(c, create)
}

// swagger:route PUT /merchants/{mer_id} 商户 MerchantUpdateParams
//
// 更新商户信息.
//
// Responses:
// 200: ServerSuccess
// 400: ClientError
// 401: ClientError
// 500: ServerError
func (m MerchantHandler) Update(c *gin.Context) {
	// 1.参数校验
	id, err := m.GetParam(c, "mer_id")
	if err != nil {
		m.InvalidParameter(c, err.Error())
		return
	}

	var params merchant.UpdateParams
	if err := m.BindParams(c, &params.Body); err != nil {
		m.InvalidParameter(c, err.Error())
		return
	}

	// 2.调用service
	update, err := merchant.Update(id, params)
	if err != nil {
		m.HandleError(c, err)
		return
	}

	// 3.返回结果
	m.Success(c, update)
}

// swagger:route DELETE /merchants/{mer_id} 商户 MerchantDeleteParams
//
// 删除商户.
//
// Responses:
// 200: ServerSuccess
// 400: ClientError
// 401: ClientError
// 500: ServerError
func (m MerchantHandler) Destroy(c *gin.Context) {
	// 1.参数校验
	id, err := m.GetParam(c, "mer_id")
	if err != nil {
		m.InvalidParameter(c, err.Error())
		return
	}

	// 2.调用service
	if err = merchant.Delete(id); err != nil {
		m.HandleError(c, err)
		return
	}

	// 3.返回结果
	m.Success(c, nil)
}
