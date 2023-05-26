package controller

import (
	reusablecode "go-template-wire/internal/reusable_code"
	"go-template-wire/pkg/failure"
	"go-template-wire/pkg/logger"
	"go-template-wire/pkg/response"

	"github.com/gin-gonic/gin"
)

type ReusableCodeController struct {
	log    *logger.Logger
	module *reusablecode.Module
}

func NewReusableCodeController(
	log *logger.Logger, m *reusablecode.Module,
) *ReusableCodeController {
	return &ReusableCodeController{log, m}
}

func (c *ReusableCodeController) GetByCode(ctx *gin.Context) {
	body, errBinding := BindJSON[reusablecode.ReusableCodeGetByCodeReq](ctx)
	if errBinding != nil {
		response.ErrBinding(ctx, &failure.BindJSONErr{
			Code:        failure.ErrReusableCodeGetByCodeBinding,
			OriginalErr: failure.ErrWithTrace(errBinding.OriginalErr),
			Model:       errBinding.Model,
		})
		return
	}

	rc, err := c.module.Repo.GetByCode(ctx.Request.Context(), body.Code)
	if err != nil {
		var errCode failure.ErrCode
		if failure.IsSQLRecordNotFound(err) {
			errCode = failure.ErrReusableCodeNotFound
		} else {
			errCode = failure.ErrReusableCodeFailed
		}
		response.ErrApp(ctx, &failure.AppErr{
			Code:        errCode,
			OriginalErr: failure.ErrWithTrace(err),
		})
		return
	}

	response.Success(ctx, rc)
}
