package main
import (
	"github.com/gin-gonic/gin"
	"net/http"
)
func main() {
	//创建默认的Gin引擎
	r := gin.Default()
	//定义一个GET请求的路由
	r.GET("/hello", func(c *gin.Context)){
		c.JSON(http.StatusOK, gin.H{
			"message": "hello Gin!",
	})
	r.Run(":8080")
}