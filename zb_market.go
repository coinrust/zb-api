package zb

import (
	"fmt"
)

// GetAllTicker 获取allTicker数据
func (z *ZB) GetAllTicker() (tickers AllTicker, err error) {
	url := z.marketURL + "allTicker"
	_, err = z.HttpGet(url, "", nil, &tickers)
	return
}

// GetTicker 获取Ticker数据
func (z *ZB) GetTicker(symbol string) (ticker Ticker, err error) {
	url := z.marketURL + fmt.Sprintf("ticker?market=%s", symbol)
	_, err = z.HttpGet(url, "", nil, &ticker)
	return
}

// GetDepth 获取市场深度
// size: 1-50
func (z *ZB) GetDepth(symbol string, size int) (depth Depth, err error) {
	url := z.marketURL + fmt.Sprintf("depth?market=%s&size=%d", symbol, size)
	_, err = z.HttpGet(url, "", nil, &depth)
	return
}

// GetTrades 获取历史成交
// symbol: btc_usdt
func (z *ZB) GetTrades(symbol string) (trades []Trade, err error) {
	url := z.marketURL + fmt.Sprintf("trades?market=%v", symbol)
	_, err = z.HttpGet(url, "", nil, &trades)
	return
}

// GetKLine 获取K线
// symbol: btc_usdt
// _type: 1min/3min/5min/15min/30min/1day/3day/1week/1hour/2hour/4hour/6hour/12hour
// since: 从这个时间戳之后的 (ms)
// size: 返回数据的条数限制（默认为1000，如果返回数据多于1000条，那么只返回1000条）
func (z *ZB) GetKLine(symbol string, _type string, since int64, size int) (kLine KLine, err error) {
	if size == 0 {
		size = 1000
	}
	url := z.marketURL + fmt.Sprintf("kline?market=%v&type=%v&since=%v&size=%v",
		symbol, _type, since, size)
	_, err = z.HttpGet(url, "", nil, &kLine)
	return
}
