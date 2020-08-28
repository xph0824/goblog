package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goModWork/databases"
	"goModWork/handler"
	"goModWork/models"
	"net/http"
)

type user struct {}

var User user

// Create create a user
func (*user) CreateUser(c *gin.Context) {
	fmt.Println(c.FullPath())
	var userForm databases.User
	bindErr := c.ShouldBind(&userForm)
	if bindErr != nil {
		handler.Error(c, bindErr, http.StatusForbidden)
		return
	}
	println(userForm.Age);
	println(userForm.Name);
	createErr := models.CreateUser(userForm)
	if createErr != nil{
		handler.Error(c, createErr, http.StatusBadRequest)
		return
	}
	handler.Success(c, "Create user successfully!")
}

