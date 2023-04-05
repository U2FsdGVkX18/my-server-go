package business

import (
	"fmt"
	"github.com/google/uuid"
	"my-server-go/config/mysql"
	"time"
)

// TrialActivationCodeInsertDB 插入试用激活码到DB
func TrialActivationCodeInsertDB() {
	db := mysql.Connect()
	//生成50个
	codes := generateBatchActivationCodes(20)
	for _, code := range codes {
		fmt.Println(code)
		var businessTrialActivationCode = mysql.BusinessTrialActivationCode{
			Code:      code,
			StartDate: time.Now(),
			EndDate:   time.Now(),
			IsUsed:    false,
		}
		db.Create(&businessTrialActivationCode)
	}
}

// RegularActivationCodeInsertDB 插入正式激活码到DB
func RegularActivationCodeInsertDB() {
	db := mysql.Connect()
	//生成50个
	codes := generateBatchActivationCodes(20)
	for _, code := range codes {
		var businessRegularActivationCode = mysql.BusinessRegularActivationCode{
			Code:      code,
			StartDate: time.Now(),
			EndDate:   time.Now(),
			IsUsed:    false,
		}
		db.Create(&businessRegularActivationCode)
	}
}

// 批量生成激活码
func generateBatchActivationCodes(count int) []string {
	codes := make([]string, count)
	for i := 0; i < count; i++ {
		codes[i] = generateActivationCode()
	}
	return codes
}

// 生成激活码
func generateActivationCode() string {
	uniqueID, _ := uuid.NewRandom()
	return uniqueID.String()
}
