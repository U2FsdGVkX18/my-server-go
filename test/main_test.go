package test

import (
	"fmt"
	"my-server-go/config/mysql"
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
	business2.TrialActivationCodeInsertDB()
	//正式
	//business2.RegularActivationCodeInsertDB()
}

func Test3(t *testing.T) {
	var cityIds []string
	mysql.DB.Model(mysql.BusinessCityList{}).Select("city_id").Scan(&cityIds)
	for _, v := range cityIds {
		fmt.Println(v)
	}
}
