package api

import (
	"github.com/gin-gonic/gin"
	"my-server-go/service/business"
	logger "my-server-go/tool/log"
	"net/http"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func BusinessCityData(ginServer *gin.Engine) {
	ginServer.Use(CORSMiddleware())
	var businessGroup = ginServer.Group("/business")
	{
		businessGroup.GET("/getRainCity", func(context *gin.Context) {
			logger.Write("扩展程序调用getRainCity接口")
			//从redis中获取城市数据
			data := business.GetRainCityForRedis()
			context.JSON(http.StatusOK, data)
			return
		})
	}
}