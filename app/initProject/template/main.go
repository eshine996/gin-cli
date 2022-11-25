package template

var MainTmp = `package main

import (
	"context"
	"{{projectName}}/docs"
	g "{{projectName}}/global"
	"{{projectName}}/router"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "petstore.swagger.io"
	docs.SwaggerInfo.BasePath = "/v2"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	docs.SwaggerInfo.BasePath = "/api/v1"

	mux := router.NewRouter()
	server := &http.Server{
		Addr:           g.Config.GetString("server.host"),
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 启动服务
	go func() {
		if err := server.ListenAndServe(); err != nil {
			g.Log.Error(err)
			cancel()
			return
		}
	}()

	// 优雅退出
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)

	select {
	case sig := <-ch: //阻塞监听
		g.Log.Info(sig)
	case <-ctx.Done():
		cancel()
	}

	if err := server.Shutdown(ctx); err != nil {
		g.Log.Error(err)
	}

	g.Log.Info("exit")
}
`
