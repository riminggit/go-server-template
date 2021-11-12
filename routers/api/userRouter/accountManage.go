package userRouter

import (
	"github.com/gin-gonic/gin"
	"go-server-template/src/accountManage"
)

func AccountManageRouter(g *gin.RouterGroup) {
	g.POST("/account-manage/change-phone", account.ChangePhoneController)
}
