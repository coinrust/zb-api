package zb

import (
	"github.com/spf13/viper"
	"log"
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
	z := NewZB(nil, accessKey, secretKey, true)
	return z
}
