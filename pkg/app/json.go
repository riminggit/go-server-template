package app

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"go-server-template/pkg/log"
	"io/ioutil"
)

// 获取json参数
func GetPostJson(c *gin.Context) string {
	// 获取请求json数据
	data, err := c.GetRawData()
	if err != nil {
		logging.Error(err.Error())
	}
	//很关键
	//把读过的字节流重新放到body
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))

	return string(data)

}
