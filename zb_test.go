package zb

import (
	"github.com/spf13/viper"
	"log"
	"testing"
)

func testZB() *ZB {
	viper.SetConfigName("test_config")
	viper.AddConfigPath("./testdata")
	err := viper.ReadInConfig()
	if err != nil {
		log.Panic(err)
	}

	accessKey := viper.GetString("access_key")
	secretKey := viper.GetString("secret_key")
	host := "zb.live"
	z := NewZB(nil, host, accessKey, secretKey, true)
	return z
}

func testZB2() *ZB {
	viper.SetConfigName("test_config")
	viper.AddConfigPath("./testdata")
	err := viper.ReadInConfig()
	if err != nil {
		log.Panic(err)
	}

	accessKey := viper.GetString("access_key")
	secretKey := viper.GetString("secret_key")
	host := "zb.com"
	z := NewZB(nil, host, accessKey, secretKey, true)
	return z
}

func TestZB_SetProxy(t *testing.T) {
	z := testZB2()
	z.SetProxy("socks5://127.0.0.1:1080")
	ticker, err := z.GetTicker("usdt_qc")
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%#v", ticker)
}
