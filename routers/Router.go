package routers

import (
	"github.com/gin-gonic/gin"
	"goModWork/controllers"
)

func Router(Router *gin.Engine) {

	// 配置静态资源目录
	Router.Static("/static", "./static")
	// 配置模板目录
	Router.LoadHTMLGlob("views/**/*")

	Router.GET("/pong", func(c *gin.Context) {
		c.String(200, "pong")
	})

	//首页
	Router.Any("/index", controllers.Index)

	//文章相关
	Router.GET("/article", controllers.Article.ArticleList)
	Router.POST("/create_article", controllers.Article.CreateArticle)

	//用户相关
	Router.POST("/create_user", controllers.User.CreateUser)

	//标签相关操作
	Router.POST("/create_lable", controllers.Lable.CreateLable)

}
