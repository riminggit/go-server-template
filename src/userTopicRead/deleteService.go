package userTopicRead

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteUserTopicRead(c *gin.Context, params DeleteParams) {

}

func DeleteUserTopicReadSimple(c *gin.Context, tx *gorm.DB, params DeleteParams) {

}
