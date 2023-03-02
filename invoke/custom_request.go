package invoke

import (
	"bytes"
	"net/http"
	"net/url"
)

// SendGet 封装GET请求
func SendGet(reqUrl string, param map[string]string, header map[string]string) (resp *http.Response, err error) {
	if param == nil && header == nil {
		resp, err := http.Get(reqUrl)
		if err != nil {
			return nil, err
		}
		return resp, err

	} else if param == nil {
		client := &http.Client{}
		req, _ := http.NewRequest("GET", reqUrl, nil)
		for k, v := range header {
			req.Header.Add(k, v)
		}
		resp, err := client.Do(req)
		if err != nil {
			return nil, err
		}
		return resp, err

	} else if header == nil {
		Url, _ := url.Parse(reqUrl)
		params := url.Values{}
		for k, v := range param {
			params.Set(k, v)
		}
		Url.RawQuery = params.Encode()
		resp, err := http.Get(Url.String())
		if err != nil {
			return nil, err
		}
		return resp, err
	}
	return nil, err
}

// SendPost 封装POST请求
func SendPost(reqUrl string, param []byte, header map[string]string) (resp *http.Response, err error) {
	client := &http.Client{}
	if header == nil {
		req, _ := http.NewRequest("POST", reqUrl, bytes.NewReader(param))
		resp, err = client.Do(req)
		if err != nil {
			return nil, err
		}
		return resp, err
	} else {
		req, _ := http.NewRequest("POST", reqUrl, bytes.NewReader(param))
		for k, v := range header {
			req.Header.Add(k, v)
		}
		resp, err = client.Do(req)
		if err != nil {
			return nil, err
		}
		return resp, err
	}
}
