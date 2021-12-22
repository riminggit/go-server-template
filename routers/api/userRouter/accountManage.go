package userRouter

import (
	"github.com/gin-gonic/gin"
	"go-server-template/src/accountManage"
	"go-server-template/pkg/apiMap"
)

func AccountManageRouter(g *gin.RouterGroup) {
	g.POST(apiMap.POST_USER_CHANGE_PHONE, account.ChangePhoneController)
}
