package main

import (
	"gin/router"

	"github.com/gin-gonic/gin"
)

func main() {
	//1.创建路由
	r := gin.Default()
	router.InitRouter(r)

}
