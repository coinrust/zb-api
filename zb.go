package zb

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/json-iterator/go"
	"net/http"
	"net/url"
	"time"
)

const (
	OK = 1000
)

var (
	MARKET_URL = "http://api.zb.live/data/v1/"
	TICKER_API = "ticker?market=%s"
	DEPTH_API  = "depth?market=%s&size=%d"
	TRADE_URL  = "https://trade.zb.live/api/"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type ZB struct {
	httpClient *http.Client
	accessKey  string
	secretKey  string
	debugMode  bool
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

func NewZB(httpClient *http.Client, accessKey string, secretKey string, debugMode bool) *ZB {
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
