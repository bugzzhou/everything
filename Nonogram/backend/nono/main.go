package nono

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var res = Nonogram{}

func init() {
	res.Gen("simple")
}

func GetFillGrid(c *gin.Context) {
	c.JSON(http.StatusOK, res)
}

func CheckFillGrid(c *gin.Context) {
	var req Nonogram

	// 绑定并解析 JSON 请求体
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	status := req.Check()
	if status {
		res.Gen("simple")
	}
	c.JSON(http.StatusOK, gin.H{
		"status": status,
	})
}
