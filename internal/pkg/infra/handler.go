package infra

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type BaseHandler struct{}

type CodeData struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

type CodeMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Error struct {
	Error string `json:"error"`
}

// 成功返回
// swagger:response ServerSuccess
type ServerSuccess struct {
	// in: body
	Body CodeData
}

// 客服端错误
// swagger:response ClientError
type ClientError struct {
	// in: body
	Body CodeMessage
}

// 服务器错误
// swagger:response ServerError
type ServerError struct {
	// in: body
	Body Error
}

// 成功
func (BaseHandler) Success(c *gin.Context, data interface{}) {
	// Store response for logging middleware
	c.Set("resp_data_for_log", data)

	c.JSON(http.StatusOK, CodeData{
		Code: 0,
		Data: data,
	})
}

// 成功(自定义code)
func (BaseHandler) ErrorCode(c *gin.Context, code int, message string) {
	// Store response for logging middleware
	c.Set("resp_data_for_log", message)

	c.JSON(http.StatusOK, CodeMessage{
		Code:    code,
		Message: message,
	})
}

// 参数无效
func (BaseHandler) InvalidParameter(c *gin.Context, msg string) {
	// Store response for logging middleware
	c.Set("resp_data_for_log", msg)

	c.JSON(http.StatusBadRequest, Error{
		msg,
	})
}

// 无权限
func (BaseHandler) Unauthorized(c *gin.Context, msg string) {
	// Store response for logging middleware
	c.Set("resp_data_for_log", msg)

	c.JSON(http.StatusUnauthorized, Error{
		msg,
	})
}

// 服务异常
func (BaseHandler) ServerError(c *gin.Context, msg string) {
	// Store response for logging middleware
	c.Set("resp_data_for_log", msg)

	c.JSON(http.StatusInternalServerError, Error{
		msg,
	})
}

func (b BaseHandler) HandleError(c *gin.Context, err error) {
	switch err {
	case gorm.ErrRecordNotFound:
		b.ErrorCode(c, 6001, "记录不存在")
	case gorm.ErrInvalidTransaction:
		b.ErrorCode(c, 6002, err.Error())
	case gorm.ErrNotImplemented:
		b.ErrorCode(c, 6003, err.Error())
	case gorm.ErrMissingWhereClause:
		b.ErrorCode(c, 6004, err.Error())
	case gorm.ErrUnsupportedRelation:
		b.ErrorCode(c, 6005, err.Error())
	case gorm.ErrPrimaryKeyRequired:
		b.ErrorCode(c, 6006, err.Error())
	case gorm.ErrModelValueRequired:
		b.ErrorCode(c, 6007, err.Error())
	case gorm.ErrInvalidData:
		b.ErrorCode(c, 6008, err.Error())
	case gorm.ErrUnsupportedDriver:
		b.ErrorCode(c, 6009, err.Error())
	case gorm.ErrRegistered:
		b.ErrorCode(c, 6010, err.Error())
	case gorm.ErrInvalidField:
		b.ErrorCode(c, 6011, err.Error())
	default:
		b.ServerError(c, err.Error())
	}
}

// 绑定参数到结构体
func (BaseHandler) BindParams(c *gin.Context, d interface{}) error {
	if err := c.ShouldBind(d); err != nil {
		// Invalid params
		return err
	}

	// Translate error message
	err := validate.Struct(d)
	if err != nil {
		if _, ok := err.(validator.ValidationErrors); ok {
			var sliceErrs []string
			for _, err := range err.(validator.ValidationErrors) {
				sliceErrs = append(sliceErrs, err.Translate(translator))
			}
			return errors.New(strings.Join(sliceErrs, ","))
		}
		return err
	}
	return nil
}

// 获取路由内参数
func (BaseHandler) GetParam(c *gin.Context, id string) (intID int, err error) {
	strID := c.Param(id)

	intID, err = strconv.Atoi(strID)

	return
}

// 获取分页参数(query内)
func (BaseHandler) GetPaginationParams(c *gin.Context) (page, size int, err error) {
	// 1.参数校验
	pageStr := c.DefaultQuery("page", "1")
	sizeStr := c.DefaultQuery("per_page", "10")

	page, err = strconv.Atoi(pageStr)

	size, err = strconv.Atoi(sizeStr)

	return
}
