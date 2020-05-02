package zb

import (
	"go.uber.org/config"
	"testing"
)

func loadConfig() config.Provider {
	provider, err := config.NewYAML(config.File("./testdata/test_config.yaml"))
	if err != nil {
		panic(err)
	}
	return provider
}

func testZB() *ZB {
	cfg := loadConfig()

	accessKey := cfg.Get("access_key").String()
	secretKey := cfg.Get("secret_key").String()
	host := "zb.live"
	z := NewZB(nil, host, accessKey, secretKey, true)
	return z
}

func testZB2() *ZB {
	cfg := loadConfig()

	accessKey := cfg.Get("access_key").String()
	secretKey := cfg.Get("secret_key").String()
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
