package log

import (
	"github.com/sirupsen/logrus"
	"os"
)

// 创建日志记录器
var logger = logrus.New()

func Write(any ...any) {
	//创建日志文件
	file, err := os.OpenFile("log/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logger.Info("file openFile err", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			logger.Info("file close err", err)
		}
	}(file)

	//设置日志输出到文件
	logger.SetOutput(file)

	//记录日志
	logger.Info(any)

}
