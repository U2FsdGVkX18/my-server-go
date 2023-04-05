package business

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

func RainCityData(ginServer *gin.Engine) {
	ginServer.Use(CORSMiddleware())
	var businessGroup = ginServer.Group("/business")
	{
		businessGroup.GET("/getRainCity", func(context *gin.Context) {
			logger.Write("调用getRainCity接口")
			//从redis中获取城市数据
			data := business.GetRainCityForRedis()
			context.JSON(http.StatusOK, data)
			return
		})
		businessGroup.GET("/verifyCode/:code", func(context *gin.Context) {
			code := context.Param("code")
			if len(code) == 36 {
				logger.Write("调用verifyCode接口,接收到的code为:", code)
				status := business.CheckActivationCode(code)
				context.JSON(http.StatusOK, gin.H{"status": status})
				return
			} else {
				context.JSON(http.StatusOK, gin.H{"status": false})
				return
			}
		})
		businessGroup.GET("/verifyCodeExpire/:code", func(context *gin.Context) {
			code := context.Param("code")
			if len(code) == 36 {
				logger.Write("调用verifyCodeExpire接口,接收到的code为:", code)
				status := business.CheckActivationCodeIsExpire(code)
				context.JSON(http.StatusOK, gin.H{"status": status})
				return
			} else {
				context.JSON(http.StatusOK, gin.H{"status": true})
				return
			}
		})
	}
}
