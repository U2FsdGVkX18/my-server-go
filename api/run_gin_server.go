package api

import "github.com/gin-gonic/gin"

func Run() {
	//创建一个默认gin web服务
	var ginServer = gin.Default()
	//调用函数创建路由
	WeChatAccess(ginServer)
	//启动服务
	err := ginServer.Run(":8000")
	if err != nil {
		return
	}
}
