package service

import (
	"strings"
	"util-go/utils"

	"github.com/gin-gonic/gin"
)

// Ping
// @Tags Ping
// @Accept json
// @Produce json
// @Success 200 {string} json{"code","data"}
// @Router /ping [get]
func Ping(c *gin.Context) {

	auth := c.Request.Header.Get("Authorization")
	if len(auth) == 0 {
		c.JSON(401, gin.H{
			"code": 401,
			"data": "Unauthorized",
		})
		return
	}

	token := strings.Fields(auth)[1]
	_, err := utils.ParseToken(token)

	if err != nil {
		c.JSON(401, gin.H{
			"code": 401,
			"data": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"code": 200,
		"data": "pong",
	})
}
