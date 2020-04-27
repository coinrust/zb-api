package zb

import "testing"

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
