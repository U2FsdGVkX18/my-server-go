package invoke

import (
	"bytes"
	logger "my-server-go/tool/log"
	"net/http"
	"net/url"
)

// SendGet 封装GET请求
func SendGet(reqUrl string, param map[string]string, header map[string]string) (resp *http.Response) {
	if param == nil && header == nil {
		resp, err := http.Get(reqUrl)
		if err != nil {
			logger.Write("SendGet get请求出错:", err)
			return nil
		}
		return resp

	} else if param == nil {
		client := &http.Client{}
		req, _ := http.NewRequest("GET", reqUrl, nil)
		for k, v := range header {
			req.Header.Add(k, v)
		}
		resp, err := client.Do(req)
		if err != nil {
			logger.Write("SendGet get请求出错:", err)
			return nil
		}
		return resp

	} else if header == nil {
		Url, _ := url.Parse(reqUrl)
		params := url.Values{}
		for k, v := range param {
			params.Set(k, v)
		}
		Url.RawQuery = params.Encode()
		resp, err := http.Get(Url.String())
		if err != nil {
			logger.Write("SendGet get请求出错:", err)
			return nil
		}
		return resp
	}
	return nil
}

// SendPost 封装POST请求
func SendPost(reqUrl string, param []byte, header map[string]string) (resp *http.Response) {
	client := &http.Client{}
	if header == nil {
		req, _ := http.NewRequest("POST", reqUrl, bytes.NewReader(param))
		resp, err := client.Do(req)
		if err != nil {
			logger.Write("SendPost post请求出错:", err)
			return nil
		}
		return resp
	} else {
		req, _ := http.NewRequest("POST", reqUrl, bytes.NewReader(param))
		for k, v := range header {
			req.Header.Add(k, v)
		}
		resp, err := client.Do(req)
		if err != nil {
			logger.Write("SendPost post请求出错:", err)
			return nil
		}
		return resp
	}
}
