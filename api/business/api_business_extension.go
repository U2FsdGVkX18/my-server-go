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

type Suggestion struct {
	Suggestion string `json:"suggestion"`
}

func ExtensionInterface(ginServer *gin.Engine) {
	ginServer.Use(CORSMiddleware())
	var businessGroup = ginServer.Group("/business")
	{
		businessGroup.GET("/getRainCity", func(context *gin.Context) {
			logger.Write("调用getRainCity接口")
			//从redis中获取城市数据
			data := business.GetRainCityForRedis()
			context.JSON(http.StatusOK, gin.H{"data": data})
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
				context.JSON(http.StatusOK, gin.H{"status": 2})
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
		businessGroup.POST("/suggestion", func(context *gin.Context) {
			var suggestion Suggestion
			err := context.BindJSON(&suggestion)
			if err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			if len(suggestion.Suggestion) > 900 {
				context.JSON(http.StatusBadRequest, gin.H{"error": "Suggestion length should not exceed 900 characters"})
				return
			}
			logger.Write("调用suggestion接口,接收到的suggestion为:", suggestion.Suggestion)
			business.SendEmail(suggestion.Suggestion)
			context.JSON(http.StatusOK, gin.H{"msg": "success"})
			return
		})
	}
}
