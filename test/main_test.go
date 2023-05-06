package test

import (
	"fmt"
	"my-server-go/config/mysql"
	"my-server-go/service/business"
	business2 "my-server-go/tool/business"
	"testing"
)

func Test1(t *testing.T) {
	//mysql.CreateTables()

	//status := business2.CheckActivationCodeIsExpire("33e032e2-e498-4db6-9ed8-a012613f884e")
	//fmt.Println(status)
	type Status int
	const (
		// Active 已激活
		Active Status = iota
		// Inactive 未激活
		Inactive
		// Invalid 无效
		Invalid
	)
	fmt.Println(Active, Inactive, Invalid)
}

// 生成激活码
func TestGenCode(t *testing.T) {
	//试用
	business2.TrialActivationCodeInsertDB(1)
	//正式
	//business2.RegularActivationCodeInsertDB(100)
}

func Test3(t *testing.T) {
	var cityIds []string
	mysql.DB.Model(mysql.BusinessCityList{}).Select("city_id").Scan(&cityIds)
	for _, v := range cityIds {
		fmt.Println(v)
	}
}

func Test4(t *testing.T) {
	business.SendEmail("测试")
}

func Test5(t *testing.T) {
	code := business.CheckActivationCode("527aa9c5-7efd-40b3-b3c8-ea30c96e13d9")
	fmt.Println(code)
}
