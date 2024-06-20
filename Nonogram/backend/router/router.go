package router

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	nono "gin/nono"
)

func InitRouter(r *gin.Engine) {
	r.Use(cors.Default())
	r.GET("/get/map", getMapHandler)
	r.GET("/get/string/:name", getStringHandler)
	r.POST("/post/map", postMapHandler)
	r.POST("/post/string", postStringHandler)

	nonoGroup := r.Group("/nono")
	{
		nonoGroup.GET("/fillGrid", nono.GetFillGrid)
		nonoGroup.POST("/check", nono.CheckFillGrid)
	}

	r.Run(":8686")
}

//-
//-
//-
//-
//-
//-
//-
//-
//-
//-
//-
//-
//-
//-
//-
//-
//-
//-

// 定义用于处理 POST 和 DELETE 请求的结构体
type MapRequest struct {
	Data map[string]interface{} `json:"data"`
}

type StringRequest struct {
	Value string `json:"value"`
}

func getMapHandler(c *gin.Context) {
	data := c.QueryMap("data")
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func getStringHandler(c *gin.Context) {
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{
		"value": name,
	})
}

func postMapHandler(c *gin.Context) {
	var req MapRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": req.Data,
	})
}

func postStringHandler(c *gin.Context) {
	var req StringRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"value": req.Value,
	})
}
