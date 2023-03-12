package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"my-server-go/tool/log"
	"time"
)

func Connect() *gorm.DB {
	dsn := "lihongwei:mujin1110@tcp(125.91.35.185:3306)/my_server?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Write("failed to connect database:", err)
	}
	return db
}

type Scheduled struct {
	ID       uint   `gorm:"primarykey"`
	CronMean string `gorm:"type:varchar(255);comment:cron表达式含义"`
	Cron     string `gorm:"type:varchar(255);comment:cron表达式"`
	Type     string `gorm:"type:varchar(255)"`
}

type QywxUserLocation struct {
	ID           uint      `gorm:"primarykey"`
	UserName     string    `gorm:"type:varchar(255);comment:企业成员微信用户名"`
	UserLocation string    `gorm:"type:varchar(255);comment:企业成员定位地址纬度+经度"`
	CreatedAt    time.Time `gorm:"comment:创建时间"`
	UpdatedAt    time.Time `gorm:"comment:更新时间"`
}

type DoubanTvshowHighscore struct {
	ID          uint      `gorm:"primarykey"`
	Name        string    `gorm:"type:varchar(255);comment:tv名称"`
	Subtitle    string    `gorm:"type:varchar(255);comment:副标题"`
	Type        string    `gorm:"type:varchar(255);comment:类型"`
	Year        string    `gorm:"type:varchar(255);comment:年份"`
	Tags        string    `gorm:"type:varchar(255);comment:标签"`
	Score       string    `gorm:"type:varchar(255);comment:评分"`
	ScorePeople uint      `gorm:"comment:评分人数"`
	HotComment  string    `gorm:"type:varchar(255);comment:热评"`
	Details     string    `gorm:"type:varchar(255);comment:详情地址"`
	ImgUrl      string    `gorm:"type:varchar(255);comment:图片地址"`
	CreatedAt   time.Time `gorm:"comment:创建时间"`
}

type DoubanNewmovieRanking struct {
	ID          uint      `gorm:"primarykey"`
	Name        string    `gorm:"type:varchar(255);comment:电影名称"`
	Intro       string    `gorm:"type:varchar(255);comment:电影简介"`
	Score       string    `gorm:"type:varchar(255);comment:评分"`
	ScorePeople uint      `gorm:"comment:评分人数"`
	Details     string    `gorm:"type:varchar(255);comment:详情地址"`
	ImgUrl      string    `gorm:"type:varchar(255);comment:图片地址"`
	CreatedAt   time.Time `gorm:"comment:创建时间"`
}

type DoubanMovieTop250 struct {
	ID                   uint      `gorm:"primarykey"`
	Name                 string    `gorm:"type:varchar(255);comment:电影名称"`
	DirectorAndActors    string    `gorm:"type:varchar(255);comment:导演和演员"`
	YearAndRegionAndType string    `gorm:"type:varchar(255);comment:年份地区类型"`
	Score                string    `gorm:"type:varchar(255);comment:评分"`
	ScorePeople          uint      `gorm:"comment:评分人数"`
	Quote                string    `gorm:"type:varchar(255);comment:引用句"`
	Details              string    `gorm:"type:varchar(255);comment:详情地址"`
	ImgUrl               string    `gorm:"type:varchar(255);comment:图片地址"`
	CreatedAt            time.Time `gorm:"comment:创建时间"`
}

type DoubanMovieNowshowing struct {
	ID          uint      `gorm:"primarykey"`
	Name        string    `gorm:"type:varchar(255);comment:电影名称"`
	Score       string    `gorm:"type:varchar(255);comment:评分"`
	ScorePeople uint      `gorm:"comment:评分人数"`
	Release     string    `gorm:"type:varchar(255);comment:电影发布年份"`
	Duration    string    `gorm:"type:varchar(255);comment:片长"`
	Region      string    `gorm:"type:varchar(255);comment:地区"`
	Director    string    `gorm:"type:varchar(255);comment:导演"`
	Actors      string    `gorm:"type:varchar(255);comment:演员"`
	Details     string    `gorm:"type:varchar(255);comment:详情地址"`
	ImgUrl      string    `gorm:"type:varchar(255);comment:图片地址"`
	CreatedAt   time.Time `gorm:"comment:创建时间"`
}

type DoubanMovieComingsoon struct {
	ID          uint      `gorm:"primarykey"`
	Name        string    `gorm:"type:varchar(255);comment:电影名称"`
	Region      string    `gorm:"type:varchar(255);comment:地区"`
	ReleaseDate string    `gorm:"type:varchar(255);comment:上映日期"`
	Type        string    `gorm:"type:varchar(255);comment:类型"`
	WantToSee   string    `gorm:"type:varchar(255);comment:想看"`
	Details     string    `gorm:"type:varchar(255);comment:详情地址"`
	ImgUrl      string    `gorm:"type:varchar(255);comment:图片地址"`
	CreatedAt   time.Time `gorm:"comment:创建时间"`
}

type DoubanBookTop250 struct {
	ID                          uint      `gorm:"primarykey"`
	Name                        string    `gorm:"type:varchar(255);comment:书名"`
	AuthorPressPublicationPrice string    `gorm:"type:varchar(255);comment:作者出版社出版年定价"`
	Score                       string    `gorm:"type:varchar(255);comment:评分"`
	ScorePeople                 uint      `gorm:"comment:评分人数"`
	Quote                       string    `gorm:"type:varchar(255);comment:引用句"`
	Details                     string    `gorm:"type:varchar(255);comment:详情地址"`
	ImgUrl                      string    `gorm:"type:varchar(255);comment:图片地址"`
	CreatedAt                   time.Time `gorm:"comment:创建时间"`
}

type DoubanBookHottestPublish struct {
	ID         uint      `gorm:"primarykey"`
	Name       string    `gorm:"type:varchar(255);comment:书名"`
	Author     string    `gorm:"type:varchar(1000);comment:作者"`
	OrigAuthor string    `gorm:"type:varchar(255);comment:原创作者"`
	Translator string    `gorm:"type:varchar(255);comment:译者"`
	Summary    string    `gorm:"comment:摘要/简介"`
	Kinds      string    `gorm:"type:varchar(255);comment:种类"`
	WordCount  string    `gorm:"type:varchar(255);comment:字数"`
	FixedPrice float32   `gorm:"comment:原价"`
	SalesPrice float32   `gorm:"comment:现价"`
	Details    string    `gorm:"type:varchar(255);comment:详情地址"`
	ImgUrl     string    `gorm:"type:varchar(255);comment:图片地址"`
	CreatedAt  time.Time `gorm:"comment:创建时间"`
}

type DoubanBookHottestOriginal struct {
	ID         uint      `gorm:"primarykey"`
	Name       string    `gorm:"type:varchar(255);comment:书名"`
	Author     string    `gorm:"type:varchar(1000);comment:作者"`
	OrigAuthor string    `gorm:"type:varchar(255);comment:原创作者"`
	Translator string    `gorm:"type:varchar(255);comment:译者"`
	Summary    string    `gorm:"comment:摘要/简介"`
	Kinds      string    `gorm:"type:varchar(255);comment:种类"`
	WordCount  string    `gorm:"type:varchar(255);comment:字数"`
	FixedPrice float32   `gorm:"comment:原价"`
	SalesPrice float32   `gorm:"comment:现价"`
	Details    string    `gorm:"type:varchar(255);comment:详情地址"`
	ImgUrl     string    `gorm:"type:varchar(255);comment:图片地址"`
	CreatedAt  time.Time `gorm:"comment:创建时间"`
}

type DoubanBookHighsalesPublish struct {
	ID         uint      `gorm:"primarykey"`
	Name       string    `gorm:"type:varchar(255);comment:书名"`
	Author     string    `gorm:"type:varchar(1000);comment:作者"`
	OrigAuthor string    `gorm:"type:varchar(255);comment:原创作者"`
	Translator string    `gorm:"type:varchar(255);comment:译者"`
	Summary    string    `gorm:"comment:摘要/简介"`
	Kinds      string    `gorm:"type:varchar(255);comment:种类"`
	WordCount  string    `gorm:"type:varchar(255);comment:字数"`
	FixedPrice float32   `gorm:"comment:原价"`
	SalesPrice float32   `gorm:"comment:现价"`
	Details    string    `gorm:"type:varchar(255);comment:详情地址"`
	ImgUrl     string    `gorm:"type:varchar(255);comment:图片地址"`
	CreatedAt  time.Time `gorm:"comment:创建时间"`
}

func CreateTables() {
	db := Connect()
	//初始化表,当表不存在则创建表
	err := db.AutoMigrate(&Scheduled{}, &QywxUserLocation{}, &DoubanTvshowHighscore{}, &DoubanNewmovieRanking{},
		&DoubanMovieTop250{}, &DoubanMovieNowshowing{}, &DoubanMovieComingsoon{}, &DoubanBookTop250{}, &DoubanBookHottestPublish{},
		&DoubanBookHottestOriginal{}, &DoubanBookHighsalesPublish{})
	if err != nil {
		log.Write("db AutoMigrate err:", err)
	}
}
