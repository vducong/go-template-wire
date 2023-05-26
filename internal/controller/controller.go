package controller

import (
	"go-template-wire/pkg/failure"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

type Controllers struct {
	HealthCheck  *HealthCheckController
	ReusableCode *ReusableCodeController
}

var ControllerSet = wire.NewSet(
	NewHealthCheckController,
	NewReusableCodeController,
	wire.Struct(new(Controllers), "HealthCheck", "ReusableCode"),
)

func BindJSON[B interface{}](ctx *gin.Context) (bindedBody *B, err *failure.BindJSONErr) {
	var body B
	if err := ctx.ShouldBindJSON(&body); err != nil {
		return nil, &failure.BindJSONErr{
			OriginalErr: err,
			Model:       reflect.TypeOf(body),
		}
	}
	return &body, nil
}
