package api

import (
	"github.com/gin-gonic/gin"
	"go-server-template/routers/api/classifyRouter"
	"go-server-template/routers/api/userRouter"
	"go-server-template/routers/api/typeRouter"
)

func InitApi(r *gin.Engine) {
	api := r.Group("/api")
	userRouter.UserInitRouter(api)
	classifyRouter.ClassifyInitRouter(api)
	typeRouter.TypeInitRouter(api)
}
