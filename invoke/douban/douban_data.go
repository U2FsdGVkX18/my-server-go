package douban

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/tidwall/gjson"
	"io"
	"my-server-go/config/mysql"
	"my-server-go/invoke"
	logger "my-server-go/tool/log"
	"regexp"
	"strconv"
	"strings"
)

func GetNewMovieRanking() {
	url := "https://movie.douban.com/chart"
	resp := invoke.SendGet(url, nil, GetHeader())
	//defer关闭io流
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		logger.Write("GetNewMovieRanking goquery.NewDocumentFromReader异常:", err)
	}
	doc.Find("div.article").Eq(0).Find("table").Each(func(i int, selection *goquery.Selection) {
		td0 := selection.Find("td").Eq(0)
		details, _ := td0.Find("a").Attr("href")
		imgUrl, _ := td0.Find("img").Attr("src")
		td1 := selection.Find("td").Eq(1)
		name := strings.Fields(strings.TrimSpace(td1.Find("a").Text()))[0]
		intro := strings.Fields(strings.TrimSpace(td1.Find("p.pl").Eq(0).Text()))[0]
		score := td1.Find("span.rating_nums").Eq(0).Text()
		scorePeople := td1.Find("span.pl").Eq(0).Text()
		compile, _ := regexp.Compile(`[^0-9]`)
		matchString := strings.TrimSpace(compile.ReplaceAllString(scorePeople, ""))
		//字符串转为uint类型,10:10进制,32:uint32
		parseUint, _ := strconv.ParseUint(matchString, 10, 32)
		var douBanDataMovieRanking = mysql.DoubanNewmovieRanking{
			Details:     details,
			ImgUrl:      imgUrl,
			Name:        name,
			Intro:       intro,
			Score:       score,
			ScorePeople: uint(parseUint),
		}
		fmt.Println(douBanDataMovieRanking)
	})
	logger.Write("getNewMovieRanking 豆瓣新片电影排行数据爬取完成")
}

func GetMovieNowShowing() {
	url := "https://movie.douban.com/cinema/nowplaying/hangzhou/"
	resp := invoke.SendGet(url, nil, GetHeader())
	//defer关闭io流
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		logger.Write("GetMovieNowShowing goquery.NewDocumentFromReader异常:", err)
	}
	doc.Find("div#nowplaying").Find("div.mod-bd").Eq(0).Find("li.list-item").Each(func(i int, selection *goquery.Selection) {
		name, _ := selection.Attr("data-title")
		score, _ := selection.Attr("data-score")
		release, _ := selection.Attr("data-release")
		duration, _ := selection.Attr("data-duration")
		region, _ := selection.Attr("data-region")
		director, _ := selection.Attr("data-director")
		actors, _ := selection.Attr("data-actors")
		scorePeople, _ := selection.Attr("data-votecount")
		parseUint, _ := strconv.ParseUint(scorePeople, 10, 32)
		details, _ := selection.Find(".poster").Eq(0).Find("a").Attr("href")
		imgUrl, _ := selection.Find(".poster").Eq(0).Find("img").Attr("src")
		var doubanMovieNowshowing = mysql.DoubanMovieNowshowing{
			Name:        name,
			Score:       score,
			Release:     release,
			Duration:    duration,
			Region:      region,
			Director:    director,
			Actors:      actors,
			ScorePeople: uint(parseUint),
			Details:     details,
			ImgUrl:      imgUrl,
		}
		fmt.Println(doubanMovieNowshowing)
	})
	logger.Write("GetMovieNowShowing 豆瓣电影正在上映数据爬取完成")
}

func GetMovieComingSoon() {
	url := "https://movie.douban.com/cinema/later/hangzhou/"
	resp := invoke.SendGet(url, nil, GetHeader())

	//defer关闭io流
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		logger.Write("GetMovieComingSoon goquery.NewDocumentFromReader异常:", err)
	}
	doc.Find("div#showing-soon").Find("div[class*='item mod']").Each(func(i int, selection *goquery.Selection) {
		details, _ := selection.Find("a[class='thumb']").Eq(0).Attr("href")
		imgUrl, _ := selection.Find("a[class='thumb']").Eq(0).Find("img").Attr("src")
		name := selection.Find("h3").Eq(0).Find("a").Eq(0).Text()
		releaseDate := selection.Find("ul").Eq(0).Find("li").Eq(0).Text()
		Type := selection.Find("ul").Eq(0).Find("li").Eq(1).Text()
		region := selection.Find("ul").Eq(0).Find("li").Eq(2).Text()
		wantToSee := selection.Find("ul").Eq(0).Find("li").Eq(3).Find("span").Eq(0).Text()

		var doubanMovieComingsoon = mysql.DoubanMovieComingsoon{
			Details:     details,
			ImgUrl:      imgUrl,
			Name:        name,
			ReleaseDate: releaseDate,
			Type:        Type,
			Region:      region,
			WantToSee:   wantToSee,
		}
		fmt.Println(doubanMovieComingsoon)
	})
	logger.Write("GetMovieComingSoon 豆瓣电影即将上映数据爬取完成")
}

func GetTop250MovieRanking() {
	db := mysql.Connect()
	for i := 0; i <= 225; i += 25 {
		url := "https://movie.douban.com/top250?start=" + strconv.Itoa(i) + "&filter="
		resp := invoke.SendGet(url, nil, GetHeader())

		//defer关闭io流
		defer func(Body io.ReadCloser) {
			_ = Body.Close()
		}(resp.Body)

		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			logger.Write("GetTop250MovieRanking goquery.NewDocumentFromReader异常:", err)
		}
		doc.Find("div#content").Find("ol[class='grid_view']").Eq(0).Find("li").Each(func(i int, selection *goquery.Selection) {
			details, _ := selection.Find("div[class='pic']").Eq(0).Find("a").Attr("href")
			imgUrl, _ := selection.Find("div[class='pic']").Eq(0).Find("a").Eq(0).Find("img").Attr("src")
			name := selection.Find("div[class='hd']").Eq(0).Find("a").Eq(0).Find("span").Eq(0).Text()
			originalStr := selection.Find("div[class='bd']").Eq(0).Find("p").Eq(0).Text()
			all := strings.ReplaceAll(originalStr, "&nbsp;", "")
			directorAndActors := strings.Split(strings.ReplaceAll(all, " ", ""), "\n")[1]
			yearAndRegionAndType := strings.Split(strings.ReplaceAll(all, " ", ""), "\n")[2]
			score := selection.Find("div[class='bd']").Eq(0).Find("div").Eq(0).Find("span").Eq(1).Text()
			scorePeople := selection.Find("div[class='bd']").Eq(0).Find("div").Eq(0).Find("span").Eq(3).Text()
			compile, _ := regexp.Compile(`[^0-9]`)
			matchString := strings.TrimSpace(compile.ReplaceAllString(scorePeople, ""))
			parseUint, _ := strconv.ParseUint(matchString, 10, 32)

			var quote string
			fmt.Println(selection.Find("div[class='bd']").Eq(0).Find("p").Length())
			if selection.Find("div[class='bd']").Eq(0).Find("p").Length() < 2 {
				quote = ""
			} else {
				quote = selection.Find("div[class='bd']").Eq(0).Find("p").Eq(1).Find("span").Text()
			}
			var doubanMovieTop250 = mysql.DoubanMovieTop250{
				Details:              details,
				ImgUrl:               imgUrl,
				Name:                 name,
				DirectorAndActors:    directorAndActors,
				YearAndRegionAndType: yearAndRegionAndType,
				Score:                score,
				ScorePeople:          uint(parseUint),
				Quote:                quote,
			}
			db.Create(&doubanMovieTop250)
			fmt.Println(doubanMovieTop250)
		})
		logger.Write("GetTop250MovieRanking 豆瓣TOP250第" + strconv.Itoa(i/25+1) + "页电影数据爬取完成")
	}
}

func GetHighScoreTVShowRanking() {
	db := mysql.Connect()
	for i := 0; i <= 180; i += 20 {
		url := "https://m.douban.com/rexxar/api/v2/tv/recommend?refresh=0&start=" + strconv.Itoa(i) + "&count=20&selected_categories=%7B%7D&uncollect=false&sort=S&tags="
		var headers = make(map[string]string)
		headers["User-Agent"] = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36"
		headers["Referer"] = "https://movie.douban.com/tv/"
		resp := invoke.SendGet(url, nil, headers)

		//defer关闭io流
		defer func(Body io.ReadCloser) {
			_ = Body.Close()
		}(resp.Body)
		body, _ := io.ReadAll(resp.Body)
		items := gjson.Get(string(body), "items").Array()
		for _, item := range items {
			name := item.Get("title").String()
			subtitl := item.Get("card_subtitle").String()
			Type := item.Get("type").String()
			year := item.Get("year").String()
			tags := ""
			for _, tag := range item.Get("tags").Array() {
				tags += tag.Get("name").String() + "/"
			}
			score := item.Get("rating.value").String()
			scorePeople := item.Get("rating.count").Int()
			hotComment := item.Get("comment.comment").String()
			id := item.Get("id").String()
			imgUrl := item.Get("pic.large").String()
			var doubanTvshowHighscore = mysql.DoubanTvshowHighscore{
				Name:        name,
				Subtitle:    subtitl,
				Type:        Type,
				Year:        year,
				Tags:        tags,
				Score:       score,
				ScorePeople: uint(scorePeople),
				HotComment:  hotComment,
				Details:     "https://www.douban.com/doubanapp/dispatch?uri=/tv/" + id,
				ImgUrl:      imgUrl,
			}
			db.Create(&doubanTvshowHighscore)
		}
		//logger.Write("GetHighScoreTVShowRanking 豆瓣高分电视剧第" + strconv.Itoa(i/25+1) + "页数据爬取完成")
	}
}

func GetTop250BookRanking() {
	for i := 0; i <= 225; i += 25 {
		url := "https://book.douban.com/top250?start=" + strconv.Itoa(i)
		resp := invoke.SendGet(url, nil, GetHeader())

		//defer关闭io流
		defer func(Body io.ReadCloser) {
			_ = Body.Close()
		}(resp.Body)
		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			logger.Write("GetTop250BookRanking goquery.NewDocumentFromReader异常:", err)
		}
		doc.Find("div#content").Find("div[class='indent']").Eq(0).Find("table").Each(func(i int, selection *goquery.Selection) {
			details, _ := selection.Find("td").Eq(0).Find("a").Attr("href")
			imgUrl, _ := selection.Find("td").Eq(0).Find("a").Eq(0).Find("img").Attr("src")
			name, _ := selection.Find("td").Eq(1).Find("div").Eq(0).Find("a").Attr("title")
			authorPressPublicationPrice := selection.Find("td").Eq(1).Find("p").Eq(0).Text()
			score := selection.Find("td").Eq(1).Find("div").Eq(1).Find("span").Eq(1).Text()
			scorePeople := selection.Find("td").Eq(1).Find("div").Eq(1).Find("span").Eq(2).Text()
			compile, _ := regexp.Compile(`[^0-9]`)
			matchString := strings.TrimSpace(compile.ReplaceAllString(scorePeople, ""))
			parseUint, _ := strconv.ParseUint(matchString, 10, 32)
			var quote string
			if selection.Find("td").Eq(1).Find("p").Length() < 2 {
				quote = ""
			} else {
				quote = selection.Find("td").Eq(1).Find("p").Eq(1).Find("span").Text()
			}
			var doubanBookTop250 = mysql.DoubanBookTop250{
				Details:                     details,
				ImgUrl:                      imgUrl,
				Name:                        name,
				AuthorPressPublicationPrice: authorPressPublicationPrice,
				Score:                       score,
				ScorePeople:                 uint(parseUint),
				Quote:                       quote,
			}
			fmt.Println(doubanBookTop250)
		})
		logger.Write("GetTop250BookRanking 豆瓣TOP250第" + strconv.Itoa(i/25+1) + "页读书数据爬取完成")
	}
}

func GetHotTestPublishBookRanking() {
	for i := 1; i <= 50; i++ {
		url := "https://read.douban.com/j/kind/"
		var params = make(map[string]any)
		params["kind"] = 1
		params["page"] = i
		params["query"] = "\n    query getFilterWorksList($works_ids: [ID!]) {\n      worksList(worksIds: $works_ids) {\n        \n    \n    title\n    cover(useSmall: false)\n    url\n    isBundle\n    coverLabel(preferVip: true)\n  \n    \n  url\n  title\n\n    \n  author {\n    name\n    url\n  }\n  origAuthor {\n    name\n    url\n  }\n  translator {\n    name\n    url\n  }\n\n    \n  abstract\n  authorHighlight\n  editorHighlight\n\n    \n    isOrigin\n    kinds {\n      \n    name @skip(if: true)\n    shortName @include(if: true)\n    id\n  \n    }\n    ... on WorksBase @include(if: true) {\n      wordCount\n      wordCountUnit\n    }\n    ... on WorksBase @include(if: false) {\n      inLibraryCount\n    }\n    ... on WorksBase @include(if: false) {\n      \n    isEssay\n    \n    ... on EssayWorks {\n      favorCount\n    }\n  \n    \n    \n    averageRating\n    ratingCount\n    url\n    isColumn\n    isFinished\n  \n  \n  \n    }\n    ... on EbookWorks @include(if: false) {\n      \n    ... on EbookWorks {\n      book {\n        url\n        averageRating\n        ratingCount\n      }\n    }\n  \n    }\n    ... on WorksBase @include(if: false) {\n      isColumn\n      isEssay\n      onSaleTime\n      ... on ColumnWorks {\n        updateTime\n      }\n    }\n    ... on WorksBase @include(if: true) {\n      isColumn\n      ... on ColumnWorks {\n        isFinished\n      }\n    }\n    ... on EssayWorks {\n      essayActivityData {\n        \n    title\n    uri\n    tag {\n      name\n      color\n      background\n      icon2x\n      icon3x\n      iconSize {\n        height\n      }\n      iconPosition {\n        x y\n      }\n    }\n  \n      }\n    }\n    highlightTags {\n      name\n    }\n    ... on WorksBase @include(if: false) {\n      fanfiction {\n        tags {\n          id\n          name\n          url\n        }\n      }\n    }\n  \n    \n  ... on WorksBase {\n    copyrightInfo {\n      newlyAdapted\n      newlyPublished\n      adaptedName\n      publishedName\n    }\n  }\n\n    isInLibrary\n    ... on WorksBase @include(if: false) {\n      \n    fixedPrice\n    salesPrice\n    isRebate\n  \n    }\n    ... on EbookWorks {\n      \n    fixedPrice\n    salesPrice\n    isRebate\n  \n    }\n    ... on WorksBase @include(if: true) {\n      ... on EbookWorks {\n        id\n        isPurchased\n        isInWishlist\n      }\n    }\n    ... on WorksBase @include(if: false) {\n      fanfiction {\n        fandoms {\n          title\n          url\n        }\n      }\n    }\n    ... on WorksBase @include(if: false) {\n      fanfiction {\n        kudoCount\n      }\n    }\n  \n        id\n        isOrigin\n      }\n    }\n  "
		params["sort"] = "hot"
		marshal, _ := json.Marshal(params)
		resp := invoke.SendPost(url, marshal, GetHeaderPost())

		//defer关闭io流
		defer func(Body io.ReadCloser) {
			_ = Body.Close()
		}(resp.Body)
		body, _ := io.ReadAll(resp.Body)

		for _, item := range gjson.Get(string(body), "list").Array() {
			name := item.Get("title").String()
			authors := ""
			for _, author := range item.Get("author").Array() {
				authors += author.Get("name").String() + "/"
			}
			origAuthors := ""
			for _, origAuthor := range item.Get("origAuthor").Array() {
				origAuthors += origAuthor.Get("name").String() + "/"
			}
			translators := ""
			for _, translator := range item.Get("translator").Array() {
				translators += translator.Get("name").String() + "/"
			}
			all := strings.ReplaceAll(item.Get("abstract").String(), " ", "")
			summary := strings.ReplaceAll(all, "\n", "")
			kindss := ""
			for _, kinds := range item.Get("kinds").Array() {
				kindss += kinds.Get("shortName").String() + "/"
			}
			wordCount := item.Get("wordCount").String()
			fixedPrice := float32(item.Get("fixedPrice").Int()) / 100
			salesPrice := float32(item.Get("salesPrice").Int()) / 100
			details := "https://read.douban.com" + item.Get("url").String() + "?dcs=category"
			imgUrl := item.Get("cover").String()
			var doubanBookHottestPublish = mysql.DoubanBookHottestPublish{
				Name:       name,
				Author:     authors,
				OrigAuthor: origAuthors,
				Translator: translators,
				Summary:    summary,
				Kinds:      kindss,
				WordCount:  wordCount,
				FixedPrice: fixedPrice,
				SalesPrice: salesPrice,
				Details:    details,
				ImgUrl:     imgUrl,
			}
			fmt.Println(doubanBookHottestPublish)
		}
		logger.Write("GetHotTestPublishBookRanking 豆瓣出版书籍中热度最高排行第" + strconv.Itoa(i) + "页数据爬取成功")
	}
}

func GetHighSalesPublishBookRanking() {
	for i := 1; i <= 50; i++ {
		url := "https://read.douban.com/j/kind/"
		var params = make(map[string]any)
		params["kind"] = 1
		params["page"] = i
		params["query"] = "\n    query getFilterWorksList($works_ids: [ID!]) {\n      worksList(worksIds: $works_ids) {\n        \n    \n    title\n    cover(useSmall: false)\n    url\n    isBundle\n    coverLabel(preferVip: true)\n  \n    \n  url\n  title\n\n    \n  author {\n    name\n    url\n  }\n  origAuthor {\n    name\n    url\n  }\n  translator {\n    name\n    url\n  }\n\n    \n  abstract\n  authorHighlight\n  editorHighlight\n\n    \n    isOrigin\n    kinds {\n      \n    name @skip(if: true)\n    shortName @include(if: true)\n    id\n  \n    }\n    ... on WorksBase @include(if: true) {\n      wordCount\n      wordCountUnit\n    }\n    ... on WorksBase @include(if: false) {\n      inLibraryCount\n    }\n    ... on WorksBase @include(if: false) {\n      \n    isEssay\n    \n    ... on EssayWorks {\n      favorCount\n    }\n  \n    \n    \n    averageRating\n    ratingCount\n    url\n    isColumn\n    isFinished\n  \n  \n  \n    }\n    ... on EbookWorks @include(if: false) {\n      \n    ... on EbookWorks {\n      book {\n        url\n        averageRating\n        ratingCount\n      }\n    }\n  \n    }\n    ... on WorksBase @include(if: false) {\n      isColumn\n      isEssay\n      onSaleTime\n      ... on ColumnWorks {\n        updateTime\n      }\n    }\n    ... on WorksBase @include(if: true) {\n      isColumn\n      ... on ColumnWorks {\n        isFinished\n      }\n    }\n    ... on EssayWorks {\n      essayActivityData {\n        \n    title\n    uri\n    tag {\n      name\n      color\n      background\n      icon2x\n      icon3x\n      iconSize {\n        height\n      }\n      iconPosition {\n        x y\n      }\n    }\n  \n      }\n    }\n    highlightTags {\n      name\n    }\n    ... on WorksBase @include(if: false) {\n      fanfiction {\n        tags {\n          id\n          name\n          url\n        }\n      }\n    }\n  \n    \n  ... on WorksBase {\n    copyrightInfo {\n      newlyAdapted\n      newlyPublished\n      adaptedName\n      publishedName\n    }\n  }\n\n    isInLibrary\n    ... on WorksBase @include(if: false) {\n      \n    fixedPrice\n    salesPrice\n    isRebate\n  \n    }\n    ... on EbookWorks {\n      \n    fixedPrice\n    salesPrice\n    isRebate\n  \n    }\n    ... on WorksBase @include(if: true) {\n      ... on EbookWorks {\n        id\n        isPurchased\n        isInWishlist\n      }\n    }\n    ... on WorksBase @include(if: false) {\n      fanfiction {\n        fandoms {\n          title\n          url\n        }\n      }\n    }\n    ... on WorksBase @include(if: false) {\n      fanfiction {\n        kudoCount\n      }\n    }\n  \n        id\n        isOrigin\n      }\n    }\n  "
		params["sort"] = "sales"
		marshal, _ := json.Marshal(params)
		resp := invoke.SendPost(url, marshal, GetHeaderPost())

		//defer关闭io流
		defer func(Body io.ReadCloser) {
			_ = Body.Close()
		}(resp.Body)
		body, _ := io.ReadAll(resp.Body)

		for _, item := range gjson.Get(string(body), "list").Array() {
			name := item.Get("title").String()
			authors := ""
			for _, author := range item.Get("author").Array() {
				authors += author.Get("name").String() + "/"
			}
			origAuthors := ""
			for _, origAuthor := range item.Get("origAuthor").Array() {
				origAuthors += origAuthor.Get("name").String() + "/"
			}
			translators := ""
			for _, translator := range item.Get("translator").Array() {
				translators += translator.Get("name").String() + "/"
			}
			all := strings.ReplaceAll(item.Get("abstract").String(), " ", "")
			summary := strings.ReplaceAll(all, "\n", "")
			kindss := ""
			for _, kinds := range item.Get("kinds").Array() {
				kindss += kinds.Get("shortName").String() + "/"
			}
			wordCount := item.Get("wordCount").String()
			fixedPrice := float32(item.Get("fixedPrice").Int()) / 100
			salesPrice := float32(item.Get("salesPrice").Int()) / 100
			details := "https://read.douban.com" + item.Get("url").String() + "?dcs=category"
			imgUrl := item.Get("cover").String()
			var doubanBookHighsalesPublish = mysql.DoubanBookHighsalesPublish{
				Name:       name,
				Author:     authors,
				OrigAuthor: origAuthors,
				Translator: translators,
				Summary:    summary,
				Kinds:      kindss,
				WordCount:  wordCount,
				FixedPrice: fixedPrice,
				SalesPrice: salesPrice,
				Details:    details,
				ImgUrl:     imgUrl,
			}
			fmt.Println(doubanBookHighsalesPublish)
		}
		logger.Write("GetHighSalesPublishBookRanking 豆瓣出版书籍中销量最高排行第" + strconv.Itoa(i) + "页数据爬取成功")
	}
}

func GetHotTestOriginalBookRanking() {
	for i := 1; i <= 50; i++ {
		url := "https://read.douban.com/j/kind/"
		var params = make(map[string]any)
		params["kind"] = 0
		params["page"] = i
		params["query"] = "\n    query getFilterWorksList($works_ids: [ID!]) {\n      worksList(worksIds: $works_ids) {\n        \n    \n    title\n    cover(useSmall: false)\n    url\n    isBundle\n    coverLabel(preferVip: true)\n  \n    \n  url\n  title\n\n    \n  author {\n    name\n    url\n  }\n  origAuthor {\n    name\n    url\n  }\n  translator {\n    name\n    url\n  }\n\n    \n  abstract\n  authorHighlight\n  editorHighlight\n\n    \n    isOrigin\n    kinds {\n      \n    name @skip(if: true)\n    shortName @include(if: true)\n    id\n  \n    }\n    ... on WorksBase @include(if: true) {\n      wordCount\n      wordCountUnit\n    }\n    ... on WorksBase @include(if: false) {\n      inLibraryCount\n    }\n    ... on WorksBase @include(if: false) {\n      \n    isEssay\n    \n    ... on EssayWorks {\n      favorCount\n    }\n  \n    \n    \n    averageRating\n    ratingCount\n    url\n    isColumn\n    isFinished\n  \n  \n  \n    }\n    ... on EbookWorks @include(if: false) {\n      \n    ... on EbookWorks {\n      book {\n        url\n        averageRating\n        ratingCount\n      }\n    }\n  \n    }\n    ... on WorksBase @include(if: false) {\n      isColumn\n      isEssay\n      onSaleTime\n      ... on ColumnWorks {\n        updateTime\n      }\n    }\n    ... on WorksBase @include(if: true) {\n      isColumn\n      ... on ColumnWorks {\n        isFinished\n      }\n    }\n    ... on EssayWorks {\n      essayActivityData {\n        \n    title\n    uri\n    tag {\n      name\n      color\n      background\n      icon2x\n      icon3x\n      iconSize {\n        height\n      }\n      iconPosition {\n        x y\n      }\n    }\n  \n      }\n    }\n    highlightTags {\n      name\n    }\n    ... on WorksBase @include(if: false) {\n      fanfiction {\n        tags {\n          id\n          name\n          url\n        }\n      }\n    }\n  \n    \n  ... on WorksBase {\n    copyrightInfo {\n      newlyAdapted\n      newlyPublished\n      adaptedName\n      publishedName\n    }\n  }\n\n    isInLibrary\n    ... on WorksBase @include(if: false) {\n      \n    fixedPrice\n    salesPrice\n    isRebate\n  \n    }\n    ... on EbookWorks {\n      \n    fixedPrice\n    salesPrice\n    isRebate\n  \n    }\n    ... on WorksBase @include(if: true) {\n      ... on EbookWorks {\n        id\n        isPurchased\n        isInWishlist\n      }\n    }\n    ... on WorksBase @include(if: false) {\n      fanfiction {\n        fandoms {\n          title\n          url\n        }\n      }\n    }\n    ... on WorksBase @include(if: false) {\n      fanfiction {\n        kudoCount\n      }\n    }\n  \n        id\n        isOrigin\n      }\n    }\n  "
		params["sort"] = "hot"
		marshal, _ := json.Marshal(params)
		resp := invoke.SendPost(url, marshal, GetHeaderPost())

		//defer关闭io流
		defer func(Body io.ReadCloser) {
			_ = Body.Close()
		}(resp.Body)
		body, _ := io.ReadAll(resp.Body)

		for _, item := range gjson.Get(string(body), "list").Array() {
			name := item.Get("title").String()
			authors := ""
			for _, author := range item.Get("author").Array() {
				authors += author.Get("name").String() + "/"
			}
			origAuthors := ""
			for _, origAuthor := range item.Get("origAuthor").Array() {
				origAuthors += origAuthor.Get("name").String() + "/"
			}
			translators := ""
			for _, translator := range item.Get("translator").Array() {
				translators += translator.Get("name").String() + "/"
			}
			//all := strings.ReplaceAll(item.Get("abstract").String(), " ", "")
			//summary := strings.ReplaceAll(all, "\n", "")
			kindss := ""
			for _, kinds := range item.Get("kinds").Array() {
				kindss += kinds.Get("shortName").String() + "/"
			}
			wordCount := item.Get("wordCount").String()
			fixedPrice := float32(item.Get("fixedPrice").Int()) / 100
			salesPrice := float32(item.Get("salesPrice").Int()) / 100
			details := "https://read.douban.com" + item.Get("url").String() + "?dcs=category"
			imgUrl := item.Get("cover").String()
			var doubanBookHottestOriginal = mysql.DoubanBookHottestOriginal{
				Name:       name,
				Author:     authors,
				OrigAuthor: origAuthors,
				Translator: translators,
				Summary:    "summary",
				Kinds:      kindss,
				WordCount:  wordCount,
				FixedPrice: fixedPrice,
				SalesPrice: salesPrice,
				Details:    details,
				ImgUrl:     imgUrl,
			}
			fmt.Println(doubanBookHottestOriginal)
		}
		logger.Write("GetHotTestOriginalBookRanking 豆瓣原创书籍中热度最高排行第" + strconv.Itoa(i) + "页数据爬取成功")
	}
}

func GetHeader() map[string]string {
	var headers = make(map[string]string)
	headers["User-Agent"] = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36"
	return headers
}

func GetHeaderPost() map[string]string {
	var headers = make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["User-Agent"] = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36"
	return headers
}
