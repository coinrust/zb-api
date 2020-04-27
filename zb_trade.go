package zb

import (
	"fmt"
	"github.com/tidwall/gjson"
	"log"
	"net/url"
	"strings"
	"time"
)

// PlaceOrder 委托下单
// tradeType: 交易类型1/0[buy/sell]
// acctType: 杠杆 1/0[杠杆/现货]（选填,默认为: 0 现货）
func (z *ZB) PlaceOrder(symbol string, price float64, amount float64, tradeType int, acctType int, customOrderID string) (result OrderResponse, err error) {
	params := url.Values{}
	params.Set("method", "order")
	params.Set("price", fmt.Sprint(price))
	params.Set("amount", fmt.Sprint(amount))
	params.Set("currency", symbol)
	params.Set("tradeType", fmt.Sprint(tradeType))
	params.Set("acctType", fmt.Sprint(acctType))
	if customOrderID != "" {
		params.Set("customerOrderId", customOrderID)
	}
	z.buildSignParams(&params)

	url := TRADE_URL + "order?" + params.Encode()
	_, err = z.HttpGet(url, "", nil, &result)
	// {"code":2009,"message":"账户余额不足"}
	return
}

// PlaceOrderMore 批量委托下单
// 交易参数，数组格式[[价格，数量]，[价格，数量]]，最多20个
// tradeType: 交易类型1/0[buy/sell]
// acctType: 杠杆 1/0[杠杆/现货]（选填,默认为: 0 现货）
func (z *ZB) PlaceOrderMore(symbol string, tradeParams [][]float64, tradeType int, acctType int) (result OrderMoreV2Response, err error) {
	var tradeParamsData []byte
	tradeParamsData, err = json.Marshal(tradeParams)
	if err != nil {
		return
	}
	params := url.Values{}
	params.Set("method", "orderMoreV2")
	params.Set("market", symbol)
	params.Set("tradeType", fmt.Sprint(tradeType))
	params.Set("acctType", fmt.Sprint(acctType))
	params.Set("tradeParams", "_tradeParams_")
	params.Set("accesskey", z.accessKey)

	payload := params.Encode()
	payload = strings.ReplaceAll(payload, "_tradeParams_", string(tradeParamsData))
	secretKeySha, _ := GetSHA(z.secretKey)
	var sign string
	sign, err = GetParamHmacMD5Sign(secretKeySha, payload)
	if err != nil {
		return
	}

	url := TRADE_URL + "orderMoreV2?" + payload + "&sign=" + sign + "&reqTime=" + fmt.Sprintf("%d", time.Now().UnixNano()/int64(time.Millisecond))
	if z.debugMode {
		log.Printf("url: %v", url)
	}
	_, err = z.HttpGet(url, "", nil, &result)
	// {"code":1003,"message":"验证不通过"}
	// {"code":2009,"message":"账户余额不足"}
	return
}

// CancelOrder 取消委托
func (z *ZB) CancelOrder(symbol string, id int64, customOrderID string) (result OrderResponse, err error) {
	params := url.Values{}
	params.Set("method", "cancelOrder")
	params.Set("currency", symbol)
	if customOrderID != "" {
		params.Set("customerOrderId", customOrderID)
	} else {
		params.Set("id", fmt.Sprint(id))
	}
	z.buildSignParams(&params)

	url := TRADE_URL + "cancelOrder?" + params.Encode()
	_, err = z.HttpGet(url, "", nil, &result)
	// {"code":3001,"message":"挂单没有找到或已完成"}
	return
}

// GetOrder 获取委托单
func (z *ZB) GetOrder(symbol string, id int64, customOrderID string) (result Order, err error) {
	params := url.Values{}
	params.Set("method", "getOrder")
	params.Set("currency", symbol)
	if customOrderID != "" {
		params.Set("customerOrderId", customOrderID)
	} else {
		params.Set("id", fmt.Sprint(id))
	}
	z.buildSignParams(&params)

	url := TRADE_URL + "getOrder?" + params.Encode()
	var resp []byte
	resp, err = z.HttpGet(url, "", nil, nil)
	if err != nil {
		return
	}
	// {"code":3001,"message":"挂单没有找到或已完成"}
	g := gjson.ParseBytes(resp)
	if codeValue := g.Get("code"); codeValue.Exists() {
		err = fmt.Errorf("code=%v message=%v",
			codeValue.Int(), g.Get("message").String())
		return
	}
	err = json.Unmarshal(resp, &result)
	return
}

// GetOrders 获取多个委托单
func (z *ZB) GetOrders(symbol string, tradeType int, pageIndex int) (result []Order, err error) {
	params := url.Values{}
	params.Set("method", "getOrders")
	params.Set("currency", symbol)
	params.Set("tradeType", fmt.Sprint(tradeType))
	params.Set("pageIndex", fmt.Sprint(pageIndex))
	z.buildSignParams(&params)

	url := TRADE_URL + "getOrders?" + params.Encode()
	var resp []byte
	resp, err = z.HttpGet(url, "", nil, nil)
	if err != nil {
		return
	}
	g := gjson.ParseBytes(resp)
	if codeValue := g.Get("code"); codeValue.Exists() {
		err = fmt.Errorf("code=%v message=%v",
			codeValue.Int(), g.Get("message").String())
		return
	}
	err = json.Unmarshal(resp, &result)
	return
}

// GetOrdersIgnoreTradeType 获取委托单忽略类型
func (z *ZB) GetOrdersIgnoreTradeType(symbol string, pageIndex int, pageSize int) (result []Order, err error) {
	params := url.Values{}
	params.Set("method", "getOrdersIgnoreTradeType")
	params.Set("currency", symbol)
	params.Set("pageIndex", fmt.Sprint(pageIndex))
	params.Set("pageSize", fmt.Sprint(pageSize))
	z.buildSignParams(&params)

	url := TRADE_URL + "getOrdersIgnoreTradeType?" + params.Encode()
	var resp []byte
	resp, err = z.HttpGet(url, "", nil, nil)
	if err != nil {
		return
	}
	g := gjson.ParseBytes(resp)
	if codeValue := g.Get("code"); codeValue.Exists() {
		err = fmt.Errorf("code=%v message=%v",
			codeValue.Int(), g.Get("message").String())
		return
	}
	err = json.Unmarshal(resp, &result)
	return
}

// GetUnfinishedOrdersIgnoreTradeType 未成交或部份成交挂单
func (z *ZB) GetUnfinishedOrdersIgnoreTradeType(symbol string, pageIndex int, pageSize int) (result []Order, err error) {
	params := url.Values{}
	params.Set("method", "getUnfinishedOrdersIgnoreTradeType")
	params.Set("currency", symbol)
	params.Set("pageIndex", fmt.Sprint(pageIndex))
	params.Set("pageSize", fmt.Sprint(pageSize))
	z.buildSignParams(&params)

	url := TRADE_URL + "getUnfinishedOrdersIgnoreTradeType?" + params.Encode()
	var resp []byte
	resp, err = z.HttpGet(url, "", nil, nil)
	if err != nil {
		return
	}
	// {"code":3001,"message":"挂单没有找到或已完成"}
	g := gjson.ParseBytes(resp)
	if codeValue := g.Get("code"); codeValue.Exists() {
		err = fmt.Errorf("code=%v message=%v",
			codeValue.Int(), g.Get("message").String())
		return
	}
	err = json.Unmarshal(resp, &result)
	return
}

// GetAccountInfo 获取用户信息
func (z *ZB) GetAccountInfo() (result AccountInfo, err error) {
	params := url.Values{}
	params.Set("method", "getAccountInfo")
	z.buildSignParams(&params)

	url := TRADE_URL + "getAccountInfo?" + params.Encode()
	var resp []byte
	resp, err = z.HttpGet(url, "", nil, nil)
	if err != nil {
		return
	}
	// {"code":3001,"message":"挂单没有找到或已完成"}
	g := gjson.ParseBytes(resp)
	if codeValue := g.Get("code"); codeValue.Exists() {
		err = fmt.Errorf("code=%v message=%v",
			codeValue.Int(), g.Get("message").String())
		return
	}
	err = json.Unmarshal(resp, &result)
	return
}
