package template

var MiddlewareCorsTmp = `package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

//Cors 跨域中间件
func Cors() gin.HandlerFunc {
	c := cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:    []string{"Content-Type", "Access-Token", "Authorization", "Token"},
		MaxAge:          6 * time.Hour,
	}
	return cors.New(c)
}`

var MiddlewareAuthTmp = `package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

//Auth 认证中间件
func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		fmt.Println(token)
	}
}`
