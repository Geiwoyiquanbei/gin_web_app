package routes

import (
	"net/http"
	"web_app/controllers"
	_ "web_app/docs" // 千万不要忘了导入把你上一步生成的docs
	"web_app/logger"
	"web_app/middleware"

	swaggerFiles "github.com/swaggo/files"

	"github.com/gin-gonic/gin"
	gs "github.com/swaggo/gin-swagger"
)

func SetUp() *gin.Engine {

	controllers.InitTrans("zh") //初始化gin框架内置的校验器使用的翻译器
	//gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	v1 := r.Group("/api/v1")
	r.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "successful")
	})
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	v1.POST("signup", controllers.SignUpHandler)
	v1.POST("login", controllers.LoginHandler)
	v1.Use(middleware.JWTAuthMiddleware())
	{
		v1.GET("/community", controllers.CommunityHandler)
		v1.GET("/community/:id", controllers.CommunityDetailHandler)
		v1.POST("/post", controllers.CreatePostHandler)
		v1.GET("/post/:id", controllers.GetPostDetailHandler)
		v1.GET("/posts/", controllers.GetPostListHandler)
		v1.GET("/posts2/", controllers.GetPostListHandler2)
		v1.POST("/vote", controllers.PostVoteController)
	}
	//r.NoRoute(func(context *gin.Context) {
	//	context.JSON(http.StatusOK, gin.H{
	//		"msg": "404",
	//	})
	//})
	r.GET("/ping", middleware.JWTAuthMiddleware(), func(context *gin.Context) {
		controllers.ResponseSuccess(context, "pong")
	})
	return r
}
