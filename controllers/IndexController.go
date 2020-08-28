package controllers

import (
	"github.com/gin-gonic/gin"
	"goModWork/handler"
	"goModWork/models"
	"net/http"
)

func Index(c *gin.Context){
	//查询作者信息
	authInfo, err := models.FirstUser()
	if err != nil {
		handler.Error(c, err, http.StatusInternalServerError)
	}

	//查询文章信息
	articleInfo, err := models.FindArticleLimit(3, 0,"")
	c.HTML(http.StatusOK, "index.html", gin.H{
		"auth_info": authInfo,
		"article_info": articleInfo,
	})

}

