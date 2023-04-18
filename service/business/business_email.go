package business

import (
	"gopkg.in/gomail.v2"
	logger "my-server-go/tool/log"
)

const smtpServer = "smtp.qq.com"

const port = 465

const from = "82008841@qq.com"

const pass = "zcdmupotbdhsbifd"

const to = "mister76@qq.com"

func SendEmail(body string) {
	message := gomail.NewMessage()
	message.SetHeader("From", from)
	message.SetHeader("To", to)
	message.SetHeader("Subject", "用户反馈")
	message.SetBody("text/plain", body)
	dialer := gomail.NewDialer(smtpServer, port, from, pass)
	err := dialer.DialAndSend(message)
	if err != nil {
		logger.Write("Error sending email:", err)
	}
	logger.Write("Email sent successfully!")
}
