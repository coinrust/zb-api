package zb

import "testing"

func TestZB_PlaceOrder(t *testing.T) {
	z := testZB()
	order, err := z.PlaceOrder("usdt_qc",
		1, 1, 1, 0, "")
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%#v", order)
}

func TestZB_PlaceOrderMore(t *testing.T) {
	var tradeParams [][]float64
	tradeParams = append(tradeParams, []float64{1, 1})
	tradeParams = append(tradeParams, []float64{2, 1})
	var tradeParamsData []byte
	var err error
	tradeParamsData, err = json.Marshal(tradeParams)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%v", string(tradeParamsData))
	z := testZB()
	orders, err := z.PlaceOrderMore("usdt_qc", tradeParams, 1, 0)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%#v", orders)
}

func TestZB_CancelOrder(t *testing.T) {
	z := testZB()
	order, err := z.CancelOrder("usdt_qc",
		1, "")
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%#v", order)
}

func TestZB_GetOrder(t *testing.T) {
	z := testZB()
	order, err := z.GetOrder("usdt_qc",
		1, "")
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%#v", order)
}

func TestZB_GetOrders(t *testing.T) {
	z := testZB()
	orders, err := z.GetOrders("usdt_qc",
		1, 1)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%#v", orders)
}

func TestZB_GetUnfinishedOrdersIgnoreTradeType(t *testing.T) {
	z := testZB()
	orders, err := z.GetUnfinishedOrdersIgnoreTradeType("usdt_qc",
		1, 10)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%#v", orders)
}

func TestZB_GetAccountInfo(t *testing.T) {
	z := testZB()
	accountInfo, err := z.GetAccountInfo()
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%#v", accountInfo)
}
