package zb

import (
	"github.com/spf13/viper"
	"log"
	"testing"
)

func testWS() *ZBWebsocket {
	viper.SetConfigName("test_config")
	viper.AddConfigPath("./testdata")
	err := viper.ReadInConfig()
	if err != nil {
		log.Panic(err)
	}

	accessKey := viper.GetString("access_key")
	secretKey := viper.GetString("secret_key")
	ws := NewZBWebsocket(accessKey, secretKey, true)
	return ws
}

func TestZBWebsocket_SubscribeTicker(t *testing.T) {
	ws := testWS()
	ws.Start()

	ws.SetTickerCallback(func(ticker *WSTicker) {

	})
	ws.SubscribeTicker("zbqc")

	select {}
}

func TestZBWebsocket_SubscribeDepth(t *testing.T) {
	ws := testWS()
	ws.Start()

	ws.SetDepthCallback(func(depth *WSDepth) {

	})
	ws.SubscribeDepth("zbqc")

	select {}
}

func TestZBWebsocket_SubscribeTrades(t *testing.T) {
	ws := testWS()
	ws.Start()

	ws.SetTradesCallback(func(trades *WSTrades) {

	})
	ws.SubscribeTrades("zbqc")

	select {}
}
