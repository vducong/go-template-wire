package router

import (
	"go-template-wire/internal/controller"
	"time"

	"github.com/gin-gonic/gin"
)

type healthCheckRouter struct {
	group      *gin.RouterGroup
	controller *controller.HealthCheckController
}

func initHealthCheckRouter(
	group *gin.RouterGroup,
	ctrl *controller.HealthCheckController,
) {
	router := newHealthCheckRouter(group, ctrl)
	router.handle()
}

func newHealthCheckRouter(
	group *gin.RouterGroup,
	ctrl *controller.HealthCheckController,
) *healthCheckRouter {
	return &healthCheckRouter{group, ctrl}
}

func (r healthCheckRouter) handle() {
	root := r.group.Group("/health")
	root.GET("", r.controller.HealthCheck(time.Now()))
}
