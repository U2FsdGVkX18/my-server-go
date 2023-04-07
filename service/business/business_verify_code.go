package business

import (
	"my-server-go/config/mysql"
	"time"
)

type Status int

// 定义激活码状态枚举
const (
	// Active 已激活
	Active Status = iota
	// Inactive 未激活
	Inactive
	// Invalid 无效
	Invalid
)

func CheckActivationCode(code string) Status {
	db := mysql.Connect()
	//查询试用激活码是否存在
	//试用激活码
	var businessTrialActivationCode mysql.BusinessTrialActivationCode
	result1 := db.Where("code = ?", code).First(&businessTrialActivationCode)
	if result1.RowsAffected == 1 {
		//判断激活码是否被使用
		if !businessTrialActivationCode.IsUsed {
			//未使用
			//标记为已使用
			businessTrialActivationCode.IsUsed = true
			//插入当前时间
			businessTrialActivationCode.StartDate = time.Now()
			//插入过期时间
			businessTrialActivationCode.EndDate = time.Now().AddDate(0, 0, 1)
			//更新数据库
			db.Save(&businessTrialActivationCode)
			return Inactive
		} else {
			//已使用
			return Active
		}
	}

	//正式激活码
	var businessRegularActivationCode mysql.BusinessRegularActivationCode
	result2 := db.Where("code = ?", code).First(&businessRegularActivationCode)
	if result2.RowsAffected == 1 {
		if !businessRegularActivationCode.IsUsed {
			//标记为已使用
			businessRegularActivationCode.IsUsed = true
			//插入当前时间
			businessRegularActivationCode.StartDate = time.Now()
			//插入过期时间
			businessRegularActivationCode.EndDate = time.Now().AddDate(50, 0, 0)
			//更新数据库
			db.Save(&businessRegularActivationCode)
			return Inactive
		} else {
			return Active
		}
	}
	//如果都不满足则返回false
	return Invalid
}
