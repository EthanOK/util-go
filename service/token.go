package service

import (
	"strings"
	"util-go/utils"

	"github.com/gin-gonic/gin"
)

// GetAuthorization
// @Tags Token
// @Accept json
// @Produce json
// @param name query string true "name"
// @param password query string true "password"
// @Success 200 {string} json{"code","data"}
// @Router /getAuthorization [get]
func GetAuthorization(c *gin.Context) {
	// 获取请求参数
	name := c.Query("name")
	password := c.Query("password")

	auth := utils.GetAuthorization(name, password)

	// 处理Ping请求
	c.JSON(200, gin.H{
		"code": 200,
		"data": auth,
	})
}

// ParseAuthorization
// @Tags Token
// @Accept json
// @Produce json
// @param authorization query string true "authorization"
// @Success 200 {string} json{"code","data"}
// @Router /parseAuthorization [get]
func ParseAuthorization(c *gin.Context) {
	// 获取请求参数
	authorization := c.Query("authorization")

	if len(authorization) == 0 {
		c.JSON(401, gin.H{
			"code": 401,
			"data": "Unauthorized",
		})
		return
	}
	token := strings.Fields(authorization)[1]

	claims, err := utils.ParseToken(token)

	if err != nil {

		c.JSON(401, gin.H{
			"code": 401,
			"data": err.Error(),
		})
		return

	}

	c.JSON(200, gin.H{
		"code": 200,
		"data": claims.Name,
	})
}
