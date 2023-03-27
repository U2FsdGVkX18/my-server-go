package notion

import (
	"github.com/tidwall/gjson"
	"io"
	"my-server-go/invoke"
)

// NotionApi
const notionBasicApi = "https://api.notion.com/v1"

// Notion机器人令牌
const notionBotKey = "secret_rvUNRXh5Imw1WmOudXxcOaF6yakOkXVSn3LuJg5uLrv"

// NotionApi版本
const notionVersion = "2022-06-28"

//func NotionSyncNewMovieRanking() {
//	dataBaseId := "7ec1b96bd47f4c5c86fab914c24b3c73"
//	pageUrl := notionBasicApi + "/pages"
//
//}

func DeleteDataBaseData(dataBaseId string) {
	dataBaseQueryUrl := notionBasicApi + "/databases/" + dataBaseId + "/query"
	for {
		resp := invoke.SendPost(dataBaseQueryUrl, nil, GetHeader())
		resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
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
