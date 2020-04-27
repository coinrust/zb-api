package zb

import (
	"fmt"
	"github.com/tidwall/gjson"
	"net/url"
)

func (z *ZB) GetLeverAssetsInfo() (result LeverAssetsInfo, err error) {
	params := url.Values{}
	params.Set("method", "getLeverAssetsInfo")
	z.buildSignParams(&params)

	url := TRADE_URL + "getLeverAssetsInfo?" + params.Encode()
	var resp []byte
	resp, err = z.HttpGet(url, "", nil, nil)
	if err != nil {
		return
	}
	// {"code":3001,"message":"挂单没有找到或已完成"}
	g := gjson.ParseBytes(resp)
	if codeValue := g.Get("code"); codeValue.Exists() {
		code := codeValue.Int()
		if code != OK {
			err = fmt.Errorf("code=%v message=%v",
				codeValue.Int(), g.Get("message").String())
			return
		}
	}
	err = json.Unmarshal(resp, &result)
	return
}

// GetLoans 获取可借贷列表
// coin: qc
func (z *ZB) GetLoans(coin string, pageIndex int, pageSize int) (result []Loan, err error) {
	params := url.Values{}
	params.Set("method", "getLoans")
	params.Set("coin", coin)
	params.Set("pageIndex", fmt.Sprint(pageIndex))
	params.Set("pageSize", fmt.Sprint(pageSize))
	z.buildSignParams(&params)

	url := TRADE_URL + "getLoans?" + params.Encode()
	var ret LoansResponse
	_, err = z.HttpGet(url, "", nil, &ret)
	if err != nil {
		return
	}
	if ret.Code != OK {
		err = fmt.Errorf("code=%v message=%v",
			ret.Code, ret.Message)
		return
	}
	result = ret.Result
	return
}

// GetLoanRecords 获取借贷记录
// loanId: 理财id（借出就传 借入不用传）
// marketName: btsqc
// status: 状态(1还款中 2已还款 3 需要平仓 4 平仓还款 5自动还款中 6自动还款)
func (z *ZB) GetLoanRecords(loanId string, marketName string, status int, pageIndex int, pageSize int) (result []LoanRecord, err error) {
	params := url.Values{}
	params.Set("method", "getLoanRecords")
	if loanId != "" {
		params.Set("loanId", loanId)
	}
	params.Set("marketName", marketName)
	params.Set("pageIndex", fmt.Sprint(pageIndex))
	params.Set("pageSize", fmt.Sprint(pageSize))
	z.buildSignParams(&params)

	url := TRADE_URL + "getLoanRecords?" + params.Encode()
	var ret LoanRecordResponse
	_, err = z.HttpGet(url, "", nil, &ret)
	if err != nil {
		return
	}
	if ret.Code != OK {
		err = fmt.Errorf("code=%v message=%v",
			ret.Code, ret.Message)
		return
	}
	result = ret.Result
	return
}

// Borrow 借款
// amount: 借入金额，小数位数不能超过4，超过系统自动截取前4位
// interestRateOfDay: 日利率 [0.05-0.2] 百分比，小数位数不能超过3，超过系统自动截取前3位
// repaymentDay: 借款期限 [10/20/30]天
// isLoop: 是否自动续借 [1/0](到期自动续借/到期自动还款)
func (z *ZB) Borrow(marketName string, coin string, amount float64, interestRateOfDay float64, repaymentDay int, isLoop int) (result GeneralResult, err error) {
	params := url.Values{}
	params.Set("method", "borrow")
	params.Set("marketName", marketName)
	params.Set("coin", coin)
	params.Set("amount", fmt.Sprint(amount))
	params.Set("interestRateOfDay", fmt.Sprint(interestRateOfDay))
	params.Set("repaymentDay", fmt.Sprint(repaymentDay))
	params.Set("isLoop", fmt.Sprint(isLoop))
	z.buildSignParams(&params)

	url := TRADE_URL + "borrow?" + params.Encode()
	_, err = z.HttpGet(url, "", nil, &result)
	// {"code":1001,"message":"您的资金安全密码已经被锁定，请等待24个小时后会自动解锁。"}
	return
}

// AutoBorrow 自动借款
// amount: 借入金额，小数位数不能超过4，超过系统自动截取前4位
// interestRateOfDay: 日利率 [0.05-0.2] 百分比，小数位数不能超过3，超过系统自动截取前3位
// repaymentDay: 借款期限 [10/20/30]天
// safePwd: 资金安全密码
func (z *ZB) AutoBorrow(marketName string, coin string, amount float64, interestRateOfDay float64, repaymentDay int, safePwd string) (result GeneralResult, err error) {
	params := url.Values{}
	params.Set("method", "autoBorrow")
	params.Set("marketName", marketName)
	params.Set("coin", coin)
	params.Set("amount", fmt.Sprint(amount))
	params.Set("interestRateOfDay", fmt.Sprint(interestRateOfDay))
	params.Set("repaymentDay", fmt.Sprint(repaymentDay))
	params.Set("safePwd", safePwd)
	z.buildSignParams(&params)

	url := TRADE_URL + "autoBorrow?" + params.Encode()
	_, err = z.HttpGet(url, "", nil, &result)
	// {"code":1001,"message":"您的资金安全密码已经被锁定，请等待24个小时后会自动解锁。"}
	return
}

// Repay 还款
// loanRecordId: 借贷记录id
// repayAmount: 还款金额
// repayType: 还款方式 [0/1](全部/部分)
func (z *ZB) Repay(loanRecordId string, repayAmount float64, repayType int) (result GeneralResult, err error) {
	params := url.Values{}
	params.Set("method", "repay")
	params.Set("loanRecordId", loanRecordId)
	params.Set("repayAmount", fmt.Sprint(repayAmount))
	params.Set("repayType", fmt.Sprint(repayType))
	z.buildSignParams(&params)

	url := TRADE_URL + "repay?" + params.Encode()
	_, err = z.HttpGet(url, "", nil, &result)
	return
}

// DoAllRepay 一键还款
// marketName: 杠杆市场，示例：btcqc
func (z *ZB) DoAllRepay(marketName string) (result GeneralResult, err error) {
	params := url.Values{}
	params.Set("method", "doAllRepay")
	params.Set("marketName", marketName)
	z.buildSignParams(&params)

	url := TRADE_URL + "doAllRepay?" + params.Encode()
	_, err = z.HttpGet(url, "", nil, &result)
	return
}

// GetRepayments 获取还款记录
// loanRecordId: 借贷记录id
func (z *ZB) GetRepayments(loanRecordId string, pageIndex int, pageSize int) (result []Repayment, err error) {
	params := url.Values{}
	params.Set("method", "getRepayments")
	params.Set("loanRecordId", loanRecordId)
	params.Set("pageIndex", fmt.Sprint(pageIndex))
	params.Set("pageSize", fmt.Sprint(pageSize))
	z.buildSignParams(&params)

	url := TRADE_URL + "getRepayments?" + params.Encode()
	var ret RepaymentsResponse
	_, err = z.HttpGet(url, "", nil, &ret)
	if err != nil {
		return
	}
	result = ret.Result
	return
}
