package zb

import (
	"fmt"
	"net/url"
)

// DoTransferFunds 主子账号内部转账
// amount: 金额
// currency: 币种
// fromUserName: 转出方用户名
// toUserName: 转入方用户名
func (z *ZB) DoTransferFunds(amount float64, currency string, fromUserName string, toUserName string) (result GeneralResult, err error) {
	params := url.Values{}
	params.Set("amount", fmt.Sprint(amount))
	params.Set("currency", currency)
	params.Set("fromUserName", fromUserName)
	params.Set("toUserName", toUserName)
	z.buildSignParams(&params)

	url := z.tradeURL + "doTransferFunds?" + params.Encode()
	var ret GeneralResult
	_, err = z.HttpGet(url, "", nil, &ret)
	if err != nil {
		return
	}
	if ret.Code != OK {
		err = fmt.Errorf("code=%v message=%v", ret.Code, ret.Message)
		return
	}
	return
}
