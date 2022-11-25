package template

var ControllerBaseTmp = `package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Api struct{}

func (a *Api) Error(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, map[string]interface{}{
		"code":    4000,
		"message": err.Error(),
	})
}

func (a *Api) Success(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code":    2000,
		"message": "ok",
	})
}

func (a *Api) SuccessWithData(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code":    2000,
		"message": "ok",
		"data":    data,
	})
}

func (a *Api) Fail(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code":    4000,
		"message": message,
	})
}
`

var ControllerUserLoginTmp = `package user

import (
	"errors"
	"{{projectName}}/controller"
	"{{projectName}}/model"
	"{{projectName}}/model/api"
	"{{projectName}}/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CtlLogin struct {
	controller.Api
}

//Login 登录
// ShowAccount godoc
// @Tags         登录
// @Summary      login
// @Description  登录
// @Accept       json
// @Produce      json
// @Param        Password  path  string  true  "密码"
// @Param        Username  path  string  true  "账号"
// @Router       /api/v1/login [post]
func (c *CtlLogin) Login(ctx *gin.Context) {
	var err error

	var req api.LoginReq
	if err = ctx.ShouldBindJSON(&req); err != nil {
		c.Fail(ctx, err.Error())
		return
	}

	var user *model.UserOutput
	if user, err = service.User().GetUserByUserName(ctx, req.Username); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.Fail(ctx, "用户名或密码不正确")
			return
		}
		c.Error(ctx, err)
		return
	}

	if req.Password != user.Password {
		c.Fail(ctx, "用户名或密码不正确")
		return
	}

	c.SuccessWithData(ctx, map[string]interface{}{
		"token": "这是token",
	})
	return
}
`
