// Package router YK-Enterprise API.
//
// golang server of enterprise
//
//     Schemes: http, https
//     Host: 192.168.110.178:8905
//     BasePath: /
//     Version: 0.0.1
//
//     Consumes:
//     - application/json
//     - multipart/form-data
//
//     Produces:
//     - application/json
//
// swagger:meta
package router

import (
	"yk/internal/app/enterprise/router/handler"
	"yk/internal/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func Register(engine *gin.Engine) {
	engine.Use(middleware.Cors())

	// Auth
	TokenGroup := engine.Group("/tokens")
	{
		//TokenGroup.GET("", handler.MerchantHandler{}.Index)
		//TokenGroup.GET("/:mer_id", handler.MerchantHandler{}.Show)
		TokenGroup.POST("", handler.TokenHandler{}.Create)
		//TokenGroup.PUT("/:mer_id", handler.MerchantHandler{}.Update)
		//TokenGroup.DELETE("/:mer_id", handler.MerchantHandler{}.Destroy)
	}

	// 商户
	MerchantGroup := engine.Group("/merchants")
	{
		MerchantGroup.GET("", handler.MerchantHandler{}.Index)
		MerchantGroup.GET("/:mer_id", handler.MerchantHandler{}.Show)
		MerchantGroup.POST("", handler.MerchantHandler{}.Create)
		MerchantGroup.PUT("/:mer_id", handler.MerchantHandler{}.Update)
		MerchantGroup.DELETE("/:mer_id", handler.MerchantHandler{}.Destroy)

		// 商户下员工
		MerchantGroup.GET("/:mer_id/employees", handler.EmployeeHandler{}.IndexOfMerchant)
	}

	// 店铺
	StoreGroup := engine.Group("/stores/:mer_id")
	{
		StoreGroup.GET("", handler.StoreHandler{}.Index)
		StoreGroup.GET("/:store_id", handler.StoreHandler{}.Show)
		StoreGroup.POST("", handler.StoreHandler{}.Create)
		StoreGroup.PUT("/:store_id", handler.StoreHandler{}.Update)
		StoreGroup.DELETE("/:store_id", handler.StoreHandler{}.Destroy)
	}

	// 员工
	EmployeeGroup := engine.Group("/employees/:mer_id/:store_id")
	{
		EmployeeGroup.GET("", handler.EmployeeHandler{}.Index)
		EmployeeGroup.GET("/:emp_id", handler.EmployeeHandler{}.Show)
		EmployeeGroup.POST("", handler.EmployeeHandler{}.Create)
		EmployeeGroup.PUT("/:emp_id", handler.EmployeeHandler{}.Update)
		EmployeeGroup.DELETE("/:emp_id", handler.EmployeeHandler{}.Destroy)
	}

	// 地区
	DistrictGroup := engine.Group("/districts")
	{
		DistrictGroup.GET("", handler.DistrictHandler{}.Index)
		DistrictGroup.GET("/:id", handler.DistrictHandler{}.Show)
		//DistrictGroup.POST("", handler.DistrictHandler{}.Create)
		//DistrictGroup.PUT("/:id", handler.DistrictHandler{}.Update)
		//DistrictGroup.DELETE("/:id", handler.DistrictHandler{}.Destroy)
	}

	// 分类
	CategoryGroup := engine.Group("/categories")
	{
		CategoryGroup.GET("", handler.CategoryHandler{}.Index)
		CategoryGroup.GET("/:id", handler.CategoryHandler{}.Show)
		//CategoryGroup.POST("", handler.DistrictHandler{}.Create)
		//CategoryGroup.PUT("/:id", handler.DistrictHandler{}.Update)
		//CategoryGroup.DELETE("/:id", handler.DistrictHandler{}.Destroy)
	}
}
