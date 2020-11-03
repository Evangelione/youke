package handler

import (
	"fmt"
	"yk/internal/app/enterprise/store"
	"yk/internal/pkg/infra"

	"github.com/gin-gonic/gin"
)

type StoreHandler struct {
	infra.BaseHandler
}

// swagger:route GET /stores/{mer_id} 店铺 StoreWhereParams
//
// 查询店铺列表.
//
// 可传入对应筛选项.
//
// Responses:
// 200: ServerSuccess
// 400: ClientError
// 401: ClientError
// 500: ServerError
func (s StoreHandler) Index(c *gin.Context) {
	// 1.参数校验
	page, size, err := s.GetPaginationParams(c)
	if err != nil {
		s.InvalidParameter(c, err.Error())
		return
	}

	pid, err := s.GetParam(c, "mer_id")
	if err != nil {
		s.InvalidParameter(c, err.Error())
		return
	}

	var where store.WhereParams
	err = s.BindParams(c, &where)
	if err != nil {
		s.InvalidParameter(c, err.Error())
		return
	}

	// 2.调用service
	countRows, err := store.FindAndCountAll(pid, page, size, where)
	if err != nil {
		s.HandleError(c, err)
		return
	}

	// 3.返回结果
	s.Success(c, countRows)
}

// swagger:route GET /stores/{mer_id}/{store_id} 店铺 StorePathParams
//
// 查询单个店铺信息.
//
// Responses:
// 200: ServerSuccess
// 400: ClientError
// 401: ClientError
// 500: ServerError
func (s StoreHandler) Show(c *gin.Context) {
	accept := c.GetHeader("Accept")
	fmt.Println(accept)
	// 1.参数校验
	pid, err := s.GetParam(c, "mer_id")
	if err != nil {
		s.InvalidParameter(c, err.Error())
		return
	}

	id, err := s.GetParam(c, "store_id")
	if err != nil {
		s.InvalidParameter(c, err.Error())
		return
	}

	// 2.调用service
	one, err := store.FindOne(pid, id)
	if err != nil {
		s.HandleError(c, err)
		return
	}

	// 3.返回结果
	s.Success(c, one)
}

// swagger:route POST /stores/{mer_id} 店铺 StoreCreateParams
//
// 创建店铺.
//
// Responses:
// 200: ServerSuccess
// 400: ClientError
// 401: ClientError
// 500: ServerError
func (s StoreHandler) Create(c *gin.Context) {
	// 1.参数校验
	pid, err := s.GetParam(c, "mer_id")
	if err != nil {
		s.InvalidParameter(c, err.Error())
		return
	}

	var params store.CreateParams
	err = s.BindParams(c, &params.Body)
	if err != nil {
		s.InvalidParameter(c, err.Error())
		return
	}

	// 2.调用service
	create, err := store.Create(pid, params)
	if err != nil {
		s.HandleError(c, err)
		return
	}

	// 3.返回结果
	s.Success(c, create)
}

// swagger:route PUT /stores/{mer_id}/{store_id} 店铺 StoreUpdateParams
//
// 更新店铺信息.
//
// Responses:
// 200: ServerSuccess
// 400: ClientError
// 401: ClientError
// 500: ServerError
func (s StoreHandler) Update(c *gin.Context) {
	// 1.参数校验
	pid, err := s.GetParam(c, "mer_id")
	if err != nil {
		s.InvalidParameter(c, err.Error())
		return
	}

	id, err := s.GetParam(c, "store_id")
	if err != nil {
		s.InvalidParameter(c, err.Error())
		return
	}

	var body store.UpdateParams
	if err := s.BindParams(c, &body); err != nil {
		s.InvalidParameter(c, err.Error())
		return
	}

	// 2.调用service
	update, err := store.Update(pid, id, body)
	if err != nil {
		s.HandleError(c, err)
		return
	}

	// 3.返回结果
	s.Success(c, update)
}

// swagger:route DELETE /stores/{mer_id}/{store_id} 店铺 StoreDeleteParams
//
// 删除店铺.
//
// Responses:
// 200: ServerSuccess
// 400: ClientError
// 401: ClientError
// 500: ServerError
func (s StoreHandler) Destroy(c *gin.Context) {
	// 1.参数校验
	pid, err := s.GetParam(c, "mer_id")
	if err != nil {
		s.InvalidParameter(c, err.Error())
		return
	}

	id, err := s.GetParam(c, "store_id")
	if err != nil {
		s.InvalidParameter(c, err.Error())
		return
	}

	// 2.调用service
	if err = store.Delete(pid, id); err != nil {
		s.HandleError(c, err)
		return
	}

	// 3.返回结果
	s.Success(c, nil)
}
