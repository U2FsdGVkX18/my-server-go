package business

import (
	"my-server-go/config/mysql"
	"time"
)

func CheckActivationCodeIsExpire(code string) bool {
	//查询试用激活码是否过期
	//试用激活码
	var businessTrialActivationCode mysql.BusinessTrialActivationCode
	result1 := mysql.DB.Where("code = ?", code).First(&businessTrialActivationCode)
	if result1.RowsAffected == 1 {
		//判断激活码是否过期
		if time.Now().After(businessTrialActivationCode.EndDate) {
			//过期
			return true
		} else {
			//未过期
			return false
		}
	}

	//正式激活码
	var businessRegularActivationCode mysql.BusinessRegularActivationCode
	result2 := mysql.DB.Where("code = ?", code).First(&businessRegularActivationCode)
	if result2.RowsAffected == 1 {
		//判断激活码是否过期
		if time.Now().After(businessRegularActivationCode.EndDate) {
			//过期
			return true
		} else {
			//未过期
			return false
		}
	}

	//过期
	return true
}
