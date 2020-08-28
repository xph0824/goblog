package controllers

import (
	"github.com/gin-gonic/gin"
	"goModWork/databases"
	"goModWork/handler"
	"goModWork/models"
	"net/http"
	"strconv"
)

type article struct {}

var Article article

func (*article) CreateArticle(c *gin.Context) {
	//fmt.Println(c.FullPath())
	var articleForm databases.Article
	bindErr := c.ShouldBind(&articleForm)
	if bindErr != nil {
		handler.Error(c, bindErr, http.StatusForbidden)
		return
	}
	//println("----")
	//print(articleForm.Uid);
	//println("title" + articleForm.Title);
	//println("content" + articleForm.Content);
	//println("pic" + articleForm.Pic);
	//println("----")
	createErr := models.CreateArticle(articleForm)
	if createErr != nil{
		handler.Error(c, createErr, http.StatusBadRequest)
		return
	}
	handler.Success(c, "Create article successfully!")
}

func (*article) ArticleList(c *gin.Context) {

	num := c.DefaultQuery("page", "1")
	keyword := c.DefaultQuery("keyword", "")
	page,err:= strconv.Atoi(num)
	if err!=nil{
		handler.Error(c, err, http.StatusInternalServerError)
	}
	lId := c.DefaultQuery("lid", "0")
	lableId,err:= strconv.Atoi(lId)

	//查询文章信息
	articleList, err := models.FindArticleLimit(page,lableId,keyword)
	if err != nil {
		handler.Error(c, err, http.StatusInternalServerError)
	}
	//获取记录数
	count := models.FinArticleCount(keyword);

	//查询标签
	lableList, err := models.ListLable()
	if err != nil {
		handler.Error(c, err, http.StatusInternalServerError)
	}

	c.HTML(http.StatusOK, "article.html", gin.H{
		"article_info": articleList,
		"count": count,
		"lable_list": lableList,
	})
}

