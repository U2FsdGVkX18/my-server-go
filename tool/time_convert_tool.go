package tool

import "time"

// GetSystemCurrentDate 获取系统当前时间,并转为日期型日期,精确到日-(yyyy-MM-dd)
func GetSystemCurrentDate() string {
	now := time.Now()
	return now.Format("2006-01-02")
}
