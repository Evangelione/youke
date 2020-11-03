package handler

import (
	"yk/internal/app/enterprise/employee"
	"yk/internal/pkg/infra"

	"github.com/gin-gonic/gin"
)

type EmployeeHandler struct {
	infra.BaseHandler
}

// swagger:route GET /employees/{mer_id}/{store_id} 员工 EmployeeWhereParams
//
// 查询员工列表.
//
// 可传入对应筛选项.
//
// Responses:
// 200: ServerSuccess
// 400: ClientError
// 401: ClientError
// 500: ServerError
func (e EmployeeHandler) Index(c *gin.Context) {
	// 1.参数校验
	page, size, err := e.GetPaginationParams(c)
	if err != nil {
		e.InvalidParameter(c, err.Error())
		return
	}

	merID, err := e.GetParam(c, "mer_id")
	if err != nil {
		e.InvalidParameter(c, err.Error())
		return
	}

	storeID, err := e.GetParam(c, "store_id")
	if err != nil {
		e.InvalidParameter(c, err.Error())
		return
	}

	var where employee.WhereParams
	err = e.BindParams(c, &where)
	if err != nil {
		e.InvalidParameter(c, err.Error())
		return
	}

	// 2.调用service
	countRows, err := employee.FindAndCountAll(merID, storeID, page, size, where)
	if err != nil {
		e.HandleError(c, err)
		return
	}

	// 3.返回结果
	e.Success(c, countRows)
}

// swagger:route GET /employees/{mer_id}/{store_id}/{emp_id} 员工 EmployeePathParams
//
// 查询单个员工信息.
//
// Responses:
// 200: ServerSuccess
// 400: ClientError
// 401: ClientError
// 500: ServerError
func (e EmployeeHandler) Show(c *gin.Context) {
	// 1.参数校验
	merID, err := e.GetParam(c, "mer_id")
	if err != nil {
		e.InvalidParameter(c, err.Error())
		return
	}

	storeID, err := e.GetParam(c, "store_id")
	if err != nil {
		e.InvalidParameter(c, err.Error())
		return
	}

	id, err := e.GetParam(c, "emp_id")
	if err != nil {
		e.InvalidParameter(c, err.Error())
		return
	}

	// 2.调用service
	one, err := employee.FindOne(merID, storeID, id)
	if err != nil {
		e.HandleError(c, err)
		return
	}

	// 3.返回结果
	e.Success(c, one)
}

// swagger:route POST /employees/{mer_id}/{store_id} 员工 EmployeeCreateParams
//
// 创建员工.
//
// Responses:
// 200: ServerSuccess
// 400: ClientError
// 401: ClientError
// 500: ServerError
func (e EmployeeHandler) Create(c *gin.Context) {
	// 1.参数校验
	merID, err := e.GetParam(c, "mer_id")
	if err != nil {
		e.InvalidParameter(c, err.Error())
		return
	}

	storeID, err := e.GetParam(c, "store_id")
	if err != nil {
		e.InvalidParameter(c, err.Error())
		return
	}

	var params employee.CreateParams
	err = e.BindParams(c, &params.Body)
	if err != nil {
		e.InvalidParameter(c, err.Error())
		return
	}

	// 2.调用service
	create, err := employee.Create(merID, storeID, params)
	if err != nil {
		e.HandleError(c, err)
		return
	}

	// 3.返回结果
	e.Success(c, create)
}

// swagger:route PUT /employees/{mer_id}/{store_id}/{emp_id} 员工 EmployeeUpdateParams
//
// 更新员工信息.
//
// Responses:
// 200: ServerSuccess
// 400: ClientError
// 401: ClientError
// 500: ServerError
func (e EmployeeHandler) Update(c *gin.Context) {
	// 1.参数校验
	merID, err := e.GetParam(c, "mer_id")
	if err != nil {
		e.InvalidParameter(c, err.Error())
		return
	}

	storeID, err := e.GetParam(c, "store_id")
	if err != nil {
		e.InvalidParameter(c, err.Error())
		return
	}

	id, err := e.GetParam(c, "emp_id")
	if err != nil {
		e.InvalidParameter(c, err.Error())
		return
	}

	var body employee.UpdateParams
	if err := e.BindParams(c, &body); err != nil {
		e.InvalidParameter(c, err.Error())
		return
	}

	// 2.调用service
	update, err := employee.Update(merID, storeID, id, body)
	if err != nil {
		e.HandleError(c, err)
		return
	}

	// 3.返回结果
	e.Success(c, update)
}

// swagger:route DELETE /employees/{mer_id}/{store_id}/{emp_id} 员工 EmployeeDeleteParams
//
// 删除员工.
//
// Responses:
// 200: ServerSuccess
// 400: ClientError
// 401: ClientError
// 500: ServerError
func (e EmployeeHandler) Destroy(c *gin.Context) {
	// 1.参数校验
	merID, err := e.GetParam(c, "mer_id")
	if err != nil {
		e.InvalidParameter(c, err.Error())
		return
	}

	storeID, err := e.GetParam(c, "store_id")
	if err != nil {
		e.InvalidParameter(c, err.Error())
		return
	}

	id, err := e.GetParam(c, "emp_id")
	if err != nil {
		e.InvalidParameter(c, err.Error())
		return
	}

	// 2.调用service
	if err = employee.Delete(merID, storeID, id); err != nil {
		e.HandleError(c, err)
		return
	}

	// 3.返回结果
	e.Success(c, nil)
}

// swagger:route GET /merchants/{mer_id}/employees 商户 EmployeeWhereParamsOfMerchant
//
// 查询商户员工列表.
//
// Responses:
// 200: ServerSuccess
// 400: ClientError
// 401: ClientError
// 500: ServerError
func (e EmployeeHandler) IndexOfMerchant(c *gin.Context) {
	// 1.参数校验
	id, err := e.GetParam(c, "mer_id")
	if err != nil {
		e.InvalidParameter(c, err.Error())
		return
	}

	page, size, err := e.GetPaginationParams(c)
	if err != nil {
		e.InvalidParameter(c, err.Error())
		return
	}

	var where employee.WhereParamsOfMerchant
	err = e.BindParams(c, &where)
	if err != nil {
		e.InvalidParameter(c, err.Error())
		return
	}

	// 2.调用service
	countRows, err := employee.FindAndCountAllOfMerchant(id, page, size, where)
	if err != nil {
		e.HandleError(c, err)
		return
	}

	// 3.返回结果
	e.Success(c, countRows)
}
