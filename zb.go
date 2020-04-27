package zb

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/json-iterator/go"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	OK = 1000
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type ZB struct {
	httpClient *http.Client
	accessKey  string
	secretKey  string
	debugMode  bool

	marketURL string // "http://api.zb.live/data/v1/"
	tradeURL  string // "https://trade.zb.live/api/"
}

func (z *ZB) HttpGet(reqUrl string, postData string, headers map[string]string, result interface{}) ([]byte, error) {
	return z.NewHttpRequest("GET", reqUrl, postData, headers, result)
}

func (z *ZB) HttpPost(reqUrl string, postData string, headers map[string]string, result interface{}) ([]byte, error) {
	return z.NewHttpRequest("POST", reqUrl, postData, headers, result)
}

func (z *ZB) NewHttpRequest(method string, reqUrl string, postData string, requestHeaders map[string]string, result interface{}) ([]byte, error) {
	if z.debugMode {
		log.Printf("reqUrl: %v", reqUrl)
	}
	req, _ := http.NewRequest(method, reqUrl, strings.NewReader(postData))
	//req.Header.Set("User-Agent", defaultUserAgent)

	if requestHeaders != nil {
		for k, v := range requestHeaders {
			req.Header.Add(k, v)
		}
	}

	resp, err := z.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	bodyData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if z.debugMode {
		log.Printf("%v", string(bodyData))
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("HttpStatusCode: %d, Desc: %s", resp.StatusCode, string(bodyData)))
	}

	if result != nil {
		err = json.Unmarshal(bodyData, result)
	}

	return bodyData, err
}

func (z *ZB) buildSignParams(params *url.Values) error {
	params.Set("accesskey", z.accessKey)
	payload := params.Encode()
	secretKeySha, _ := GetSHA(z.secretKey)
	sign, err := GetParamHmacMD5Sign(secretKeySha, payload)
	if err != nil {
		return err
	}
	params.Set("sign", sign)
	params.Set("reqTime", fmt.Sprintf("%d", time.Now().UnixNano()/int64(time.Millisecond)))
	return nil
}

// SetProxy 设置代理
// proxyURL: "socks5://127.0.0.1:1080"
func (z *ZB) SetProxy(proxyURL string) (err error) {
	transport := CloneDefaultTransport()
	transport.Proxy, err = ParseProxy(proxyURL)
	if err != nil {
		return
	}
	z.httpClient.Transport = transport
	return nil
}

// domain: zb.live/zb.com
func NewZB(httpClient *http.Client, host string, accessKey string, secretKey string, debugMode bool) *ZB {
	marketURL := fmt.Sprintf("http://api.%v/data/v1/", host) // "http://api.zb.live/data/v1/"
	tradeURL := fmt.Sprintf("https://trade.%v/api/", host)   // "https://trade.zb.live/api/"
	if httpClient == nil {
		httpClient = &http.Client{
			Timeout: 10 * time.Second,
		}
	}
	return &ZB{
		httpClient: httpClient,
		accessKey:  accessKey,
		secretKey:  secretKey,
		debugMode:  debugMode,
		marketURL:  marketURL,
		tradeURL:   tradeURL,
	}
}

func GetSHA(text string) (string, error) {
	sha := sha1.New()
	_, err := sha.Write([]byte(text))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(sha.Sum(nil)), nil
}

func GetParamHmacSHA256Sign(secret, params string) (string, error) {
	mac := hmac.New(sha256.New, []byte(secret))
	_, err := mac.Write([]byte(params))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(mac.Sum(nil)), nil
}

func GetParamHmacMD5Sign(secret, params string) (string, error) {
	mac := hmac.New(md5.New, []byte(secret))
	_, err := mac.Write([]byte(params))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(mac.Sum(nil)), nil
}

/**
 *md5签名
 */
func GetParamMD5Sign(secret, params string) (string, error) {
	hash := md5.New()
	_, err := hash.Write([]byte(params))

	if err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}

func GetParamHmacSHA256Base64Sign(secret, params string) (string, error) {
	mac := hmac.New(sha256.New, []byte(secret))
	_, err := mac.Write([]byte(params))
	if err != nil {
		return "", err
	}
	signByte := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(signByte), nil
}
