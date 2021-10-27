package app

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func GetHeaderToken(c *gin.Context) string {
	initToken := c.Request.Header.Get("Authorization")
	// token := initToken[7:]
	token := strings.Replace(initToken, "Bearer ", "", 1)
	return token
}
