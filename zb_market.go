package zb

import (
	"fmt"
	"log"
)

// GetTicker 获取Ticker数据
func (z *ZB) GetTicker(symbol string) (ticker Ticker, err error) {
	var resp []byte
	url := MARKET_URL + fmt.Sprintf(TICKER_API, symbol)
	resp, err = HttpGet(z.httpClient, url, "", nil, &ticker)
	if err != nil {
		return
	}
	if z.debugMode {
		log.Printf("%v", string(resp))
	}
	return
}

// GetDepth 获取市场深度
// size: 1-50
func (z *ZB) GetDepth(symbol string, size int) (depth Depth, err error) {
	var resp []byte
	url := MARKET_URL + fmt.Sprintf(DEPTH_API, symbol, size)
	resp, err = HttpGet(z.httpClient, url, "", nil, &depth)
	if err != nil {
		return
	}
	if z.debugMode {
		log.Printf("%v", string(resp))
	}
	return
}
