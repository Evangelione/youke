package middleware

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// 日志中间件
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		// token验证（是否生效，是否过期，是否已过签发时间）
		token, err := jwt.Parse(header, func(token *jwt.Token) (interface{}, error) {
			// 验证加密方式
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			// 返回盐
			return []byte("enterprise"), nil
		})

		// token是否有效
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "请登录",
			})
			c.Abort()
			return
		}

		// 读取token内信息
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			id := claims["id"].(string)
			if id != "" {
				c.Set("id", id)
			} else {
				err = errors.New("无效身份，请重新登录")
			}
		} else {
			err = errors.New("无效token，请重新登录")
		}

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
