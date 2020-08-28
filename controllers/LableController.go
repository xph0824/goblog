package controllers

import (
	"github.com/gin-gonic/gin"
	"goModWork/databases"
	"goModWork/handler"
	"goModWork/models"
	"net/http"
)

type lable struct {

}

var Lable lable

func (*lable) CreateLable(c *gin.Context) {
	var lableForm databases.Lable
	bindErr := c.ShouldBind(&lableForm)
	if bindErr != nil {
		handler.Error(c, bindErr, http.StatusForbidden)
		return
	}

	createErr := models.CreateLable(lableForm)
	if createErr != nil{
		handler.Error(c, createErr, http.StatusBadRequest)
		return
	}
	handler.Success(c, "Create lable successfully!")
}

func (*lable) ListLable(c *gin.Context) {
	//查询标签
	lableList, err := models.ListLable()
	if err != nil {
		handler.Error(c, err, http.StatusInternalServerError)
	}
	c.JSON(200, gin.H{
		"lable_lis": lableList,
	})

}