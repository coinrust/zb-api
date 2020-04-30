package zb

import "testing"

func TestZB_DoTransferFunds(t *testing.T) {
	z := testZB()
	result, err := z.DoTransferFunds(1,
		"qc", "177xxxxxxxx", "177xxxxxxxx@c01")
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%#v", result)
}
