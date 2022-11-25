package template

var RouterTmp = `package router

import (
	"{{projectName}}/controller/user"
	"{{projectName}}/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func NewRouter() http.Handler {
	engine := gin.Default()
	setApiRouter(engine) // 设置API路由
	setSwagger(engine)
	return engine.Handler()
}

func setApiRouter(engine *gin.Engine) {
	api := engine.Group("/api")
	api.Use(middleware.Cors()) //跨域
	{
		v1 := api.Group("/v1")
		{
			login := user.CtlLogin{}
			v1.POST("/login", login.Login) //登录
		}
	}
}

func setSwagger(engine *gin.Engine) {
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}`
