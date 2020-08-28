package main

import (
	"github.com/gin-gonic/gin"
	"goModWork/databases"
	"goModWork/routers"
)



func main() {
	//初始化数据库
	databases.InitDB()
	defer databases.CloserDB()
	router := gin.Default()
	//加载路由
	routers.Router(router)
	//启动服务
	router.Run(":8080")

}