package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
	"encoding/json"
)

func main() {
	//创建一个服务
	ginServer := gin.Default()
	ginServer.Use(favicon.New("favicon.ico"))
	//加载静态页面
	ginServer.LoadHTMLGlob("templates/*")
	//加载资源文件
	ginServer.Static("/static", "./static")
	//相应一个页面给前端
	ginServer.GET("/index", func(context *gin.Context){
		//context.JSON() json数据
		context.HTML(200, "index.html", gin.H{
			"msg" : "这是一个go后台传递来的数据",
		})
		})
		//usl? userid=xxx&username=index
		ginServer.GET("/user/info", func(context *gin.Context) {
		userid := context.Query("userid")
		username := context.Query("username")
		context.JSON(200, gin.H{
			"userid":  userid,
			"username": username,
		})
	})
		// /user/info/1/index
		ginServer.GET("/user/info/:userid/:username", func(context *gin.Context){
		userid := context.Param("userid")
		username := context.Param("username")
			context.JSON(200, gin.H{
				"userid" : userid,
				"username" : username,
		})
	})

	//前端给后端传递json
	ginServer.POST("/json", func(context *gin.Context) {
    // request body
    data, _ := context.GetRawData()

    var m map[string]interface{}
    if err := json.Unmarshal(data, &m); err != nil {
        context.JSON(200, gin.H{"error": "Invalid JSON"})
        return
    }
    context.JSON(200, m) // 确保 m 是一个 object 类型的数据
})
	//服务器端口
	ginServer.Run(":8082")

}