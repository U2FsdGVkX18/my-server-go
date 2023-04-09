package notion

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"io"
	"my-server-go/config/mysql"
	"my-server-go/invoke"
	"my-server-go/tool"
	logger "my-server-go/tool/log"
	"strconv"
)

// NotionApi
const notionBasicApi = "https://api.notion.com/v1"

// Notion机器人令牌
const notionBotKey = "secret_rvUNRXh5Imw1WmOudXxcOaF6yakOkXVSn3LuJg5uLrv"

// NotionApi版本
const notionVersion = "2022-06-28"

func SyncNewMovieRanking() {
	dataBaseId := "7ec1b96bd47f4c5c86fab914c24b3c73"
	pageUrl := notionBasicApi + "/pages"
	DeleteDataBaseData(dataBaseId)
	var doubanNewmovieRankings []mysql.DoubanNewmovieRanking
	mysql.DB.Select("id,name,intro,score,score_people,details,img_url,created_at").Find(&doubanNewmovieRankings)
	for _, doubanNewmovieRanking := range doubanNewmovieRankings {
		var id = ID{Number: doubanNewmovieRanking.ID}
		var name = Name{Title: GetTitle(doubanNewmovieRanking.Name)}
		var dataCreateTime = DataCreateTime{RichText: GetRichText(tool.GetDateTime13(doubanNewmovieRanking.CreatedAt))}
		var score = Score{GetRichText(doubanNewmovieRanking.Score)}
		var scorePeople = ScorePeople{RichText: GetRichText(strconv.FormatUint(uint64(doubanNewmovieRanking.ScorePeople), 10))}
		var intro = Intro{RichText: GetRichText(doubanNewmovieRanking.Intro)}
		var img = Img{
			Files: GetFiles(doubanNewmovieRanking.ImgUrl),
			Type:  "files",
		}
		var imgUrl = ImgUrl{
			Type: "url",
			URL:  doubanNewmovieRanking.ImgUrl,
		}
		var details = Details{
			Type: "url",
			URL:  doubanNewmovieRanking.Details,
		}
		var jsonRootBean = JsonRootBean{
			Parent: Parent{
				DatabaseID: dataBaseId,
				Type:       "database_id",
			},
			Properties: struct {
				Properties
				NewMovieRankingProperties
			}{
				Properties: Properties{
					ID:             id,
					Name:           name,
					Img:            img,
					ImgUrl:         imgUrl,
					Details:        details,
					DataCreateTime: dataCreateTime,
				},
				NewMovieRankingProperties: NewMovieRankingProperties{
					ScorePeople: scorePeople,
					Score:       score,
					Intro:       intro,
				},
			},
		}
		marshal, _ := json.Marshal(&jsonRootBean)
		invoke.SendPost(pageUrl, marshal, GetHeader())
	}
	logger.Write("SyncNewMovieRanking : 同步成功!")
}

func SyncMovieNowShowing() {
	dataBaseId := "3cf8045ba71b45099edc805fee5bcac6"
	pageUrl := notionBasicApi + "/pages"
	DeleteDataBaseData(dataBaseId)
	var doubanMovieNowshowings []mysql.DoubanMovieNowshowing
	mysql.DB.Select("id,name,score,score_people,`release`,duration,region,director,actors,details,img_url,created_at").Find(&doubanMovieNowshowings)
	for _, doubanMovieNowshowing := range doubanMovieNowshowings {
		var id = ID{Number: doubanMovieNowshowing.ID}
		var name = Name{Title: GetTitle(doubanMovieNowshowing.Name)}
		var dataCreateTime = DataCreateTime{RichText: GetRichText(tool.GetDateTime13(doubanMovieNowshowing.CreatedAt))}
		var score = Score{GetRichText(doubanMovieNowshowing.Score)}
		var scorePeople = ScorePeople{RichText: GetRichText(strconv.FormatUint(uint64(doubanMovieNowshowing.ScorePeople), 10))}
		var img = Img{
			Files: GetFiles(doubanMovieNowshowing.ImgUrl),
			Type:  "files",
		}
		var imgUrl = ImgUrl{
			Type: "url",
			URL:  doubanMovieNowshowing.ImgUrl,
		}
		var details = Details{
			Type: "url",
			URL:  doubanMovieNowshowing.Details,
		}
		var release = Release{RichText: GetRichText(doubanMovieNowshowing.Release)}
		var duration = Duration{RichText: GetRichText(doubanMovieNowshowing.Duration)}
		var region = Region{RichText: GetRichText(doubanMovieNowshowing.Region)}
		var director = Director{RichText: GetRichText(doubanMovieNowshowing.Director)}
		var actors = Actors{RichText: GetRichText(doubanMovieNowshowing.Actors)}
		var jsonRootBean = JsonRootBean{
			Parent: Parent{
				DatabaseID: dataBaseId,
				Type:       "database_id",
			},
			Properties: struct {
				Properties
				MovieNowShowingProperties
			}{
				Properties: Properties{
					ID:             id,
					Name:           name,
					Img:            img,
					ImgUrl:         imgUrl,
					Details:        details,
					DataCreateTime: dataCreateTime,
				},
				MovieNowShowingProperties: MovieNowShowingProperties{
					ScorePeople: scorePeople,
					Score:       score,
					Release:     release,
					Duration:    duration,
					Region:      region,
					Director:    director,
					Actors:      actors,
				},
			},
		}
		marshal, _ := json.Marshal(&jsonRootBean)
		invoke.SendPost(pageUrl, marshal, GetHeader())
	}
	logger.Write("SyncMovieNowShowing : 同步成功!")
}

func SyncMovieComingSoon() {
	dataBaseId := "2ba22a88135147ef9ed6d9df0bfc4243"
	pageUrl := notionBasicApi + "/pages"
	DeleteDataBaseData(dataBaseId)
	var doubanMovieComingsoons []mysql.DoubanMovieComingsoon
	mysql.DB.Select("id,name,region,release_date,type,want_to_see,details,img_url,created_at").Find(&doubanMovieComingsoons)
	for _, doubanMovieComingsoon := range doubanMovieComingsoons {
		var id = ID{Number: doubanMovieComingsoon.ID}
		var name = Name{Title: GetTitle(doubanMovieComingsoon.Name)}
		var dataCreateTime = DataCreateTime{RichText: GetRichText(tool.GetDateTime13(doubanMovieComingsoon.CreatedAt))}
		var img = Img{
			Files: GetFiles(doubanMovieComingsoon.ImgUrl),
			Type:  "files",
		}
		var imgUrl = ImgUrl{
			Type: "url",
			URL:  doubanMovieComingsoon.ImgUrl,
		}
		var details = Details{
			Type: "url",
			URL:  doubanMovieComingsoon.Details,
		}
		var region = Region{RichText: GetRichText(doubanMovieComingsoon.Region)}
		var releaseDate = ReleaseDate{RichText: GetRichText(doubanMovieComingsoon.ReleaseDate)}
		var typ = Type{RichText: GetRichText(doubanMovieComingsoon.Type)}
		var wantToSee = WantToSee{RichText: GetRichText(doubanMovieComingsoon.WantToSee)}

		var jsonRootBean = JsonRootBean{
			Parent: Parent{
				DatabaseID: dataBaseId,
				Type:       "database_id",
			},
			Properties: struct {
				Properties
				MovieComingSoonProperties
			}{
				Properties: Properties{
					ID:             id,
					Name:           name,
					Img:            img,
					ImgUrl:         imgUrl,
					Details:        details,
					DataCreateTime: dataCreateTime,
				},
				MovieComingSoonProperties: MovieComingSoonProperties{
					Region:      region,
					ReleaseDate: releaseDate,
					Type:        typ,
					WantToSee:   wantToSee,
				},
			},
		}
		marshal, _ := json.Marshal(&jsonRootBean)
		invoke.SendPost(pageUrl, marshal, GetHeader())
	}
	logger.Write("SyncMovieComingSoon : 同步成功!")
}

func SyncTop250MovieRanking() {
	dataBaseId := "27fd20d5025b48189b93b4a7f00aed5d"
	pageUrl := notionBasicApi + "/pages"
	DeleteDataBaseData(dataBaseId)
	var doubanMovieTop250s []mysql.DoubanMovieTop250
	mysql.DB.Select("id,name,score,score_people,director_and_actors,year_and_region_and_type,quote,details,img_url,created_at").Find(&doubanMovieTop250s)
	for _, doubanMovieTop250 := range doubanMovieTop250s {
		var id = ID{Number: doubanMovieTop250.ID}
		var name = Name{Title: GetTitle(doubanMovieTop250.Name)}
		var dataCreateTime = DataCreateTime{RichText: GetRichText(tool.GetDateTime13(doubanMovieTop250.CreatedAt))}
		var score = Score{GetRichText(doubanMovieTop250.Score)}
		var scorePeople = ScorePeople{RichText: GetRichText(strconv.FormatUint(uint64(doubanMovieTop250.ScorePeople), 10))}
		var img = Img{
			Files: GetFiles(doubanMovieTop250.ImgUrl),
			Type:  "files",
		}
		var imgUrl = ImgUrl{
			Type: "url",
			URL:  doubanMovieTop250.ImgUrl,
		}
		var details = Details{
			Type: "url",
			URL:  doubanMovieTop250.Details,
		}
		var directorAndActors = DirectorAndActors{RichText: GetRichText(doubanMovieTop250.DirectorAndActors)}
		var yearAndRegionAndType = YearAndRegionAndType{RichText: GetRichText(doubanMovieTop250.YearAndRegionAndType)}
		var quote = Quote{RichText: GetRichText(doubanMovieTop250.Quote)}

		var jsonRootBean = JsonRootBean{
			Parent: Parent{
				DatabaseID: dataBaseId,
				Type:       "database_id",
			},
			Properties: struct {
				Properties
				Top250MovieRankingProperties
			}{
				Properties: Properties{
					ID:             id,
					Name:           name,
					Img:            img,
					ImgUrl:         imgUrl,
					Details:        details,
					DataCreateTime: dataCreateTime,
				},
				Top250MovieRankingProperties: Top250MovieRankingProperties{
					ScorePeople:          scorePeople,
					Score:                score,
					DirectorAndActors:    directorAndActors,
					YearAndRegionAndType: yearAndRegionAndType,
					Quote:                quote,
				},
			},
		}
		marshal, _ := json.Marshal(&jsonRootBean)
		invoke.SendPost(pageUrl, marshal, GetHeader())
	}
	logger.Write("SyncTop250MovieRanking : 同步成功!")
}

func SyncHighScoreTVShowRanking() {
	dataBaseId := "950264e99ab744968a7ea046d23f8423"
	pageUrl := notionBasicApi + "/pages"
	DeleteDataBaseData(dataBaseId)
	var doubanTvshowHighscores []mysql.DoubanTvshowHighscore
	mysql.DB.Select("id,name,score,score_people,subtitle,type,year,tags,hot_comment,details,img_url,created_at").Find(&doubanTvshowHighscores)
	for _, doubanTvshowHighscore := range doubanTvshowHighscores {
		var id = ID{Number: doubanTvshowHighscore.ID}
		var name = Name{Title: GetTitle(doubanTvshowHighscore.Name)}
		var dataCreateTime = DataCreateTime{RichText: GetRichText(tool.GetDateTime13(doubanTvshowHighscore.CreatedAt))}
		var score = Score{GetRichText(doubanTvshowHighscore.Score)}
		var scorePeople = ScorePeople{RichText: GetRichText(strconv.FormatUint(uint64(doubanTvshowHighscore.ScorePeople), 10))}
		var img = Img{
			Files: GetFiles(doubanTvshowHighscore.ImgUrl),
			Type:  "files",
		}
		var imgUrl = ImgUrl{
			Type: "url",
			URL:  doubanTvshowHighscore.ImgUrl,
		}
		var details = Details{
			Type: "url",
			URL:  doubanTvshowHighscore.Details,
		}
		var subtitle = Subtitle{RichText: GetRichText(doubanTvshowHighscore.Subtitle)}
		var typ = Type{RichText: GetRichText(doubanTvshowHighscore.Type)}
		var year = Year{RichText: GetRichText(doubanTvshowHighscore.Year)}
		var tags = Tags{RichText: GetRichText(doubanTvshowHighscore.Tags)}
		var hotComment = HotComment{RichText: GetRichText(doubanTvshowHighscore.HotComment)}

		var jsonRootBean = JsonRootBean{
			Parent: Parent{
				DatabaseID: dataBaseId,
				Type:       "database_id",
			},
			Properties: struct {
				Properties
				HighScoreTVShowRankingProperties
			}{
				Properties: Properties{
					ID:             id,
					Name:           name,
					Img:            img,
					ImgUrl:         imgUrl,
					Details:        details,
					DataCreateTime: dataCreateTime,
				},
				HighScoreTVShowRankingProperties: HighScoreTVShowRankingProperties{
					ScorePeople: scorePeople,
					Score:       score,
					Subtitle:    subtitle,
					Type:        typ,
					Year:        year,
					Tags:        tags,
					HotComment:  hotComment,
				},
			},
		}
		marshal, _ := json.Marshal(&jsonRootBean)
		invoke.SendPost(pageUrl, marshal, GetHeader())
	}
	logger.Write("SyncHighScoreTVShowRanking : 同步成功!")
}

func SyncTop250BookRanking() {
	dataBaseId := "10f0c3cf30054e91bd4b52f7efb62c2b"
	pageUrl := notionBasicApi + "/pages"
	DeleteDataBaseData(dataBaseId)
	var doubanBookTop250s []mysql.DoubanBookTop250
	mysql.DB.Select("id,name,score,score_people,author_press_publication_price,quote,details,img_url,created_at").Find(&doubanBookTop250s)
	for _, doubanBookTop250 := range doubanBookTop250s {
		var id = ID{Number: doubanBookTop250.ID}
		var name = Name{Title: GetTitle(doubanBookTop250.Name)}
		var dataCreateTime = DataCreateTime{RichText: GetRichText(tool.GetDateTime13(doubanBookTop250.CreatedAt))}
		var score = Score{GetRichText(doubanBookTop250.Score)}
		var scorePeople = ScorePeople{RichText: GetRichText(strconv.FormatUint(uint64(doubanBookTop250.ScorePeople), 10))}
		var img = Img{
			Files: GetFiles(doubanBookTop250.ImgUrl),
			Type:  "files",
		}
		var imgUrl = ImgUrl{
			Type: "url",
			URL:  doubanBookTop250.ImgUrl,
		}
		var details = Details{
			Type: "url",
			URL:  doubanBookTop250.Details,
		}
		var authorPressPublicationPrice = AuthorPressPublicationPrice{RichText: GetRichText(doubanBookTop250.AuthorPressPublicationPrice)}
		var quote = Quote{RichText: GetRichText(doubanBookTop250.Quote)}

		var jsonRootBean = JsonRootBean{
			Parent: Parent{
				DatabaseID: dataBaseId,
				Type:       "database_id",
			},
			Properties: struct {
				Properties
				Top250BookRankingProperties
			}{
				Properties: Properties{
					ID:             id,
					Name:           name,
					Img:            img,
					ImgUrl:         imgUrl,
					Details:        details,
					DataCreateTime: dataCreateTime,
				},
				Top250BookRankingProperties: Top250BookRankingProperties{
					ScorePeople:                 scorePeople,
					Score:                       score,
					AuthorPressPublicationPrice: authorPressPublicationPrice,
					Quote:                       quote,
				},
			},
		}
		marshal, _ := json.Marshal(&jsonRootBean)
		invoke.SendPost(pageUrl, marshal, GetHeader())
	}
	logger.Write("SyncTop250BookRanking : 同步成功!")
}

func SyncHotTestPublishBookRanking() {
	dataBaseId := "def123f5c0da41c4a7e8b402175ef570"
	pageUrl := notionBasicApi + "/pages"
	DeleteDataBaseData(dataBaseId)
	var doubanBookHottestPublishs []mysql.DoubanBookHottestPublish
	mysql.DB.Select("id,name,author,orig_author,translator,summary,kinds,word_count,fixed_price,sales_price,details,img_url,created_at").Find(&doubanBookHottestPublishs)
	for _, doubanBookHottestPublish := range doubanBookHottestPublishs {
		var id = ID{Number: doubanBookHottestPublish.ID}
		var name = Name{Title: GetTitle(doubanBookHottestPublish.Name)}
		var dataCreateTime = DataCreateTime{RichText: GetRichText(tool.GetDateTime13(doubanBookHottestPublish.CreatedAt))}
		var img = Img{
			Files: GetFiles(doubanBookHottestPublish.ImgUrl),
			Type:  "files",
		}
		var imgUrl = ImgUrl{
			Type: "url",
			URL:  doubanBookHottestPublish.ImgUrl,
		}
		var details = Details{
			Type: "url",
			URL:  doubanBookHottestPublish.Details,
		}
		var author = Author{RichText: GetRichText(doubanBookHottestPublish.Author)}
		var origAuthor = OrigAuthor{RichText: GetRichText(doubanBookHottestPublish.OrigAuthor)}
		var translator = Translator{RichText: GetRichText(doubanBookHottestPublish.Translator)}
		var summary = Summary{RichText: GetRichText(doubanBookHottestPublish.Summary)}
		var kinds = Kinds{RichText: GetRichText(doubanBookHottestPublish.Kinds)}
		var wordCount = WordCount{RichText: GetRichText(doubanBookHottestPublish.WordCount)}
		var fixedPrice = FixedPrice{RichText: GetRichText(strconv.FormatFloat(float64(doubanBookHottestPublish.FixedPrice), 'f', -2, 32))}
		var salesPrice = SalesPrice{RichText: GetRichText(strconv.FormatFloat(float64(doubanBookHottestPublish.SalesPrice), 'f', -2, 32))}

		var jsonRootBean = JsonRootBean{
			Parent: Parent{
				DatabaseID: dataBaseId,
				Type:       "database_id",
			},
			Properties: struct {
				Properties
				BookProperties
			}{
				Properties: Properties{
					ID:             id,
					Name:           name,
					Img:            img,
					ImgUrl:         imgUrl,
					Details:        details,
					DataCreateTime: dataCreateTime,
				},
				BookProperties: BookProperties{
					Author:     author,
					OrigAuthor: origAuthor,
					Translator: translator,
					Summary:    summary,
					Kinds:      kinds,
					WordCount:  wordCount,
					FixedPrice: fixedPrice,
					SalesPrice: salesPrice,
				},
			},
		}
		marshal, _ := json.Marshal(&jsonRootBean)
		invoke.SendPost(pageUrl, marshal, GetHeader())
	}
	logger.Write("SyncHotTestPublishBookRanking : 同步成功!")
}

func SyncHighSalesPublishBookRanking() {
	dataBaseId := "ffd04fb305614488a31c6b460bd0e2df"
	pageUrl := notionBasicApi + "/pages"
	DeleteDataBaseData(dataBaseId)
	var doubanBookHighsalesPublishs []mysql.DoubanBookHighsalesPublish
	mysql.DB.Select("id,name,author,orig_author,translator,summary,kinds,word_count,fixed_price,sales_price,details,img_url,created_at").Find(&doubanBookHighsalesPublishs)
	for _, doubanBookHighsalesPublish := range doubanBookHighsalesPublishs {
		var id = ID{Number: doubanBookHighsalesPublish.ID}
		var name = Name{Title: GetTitle(doubanBookHighsalesPublish.Name)}
		var dataCreateTime = DataCreateTime{RichText: GetRichText(tool.GetDateTime13(doubanBookHighsalesPublish.CreatedAt))}
		var img = Img{
			Files: GetFiles(doubanBookHighsalesPublish.ImgUrl),
			Type:  "files",
		}
		var imgUrl = ImgUrl{
			Type: "url",
			URL:  doubanBookHighsalesPublish.ImgUrl,
		}
		var details = Details{
			Type: "url",
			URL:  doubanBookHighsalesPublish.Details,
		}
		var author = Author{RichText: GetRichText(doubanBookHighsalesPublish.Author)}
		var origAuthor = OrigAuthor{RichText: GetRichText(doubanBookHighsalesPublish.OrigAuthor)}
		var translator = Translator{RichText: GetRichText(doubanBookHighsalesPublish.Translator)}
		var summary = Summary{RichText: GetRichText(doubanBookHighsalesPublish.Summary)}
		var kinds = Kinds{RichText: GetRichText(doubanBookHighsalesPublish.Kinds)}
		var wordCount = WordCount{RichText: GetRichText(doubanBookHighsalesPublish.WordCount)}
		var fixedPrice = FixedPrice{RichText: GetRichText(strconv.FormatFloat(float64(doubanBookHighsalesPublish.FixedPrice), 'f', -2, 32))}
		var salesPrice = SalesPrice{RichText: GetRichText(strconv.FormatFloat(float64(doubanBookHighsalesPublish.SalesPrice), 'f', -2, 32))}

		var jsonRootBean = JsonRootBean{
			Parent: Parent{
				DatabaseID: dataBaseId,
				Type:       "database_id",
			},
			Properties: struct {
				Properties
				BookProperties
			}{
				Properties: Properties{
					ID:             id,
					Name:           name,
					Img:            img,
					ImgUrl:         imgUrl,
					Details:        details,
					DataCreateTime: dataCreateTime,
				},
				BookProperties: BookProperties{
					Author:     author,
					OrigAuthor: origAuthor,
					Translator: translator,
					Summary:    summary,
					Kinds:      kinds,
					WordCount:  wordCount,
					FixedPrice: fixedPrice,
					SalesPrice: salesPrice,
				},
			},
		}
		marshal, _ := json.Marshal(&jsonRootBean)
		invoke.SendPost(pageUrl, marshal, GetHeader())
	}
	logger.Write("SyncHighSalesPublishBookRanking : 同步成功!")
}

func SyncHotTestOriginalBookRanking() {
	dataBaseId := "c3b034d6598645d0bc87e1d12c4aecd7"
	pageUrl := notionBasicApi + "/pages"
	DeleteDataBaseData(dataBaseId)
	var doubanBookHottestOriginals []mysql.DoubanBookHottestOriginal
	mysql.DB.Select("id,name,author,orig_author,translator,summary,kinds,word_count,fixed_price,sales_price,details,img_url,created_at").Find(&doubanBookHottestOriginals)
	for _, doubanBookHottestOriginal := range doubanBookHottestOriginals {
		var id = ID{Number: doubanBookHottestOriginal.ID}
		var name = Name{Title: GetTitle(doubanBookHottestOriginal.Name)}
		var dataCreateTime = DataCreateTime{RichText: GetRichText(tool.GetDateTime13(doubanBookHottestOriginal.CreatedAt))}
		var img = Img{
			Files: GetFiles(doubanBookHottestOriginal.ImgUrl),
			Type:  "files",
		}
		var imgUrl = ImgUrl{
			Type: "url",
			URL:  doubanBookHottestOriginal.ImgUrl,
		}
		var details = Details{
			Type: "url",
			URL:  doubanBookHottestOriginal.Details,
		}
		var author = Author{RichText: GetRichText(doubanBookHottestOriginal.Author)}
		var origAuthor = OrigAuthor{RichText: GetRichText(doubanBookHottestOriginal.OrigAuthor)}
		var translator = Translator{RichText: GetRichText(doubanBookHottestOriginal.Translator)}
		var summary = Summary{RichText: GetRichText(doubanBookHottestOriginal.Summary)}
		var kinds = Kinds{RichText: GetRichText(doubanBookHottestOriginal.Kinds)}
		var wordCount = WordCount{RichText: GetRichText(doubanBookHottestOriginal.WordCount)}
		var fixedPrice = FixedPrice{RichText: GetRichText(strconv.FormatFloat(float64(doubanBookHottestOriginal.FixedPrice), 'f', -2, 32))}
		var salesPrice = SalesPrice{RichText: GetRichText(strconv.FormatFloat(float64(doubanBookHottestOriginal.SalesPrice), 'f', -2, 32))}

		var jsonRootBean = JsonRootBean{
			Parent: Parent{
				DatabaseID: dataBaseId,
				Type:       "database_id",
			},
			Properties: struct {
				Properties
				BookProperties
			}{
				Properties: Properties{
					ID:             id,
					Name:           name,
					Img:            img,
					ImgUrl:         imgUrl,
					Details:        details,
					DataCreateTime: dataCreateTime,
				},
				BookProperties: BookProperties{
					Author:     author,
					OrigAuthor: origAuthor,
					Translator: translator,
					Summary:    summary,
					Kinds:      kinds,
					WordCount:  wordCount,
					FixedPrice: fixedPrice,
					SalesPrice: salesPrice,
				},
			},
		}
		marshal, _ := json.Marshal(&jsonRootBean)
		invoke.SendPost(pageUrl, marshal, GetHeader())
	}
	logger.Write("SyncHotTestOriginalBookRanking : 同步成功!")
}

func GetFiles(url string) []Files {
	var external = External{URL: url}
	var files = Files{
		External: external,
		Name:     "url",
		Type:     "external",
	}
	var fileList []Files
	return append(fileList, files)
}

func GetRichText(content string) []RichText {
	var text = Text{Content: content}
	var richText = RichText{
		Text: text,
		Type: "text",
	}
	var richTextList []RichText
	return append(richTextList, richText)
}

func GetTitle(content string) []Title {
	var titleList []Title
	var text = Text{Content: content}
	var title = Title{
		Type: "text",
		Text: text,
	}
	return append(titleList, title)
}

func DeleteDataBaseData(dataBaseId string) {
	dataBaseQueryUrl := notionBasicApi + "/databases/" + dataBaseId + "/query"
	for {
		resp := invoke.SendPost(dataBaseQueryUrl, nil, GetHeader())
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		results := gjson.Get(string(body), "results").Array()
		if len(results) != 0 {
			for _, v := range results {
				id := v.Get("id").String()
				invoke.SendDelete(notionBasicApi+"/blocks/"+id, nil, GetHeader())
			}
		} else {
			break
		}
	}
}

func GetHeader() map[string]string {
	var headers = make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Notion-Version"] = notionVersion
	headers["Authorization"] = "Bearer " + notionBotKey
	return headers
}
