package router

import (
	"go-template-wire/internal/controller"
	"go-template-wire/internal/middleware"

	"github.com/gin-gonic/gin"
)

type reusableCodeRouter struct {
	group        *gin.RouterGroup
	controller   *controller.ReusableCodeController
	internalAuth *middleware.InternalAuthMiddleware
}

func initReusableCodeRouter(
	group *gin.RouterGroup,
	ctrl *controller.ReusableCodeController,
	internalAuth *middleware.InternalAuthMiddleware,
) {
	router := newReusableCodeRouter(group, ctrl, internalAuth)
	router.handle()
}

func newReusableCodeRouter(
	group *gin.RouterGroup,
	ctrl *controller.ReusableCodeController,
	internalAuth *middleware.InternalAuthMiddleware,
) *reusableCodeRouter {
	return &reusableCodeRouter{group, ctrl, internalAuth}
}

func (r reusableCodeRouter) handle() {
	root := r.group.Group("/reusable-code")
	root.POST("", r.controller.GetByCode)
}
