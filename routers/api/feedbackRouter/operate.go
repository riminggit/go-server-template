package feedbackRouter

import (
	"github.com/gin-gonic/gin"
	"go-server-template/middleware/global/auth"
	"go-server-template/pkg/apiMap"
	"go-server-template/src/feedback"
)

func userFeedbackRouter(g *gin.RouterGroup) {
	adminAuth := authMiddleware.UserAuth()
	g.POST(apiMap.POST_QUERY_FEEDBACK, adminAuth, feedback.QueryController)
	g.POST(apiMap.POST_DELETE_FEEDBACK, adminAuth, feedback.DeleteController)
	g.POST(apiMap.POST_UPDATE_FEEDBACK, adminAuth, feedback.AnswerFeedbackController)
	
	g.POST(apiMap.POST_CREATE_FEEDBACK_USER, feedback.CreateController)
	g.POST(apiMap.POST_QUERY_FEEDBACK_USER, feedback.UserQueryController)
	g.POST(apiMap.POST_DELETE_FEEDBACK_USER, feedback.UserDeleteController)
	g.POST(apiMap.POST_UPDATE_FEEDBACK_USER, feedback.UserUpdateController)
}
