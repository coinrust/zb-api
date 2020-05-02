package zb

import (
	"testing"
)

func testWS() *ZBWebsocket {
	cfg := loadConfig()

	accessKey := cfg.Get("access_key").String()
	secretKey := cfg.Get("secret_key").String()
	ws := NewZBWebsocket("zb.live", accessKey, secretKey, true)
	return ws
}

func TestZBWebsocket_SubscribeTicker(t *testing.T) {
	ws := testWS()

	ws.SetTickerCallback(func(ticker *WSTicker) {

	})
	ws.SubscribeTicker("zbqc")

	select {}
}

func TestZBWebsocket_SubscribeDepth(t *testing.T) {
	ws := testWS()

	ws.SetDepthCallback(func(depth *WSDepth) {

	})
	ws.SubscribeDepth("zbqc")

	select {}
}

func TestZBWebsocket_SubscribeTrades(t *testing.T) {
	ws := testWS()

	ws.SetTradesCallback(func(trades *WSTrades) {

	})
	ws.SubscribeTrades("zbqc")

	select {}
}
