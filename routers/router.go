package routers

import (
	"club/app/admin"
	"club/app/api"
	v1 "club/app/api/v1"
	"club/app/web"
	"club/middleware/jwt"
	"club/pkg/setting"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	r.GET("/auth", api.GetAuth)

	apiV1 := r.Group("/api/v1")
	apiV1.Use(jwt.JWT())
	//api
	{
		apiV1.GET("/articles", v1.GetArticles)
		//获取指定文章
		apiV1.GET("/articles/:id", v1.GetArticle)
		//新建文章
		apiV1.POST("/articles", v1.AddArticle)
		//更新指定文章
		apiV1.PUT("/articles/:id", v1.EditArticle)
		//删除指定文章
		apiV1.DELETE("/articles/:id", v1.DeleteArticle)
	}
	//web
	{
		r.GET("/", web.GetAll)
		r.GET("/player", web.Player)
	}
	//admin
	adminGroup := r.Group("/admin")
	{
		adminGroup.GET("/index", admin.Index)
		adminGroup.GET("/login", admin.Login)
		adminGroup.POST("/login", admin.DoLogin)
	}
	return r
}
