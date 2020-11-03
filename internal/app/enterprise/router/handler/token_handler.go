package handler

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
	"yk/internal/app/enterprise/merchant"
	"yk/internal/pkg/infra"

	"github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
)

type TokenHandler struct {
	infra.BaseHandler
}

// swagger:parameters MerchantCreateParams
type CreateParams struct {
	// in: body
	Body struct {
		Account  string `json:"account" validate:"required,alphanum,min=6,max=12"`
		Password string `json:"password" validate:"required,min=6,max=18"`
		//ClientID int    `json:"status" validate:"required,oneof=0 1"`
		//Contact      string    `json:"contact" validate:"required,alphanumunicode"`
	}
}

func (t TokenHandler) Index(c *gin.Context) {

}

func (t TokenHandler) Show(c *gin.Context) {

}

func (t TokenHandler) Create(c *gin.Context) {
	// 1.参数校验
	var params CreateParams
	err := t.BindParams(c, &params.Body)
	if err != nil {
		t.InvalidParameter(c, err.Error())
		return
	}

	fmt.Println(params)

	// 2.调用service
	m := merchant.Merchant{
		DB: infra.Mysql,
	}

	hash := md5.New()
	hash.Write([]byte(params.Body.Password))
	pwd := hex.EncodeToString(hash.Sum(nil))

	if err := m.DB.Where("account = ? AND pwd = ?", params.Body.Account, pwd).First(&m).Error; err != nil {
		t.ErrorCode(c, 2001, "账号或密码不存在")
		return
	}

	// 获取当前时间
	now := time.Now()
	h2, _ := time.ParseDuration("2h")

	// 设置过期时间
	exp := now.Add(h2).Unix()
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "enterprise",
		"exp": exp,
		"sub": "admin-api",
		"aud": "merchant",
		"nbf": now.Unix(),
		"iat": now.Unix(),
		"id":  m.ID,
	})

	tk, err := claims.SignedString([]byte("enterprise"))
	if err != nil {
		t.HandleError(c, err)
		return
	}

	// 3.返回结果
	t.Success(c, tk)
}

func (t TokenHandler) Update(c *gin.Context) {

}

func (t TokenHandler) Destroy(c *gin.Context) {

}
