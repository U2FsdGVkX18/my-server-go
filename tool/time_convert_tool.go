package tool

import "time"

// GetSystemCurrentDate 获取系统当前时间,并转为日期型日期,精确到日-(yyyy-MM-dd)
func GetSystemCurrentDate() string {
	now := time.Now()
	return now.Format("2006-01-02")
}

// GetDateTime13 传入国际格式时间, 获取日期时间型时间, 精确到秒
func GetDateTime13(time time.Time) string {
	return time.Format("2006-01-02 15:04:05")
}
