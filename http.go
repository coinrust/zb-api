package zb

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func HttpGet(client *http.Client, reqUrl string, postData string, headers map[string]string, result interface{}) ([]byte, error) {
	return NewHttpRequest(client, "GET", reqUrl, postData, headers, result)
}

func HttpPost(client *http.Client, reqUrl string, postData string, headers map[string]string, result interface{}) ([]byte, error) {
	return NewHttpRequest(client, "POST", reqUrl, postData, headers, result)
}

func NewHttpRequest(client *http.Client, method string, reqUrl string, postData string, requestHeaders map[string]string, result interface{}) ([]byte, error) {
	req, _ := http.NewRequest(method, reqUrl, strings.NewReader(postData))
	//req.Header.Set("User-Agent", defaultUserAgent)

	if requestHeaders != nil {
		for k, v := range requestHeaders {
			req.Header.Add(k, v)
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	bodyData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("HttpStatusCode: %d, Desc: %s", resp.StatusCode, string(bodyData)))
	}

	if result != nil {
		err = json.Unmarshal(bodyData, result)
	}

	return bodyData, err
}
