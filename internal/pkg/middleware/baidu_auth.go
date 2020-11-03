package middleware

import (
	"net/http"
	"time"
	"yk/internal/pkg/constants"
	"yk/internal/pkg/infra"

	"github.com/imroc/req"

	"github.com/gin-gonic/gin"
)

// unit鉴权
func UnitAccessToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Read AccessToken.
		accessToken := infra.GetRedisValue(constants.DialogAccessTokenKey)
		if accessToken != "" {
			c.Next()
			return
		}

		// No AccessToken, Get AccessToken.
		r, err := req.Post(constants.AccessTokenAPI, req.Param{
			"client_id":     constants.DiaLogClientID,
			"client_secret": constants.DialogClientSecret,
			"grant_type":    constants.GrantType,
		})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			c.Abort()
			return
		}

		// Conversion json to Struct or Map.
		token := struct {
			ExpiresIn   int64  `json:"expires_in"`
			AccessToken string `json:"access_token"`
		}{}
		if err = r.ToJSON(&token); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			c.Abort()
			return
		}

		// Store AccessToken.
		if err = infra.SetRedisValue(constants.DialogAccessTokenKey, token.AccessToken, time.Duration(token.ExpiresIn)*time.Second); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			c.Abort()
		}
	}
}

// 人脸鉴权
func FaceAccessToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Read AccessToken.
		accessToken := infra.GetRedisValue(constants.FaceAccessTokenKey)
		if accessToken != "" {
			c.Next()
			return
		}

		// No AccessToken, Get AccessToken.
		r, err := req.Post(constants.AccessTokenAPI, req.Param{
			"client_id":     constants.FaceClientID,
			"client_secret": constants.FaceClientSecret,
			"grant_type":    constants.GrantType,
		})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			c.Abort()
			return
		}

		// Conversion json to Struct or Map.
		token := struct {
			ExpiresIn   int64  `json:"expires_in"`
			AccessToken string `json:"access_token"`
		}{}
		if err = r.ToJSON(&token); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			c.Abort()
			return
		}

		// Store AccessToken.
		if err = infra.SetRedisValue(constants.FaceAccessTokenKey, token.AccessToken, time.Duration(token.ExpiresIn)*time.Second); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			c.Abort()
		}
	}
}
