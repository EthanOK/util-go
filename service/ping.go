package service

import (
	"fmt"
	"strings"

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
	fmt.Println("token:", token)
	// 处理Ping请求
	c.JSON(200, gin.H{
		"code": 200,
		"data": "pong",
	})
}
