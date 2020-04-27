package zb

import (
	"testing"
	"time"
)

func TestZB_GetAllTicker(t *testing.T) {
	z := testZB()
	tickers, err := z.GetAllTicker()
	if err != nil {
		t.Error(err)
		return
	}
	for k, ticker := range tickers {
		t.Logf("%v: %#v", k, ticker)
	}
}

func TestZB_GetTicker(t *testing.T) {
	z := testZB()
	ticker, err := z.GetTicker("btc_usdt")
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%#v", ticker)
}

func TestZB_GetDepth(t *testing.T) {
	z := testZB()
	depth, err := z.GetDepth("btc_usdt", 10)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%#v", depth)
}

func TestZB_GetTrades(t *testing.T) {
	z := testZB()
	trades, err := z.GetTrades("btc_usdt")
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%#v", trades)
}

func TestZB_GetKLine(t *testing.T) {
	z := testZB()
	since := time.Now().Add(-1 * time.Hour)
	kLine, err := z.GetKLine("btc_usdt",
		"1min", since.UnixNano()/int64(time.Millisecond), 1000)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%#v", kLine)
}
