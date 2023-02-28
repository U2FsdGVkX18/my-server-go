package log

import (
	"log"
	"os"
)

func Write(any ...any) {
	file, err := os.OpenFile("server.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0664)
	if err != nil {
		return
	}
	logger := log.New(file, "log : ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println(any)
}
