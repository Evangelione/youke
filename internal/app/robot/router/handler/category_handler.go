package handler

import (
	"yk/internal/app/enterprise/category"
	"yk/internal/pkg/infra"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	infra.BaseHandler
}

// swagger:route GET /categories 商户 CategoryWhereParams
//
// 查询一级分类列表.
//
// Responses:
// 200: ServerSuccess
// 400: ClientError
// 401: ClientError
// 500: ServerError
func (ca CategoryHandler) Index(c *gin.Context) {
	// 1.参数校验
	// ...

	// 2.调用service
	all, err := category.FindAll()
	if err != nil {
		ca.HandleError(c, err)
		return
	}

	// 3.返回结果
	ca.Success(c, all)
}

// swagger:route GET /categories/{id} 地区 CategoryPathParams
//
// 查询一级分类子集.
//
// Responses:
// 200: ServerSuccess
// 400: ClientError
// 401: ClientError
// 500: ServerError
func (ca CategoryHandler) Show(c *gin.Context) {
	// 1.参数校验
	id, err := ca.GetParam(c, "id")
	if err != nil {
		ca.InvalidParameter(c, err.Error())
		return
	}

	// 2.调用service
	one, err := category.FindOne(id)
	if err != nil {
		ca.HandleError(c, err)
		return
	}

	// 3.返回结果
	ca.Success(c, one)

}

func (ca CategoryHandler) Create(c *gin.Context) {}

func (ca CategoryHandler) Update(c *gin.Context) {}

func (ca CategoryHandler) Destroy(c *gin.Context) {}
