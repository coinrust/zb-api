package zb

import (
	"net"
	"net/http"
	"net/url"
	"time"
)

var (
	DefaultTransport *http.Transport
)

func init() {
	DefaultTransport = &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   5 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		IdleConnTimeout:        1 * time.Minute,
		TLSHandshakeTimeout:    10 * time.Second,
		ExpectContinueTimeout:  1 * time.Second,
		DisableKeepAlives:      false,
		MaxResponseHeaderBytes: 1 << 15,
	}
}

func CloneDefaultTransport() *http.Transport {
	return &http.Transport{DialContext: DefaultTransport.DialContext,
		IdleConnTimeout:        DefaultTransport.IdleConnTimeout,
		TLSHandshakeTimeout:    DefaultTransport.TLSHandshakeTimeout,
		ExpectContinueTimeout:  DefaultTransport.ExpectContinueTimeout,
		MaxResponseHeaderBytes: DefaultTransport.MaxResponseHeaderBytes,
		Proxy:                  DefaultTransport.Proxy,
		DisableKeepAlives:      DefaultTransport.DisableKeepAlives,
	}
}

// "socks5://127.0.0.1:1080"
func ParseProxy(proxyURL string) (res func(*http.Request) (*url.URL, error), err error) {
	var purl *url.URL
	purl, err = url.Parse(proxyURL)
	if err != nil {
		return
	}
	res = http.ProxyURL(purl)
	return
}
