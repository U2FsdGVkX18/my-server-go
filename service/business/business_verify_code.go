package business

import "my-server-go/config/mysql"

func CheckActivationCode(code string) bool {
	db := mysql.Connect()
	//查询试用激活码是否存在
	//试用激活码
	var businessTrialActivationCode mysql.BusinessTrialActivationCode
	result1 := db.Where("code = ?", code).First(&businessTrialActivationCode)
	if result1.RowsAffected == 1 {
		if !businessTrialActivationCode.IsUsed {
			//标记为已使用
			businessTrialActivationCode.IsUsed = true
			//更新数据库
			db.Save(&businessTrialActivationCode)
			return true
		} else {
			return false
		}
	}

	//正式激活码
	var businessRegularActivationCode mysql.BusinessRegularActivationCode
	result2 := db.Where("code = ?", code).First(&businessRegularActivationCode)
	if result2.RowsAffected == 1 {
		if !businessRegularActivationCode.IsUsed {
			//标记为已使用
			businessRegularActivationCode.IsUsed = true
			//更新数据库
			db.Save(&businessRegularActivationCode)
			return true
		} else {
			return false
		}
	}
	//如果都不满足则返回false
	return false
}
