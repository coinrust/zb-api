package zb

import (
	"testing"
)

func TestZB_GetLeverAssetsInfo(t *testing.T) {
	z := testZB()
	leverAssetsInfo, err := z.GetLeverAssetsInfo()
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%#v", leverAssetsInfo)
}

func TestParseLeverAssetsInfo(t *testing.T) {
	s := `{"code":1000,"message":{"des": "success", "isSuc": true, "datas": {"leverPerm":true,"levers":[{"totalConvertCNY":"0.00","loanOutConvertUSD":"0.00","unwindPrice":"0.00","fOverdraft":"0.00","fShowName":"QC","statusShow":"%E6%AD%A3%E5%B8%B8","repayLeverShow":"-","cCanLoanIn":0,"fLoanIn":"0.00","cOverdraft":"0.00","netConvertUSD":"0.00","loanInConvertCNY":"0.00","key":"btcqc","fUnitDecimal":8,"repayLevel":0,"showName":"BTC/QC","loanInConvertUSD":"0.00","level":0,"netConvertCNY":"0.00","cFreeze":"0.00","cUnitTag":"BTC","couldTransferOutCoin":"0.00","totalConvertUSD":"0.00","cEnName":"BTC","loanOutConvertCNY":"0.00","fAvailable":"0.00","fEnName":"QC","cShowName":"BTC","leverMultiple":"3.00","couldTransferOutFiat":"0.00","fFreeze":"0.00","fCanLoanIn":0,"cLoanOut":"0.00","cUnitDecimal":8,"cLoanIn":"0.00","cAvailable":"0.00","fLoanOut":"0.00","repayLock":false,"status":1},{"totalConvertCNY":"0.00","loanOutConvertUSD":"0.00","unwindPrice":"0.00","fOverdraft":"0.00","fShowName":"QC","statusShow":"%E6%AD%A3%E5%B8%B8","repayLeverShow":"-","cCanLoanIn":0,"fLoanIn":"0.00","cOverdraft":"0.00","netConvertUSD":"0.00","loanInConvertCNY":"0.00","key":"ethqc","fUnitDecimal":8,"repayLevel":0,"showName":"ETH/QC","loanInConvertUSD":"0.00","level":0,"netConvertCNY":"0.00","cFreeze":"0.00","cUnitTag":"ETH","couldTransferOutCoin":"0.00","totalConvertUSD":"0.00","cEnName":"ETH","loanOutConvertCNY":"0.00","fAvailable":"0.00","fEnName":"QC","cShowName":"ETH","leverMultiple":"3.00","couldTransferOutFiat":"0.00","fFreeze":"0.00","fCanLoanIn":0,"cLoanOut":"0.00","cUnitDecimal":8,"cLoanIn":"0.00","cAvailable":"0.00","fLoanOut":"0.00","repayLock":false,"status":1},{"totalConvertCNY":"0.00","loanOutConvertUSD":"0.00","unwindPrice":"0.00","fOverdraft":"0.00","fShowName":"QC","statusShow":"%E6%AD%A3%E5%B8%B8","repayLeverShow":"-","cCanLoanIn":0,"fLoanIn":"0.00","cOverdraft":"0.00","netConvertUSD":"0.00","loanInConvertCNY":"0.00","key":"xrpqc","fUnitDecimal":8,"repayLevel":0,"showName":"XRP/QC","loanInConvertUSD":"0.00","level":0,"netConvertCNY":"0.00","cFreeze":"0.00","cUnitTag":"XRP","couldTransferOutCoin":"0.00","totalConvertUSD":"0.00","cEnName":"XRP","loanOutConvertCNY":"0.00","fAvailable":"0.00","fEnName":"QC","cShowName":"XRP","leverMultiple":"3.00","couldTransferOutFiat":"0.00","fFreeze":"0.00","fCanLoanIn":0,"cLoanOut":"0.00","cUnitDecimal":8,"cLoanIn":"0.00","cAvailable":"0.00","fLoanOut":"0.00","repayLock":false,"status":1},{"totalConvertCNY":"0.00","loanOutConvertUSD":"0.00","unwindPrice":"0.00","fOverdraft":"0.00","fShowName":"QC","statusShow":"%E6%AD%A3%E5%B8%B8","repayLeverShow":"-","cCanLoanIn":0,"fLoanIn":"0.00","cOverdraft":"0.00","netConvertUSD":"0.00","loanInConvertCNY":"0.00","key":"eosqc","fUnitDecimal":8,"repayLevel":0,"showName":"EOS/QC","loanInConvertUSD":"0.00","level":0,"netConvertCNY":"0.00","cFreeze":"0.00","cUnitTag":"EOS","couldTransferOutCoin":"0.00","totalConvertUSD":"0.00","cEnName":"EOS","loanOutConvertCNY":"0.00","fAvailable":"0.00","fEnName":"QC","cShowName":"EOS","leverMultiple":"3.00","couldTransferOutFiat":"0.00","fFreeze":"0.00","fCanLoanIn":0,"cLoanOut":"0.00","cUnitDecimal":8,"cLoanIn":"0.00","cAvailable":"0.00","fLoanOut":"0.00","repayLock":false,"status":1},{"totalConvertCNY":"0.00","loanOutConvertUSD":"0.00","unwindPrice":"0.00","fOverdraft":"0.00","fShowName":"QC","statusShow":"%E6%AD%A3%E5%B8%B8","repayLeverShow":"-","cCanLoanIn":0,"fLoanIn":"0.00","cOverdraft":"0.00","netConvertUSD":"0.00","loanInConvertCNY":"0.00","key":"xlmqc","fUnitDecimal":8,"repayLevel":0,"showName":"XLM/QC","loanInConvertUSD":"0.00","level":0,"netConvertCNY":"0.00","cFreeze":"0.00","cUnitTag":"XLM","couldTransferOutCoin":"0.00","totalConvertUSD":"0.00","cEnName":"XLM","loanOutConvertCNY":"0.00","fAvailable":"0.00","fEnName":"QC","cShowName":"XLM","leverMultiple":"3.00","couldTransferOutFiat":"0.00","fFreeze":"0.00","fCanLoanIn":0,"cLoanOut":"0.00","cUnitDecimal":7,"cLoanIn":"0.00","cAvailable":"0.00","fLoanOut":"0.00","repayLock":false,"status":1},{"totalConvertCNY":"0.00","loanOutConvertUSD":"0.00","unwindPrice":"0.00","fOverdraft":"0.00","fShowName":"QC","statusShow":"%E6%AD%A3%E5%B8%B8","repayLeverShow":"-","cCanLoanIn":0,"fLoanIn":"0.00","cOverdraft":"0.00","netConvertUSD":"0.00","loanInConvertCNY":"0.00","key":"ltcqc","fUnitDecimal":8,"repayLevel":0,"showName":"LTC/QC","loanInConvertUSD":"0.00","level":0,"netConvertCNY":"0.00","cFreeze":"0.00","cUnitTag":"LTC","couldTransferOutCoin":"0.00","totalConvertUSD":"0.00","cEnName":"LTC","loanOutConvertCNY":"0.00","fAvailable":"0.00","fEnName":"QC","cShowName":"LTC","leverMultiple":"3.00","couldTransferOutFiat":"0.00","fFreeze":"0.00","fCanLoanIn":0,"cLoanOut":"0.00","cUnitDecimal":8,"cLoanIn":"0.00","cAvailable":"0.00","fLoanOut":"0.00","repayLock":false,"status":1},{"totalConvertCNY":"0.00","loanOutConvertUSD":"0.00","unwindPrice":"0.00","fOverdraft":"0.00","fShowName":"QC","statusShow":"%E6%AD%A3%E5%B8%B8","repayLeverShow":"-","cCanLoanIn":0,"fLoanIn":"0.00","cOverdraft":"0.00","netConvertUSD":"0.00","loanInConvertCNY":"0.00","key":"usdtqc","fUnitDecimal":8,"repayLevel":0,"showName":"USDT/QC","loanInConvertUSD":"0.00","level":0,"netConvertCNY":"0.00","cFreeze":"0.00","cUnitTag":"USDT","couldTransferOutCoin":"0.00","totalConvertUSD":"0.00","cEnName":"USDT","loanOutConvertCNY":"0.00","fAvailable":"0.00","fEnName":"QC","cShowName":"USDT","leverMultiple":"10.00","couldTransferOutFiat":"0.00","fFreeze":"0.00","fCanLoanIn":0,"cLoanOut":"0.00","cUnitDecimal":8,"cLoanIn":"0.00","cAvailable":"0.00","fLoanOut":"0.00","repayLock":false,"status":1},{"totalConvertCNY":"0.00","loanOutConvertUSD":"0.00","unwindPrice":"0.00","fOverdraft":"0.00","fShowName":"QC","statusShow":"%E6%AD%A3%E5%B8%B8","repayLeverShow":"-","cCanLoanIn":0,"fLoanIn":"0.00","cOverdraft":"0.00","netConvertUSD":"0.00","loanInConvertCNY":"0.00","key":"adaqc","fUnitDecimal":8,"repayLevel":0,"showName":"ADA/QC","loanInConvertUSD":"0.00","level":0,"netConvertCNY":"0.00","cFreeze":"0.00","cUnitTag":"ADA","couldTransferOutCoin":"0.00","totalConvertUSD":"0.00","cEnName":"ADA","loanOutConvertCNY":"0.00","fAvailable":"0.00","fEnName":"QC","cShowName":"ADA","leverMultiple":"3.00","couldTransferOutFiat":"0.00","fFreeze":"0.00","fCanLoanIn":0,"cLoanOut":"0.00","cUnitDecimal":8,"cLoanIn":"0.00","cAvailable":"0.00","fLoanOut":"0.00","repayLock":false,"status":1},{"totalConvertCNY":"0.00","loanOutConvertUSD":"0.00","unwindPrice":"0.00","fOverdraft":"0.00","fShowName":"QC","statusShow":"%E6%AD%A3%E5%B8%B8","repayLeverShow":"-","cCanLoanIn":0,"fLoanIn":"0.00","cOverdraft":"0.00","netConvertUSD":"0.00","loanInConvertCNY":"0.00","key":"dashqc","fUnitDecimal":8,"repayLevel":0,"showName":"DASH/QC","loanInConvertUSD":"0.00","level":0,"netConvertCNY":"0.00","cFreeze":"0.00","cUnitTag":"DASH","couldTransferOutCoin":"0.00","totalConvertUSD":"0.00","cEnName":"DASH","loanOutConvertCNY":"0.00","fAvailable":"0.00","fEnName":"QC","cShowName":"DASH","leverMultiple":"3.00","couldTransferOutFiat":"0.00","fFreeze":"0.00","fCanLoanIn":0,"cLoanOut":"0.00","cUnitDecimal":8,"cLoanIn":"0.00","cAvailable":"0.00","fLoanOut":"0.00","repayLock":false,"status":1},{"totalConvertCNY":"0.00","loanOutConvertUSD":"0.00","unwindPrice":"0.00","fOverdraft":"0.00","fShowName":"QC","statusShow":"%E6%AD%A3%E5%B8%B8","repayLeverShow":"-","cCanLoanIn":0,"fLoanIn":"0.00","cOverdraft":"0.00","netConvertUSD":"0.00","loanInConvertCNY":"0.00","key":"neoqc","fUnitDecimal":8,"repayLevel":0,"showName":"NEO/QC","loanInConvertUSD":"0.00","level":0,"netConvertCNY":"0.00","cFreeze":"0.00","cUnitTag":"NEO","couldTransferOutCoin":"0.00","totalConvertUSD":"0.00","cEnName":"NEO","loanOutConvertCNY":"0.00","fAvailable":"0.00","fEnName":"QC","cShowName":"NEO","leverMultiple":"3.00","couldTransferOutFiat":"0.00","fFreeze":"0.00","fCanLoanIn":0,"cLoanOut":"0.00","cUnitDecimal":8,"cLoanIn":"0.00","cAvailable":"0.00","fLoanOut":"0.00","repayLock":false,"status":1},{"totalConvertCNY":"0.00","loanOutConvertUSD":"0.00","unwindPrice":"0.00","fOverdraft":"0.00","fShowName":"QC","statusShow":"%E6%AD%A3%E5%B8%B8","repayLeverShow":"-","cCanLoanIn":0,"fLoanIn":"0.00","cOverdraft":"0.00","netConvertUSD":"0.00","loanInConvertCNY":"0.00","key":"etcqc","fUnitDecimal":8,"repayLevel":0,"showName":"ETC/QC","loanInConvertUSD":"0.00","level":0,"netConvertCNY":"0.00","cFreeze":"0.00","cUnitTag":"ETC","couldTransferOutCoin":"0.00","totalConvertUSD":"0.00","cEnName":"ETC","loanOutConvertCNY":"0.00","fAvailable":"0.00","fEnName":"QC","cShowName":"ETC","leverMultiple":"3.00","couldTransferOutFiat":"0.00","fFreeze":"0.00","fCanLoanIn":0,"cLoanOut":"0.00","cUnitDecimal":8,"cLoanIn":"0.00","cAvailable":"0.00","fLoanOut":"0.00","repayLock":false,"status":1},{"totalConvertCNY":"0.00","loanOutConvertUSD":"0.00","unwindPrice":"0.00","fOverdraft":"0.00","fShowName":"QC","statusShow":"%E6%AD%A3%E5%B8%B8","repayLeverShow":"-","cCanLoanIn":0,"fLoanIn":"0.00","cOverdraft":"0.00","netConvertUSD":"0.00","loanInConvertCNY":"0.00","key":"xemqc","fUnitDecimal":8,"repayLevel":0,"showName":"XEM/QC","loanInConvertUSD":"0.00","level":0,"netConvertCNY":"0.00","cFreeze":"0.00","cUnitTag":"XEM","couldTransferOutCoin":"0.00","totalConvertUSD":"0.00","cEnName":"XEM","loanOutConvertCNY":"0.00","fAvailable":"0.00","fEnName":"QC","cShowName":"XEM","leverMultiple":"3.00","couldTransferOutFiat":"0.00","fFreeze":"0.00","fCanLoanIn":0,"cLoanOut":"0.00","cUnitDecimal":8,"cLoanIn":"0.00","cAvailable":"0.00","fLoanOut":"0.00","repayLock":false,"status":1},{"totalConvertCNY":"0.00","loanOutConvertUSD":"0.00","unwindPrice":"0.00","fOverdraft":"0.00","fShowName":"QC","statusShow":"%E6%AD%A3%E5%B8%B8","repayLeverShow":"-","cCanLoanIn":0,"fLoanIn":"0.00","cOverdraft":"0.00","netConvertUSD":"0.00","loanInConvertCNY":"0.00","key":"dogeqc","fUnitDecimal":8,"repayLevel":0,"showName":"DOGE/QC","loanInConvertUSD":"0.00","level":0,"netConvertCNY":"0.00","cFreeze":"0.00","cUnitTag":"DOGE","couldTransferOutCoin":"0.00","totalConvertUSD":"0.00","cEnName":"DOGE","loanOutConvertCNY":"0.00","fAvailable":"0.00","fEnName":"QC","cShowName":"DOGE","leverMultiple":"3.00","couldTransferOutFiat":"0.00","fFreeze":"0.00","fCanLoanIn":0,"cLoanOut":"0.00","cUnitDecimal":8,"cLoanIn":"0.00","cAvailable":"0.00","fLoanOut":"0.00","repayLock":false,"status":1},{"totalConvertCNY":"0.00","loanOutConvertUSD":"0.00","unwindPrice":"0.00","fOverdraft":"0.00","fShowName":"QC","statusShow":"%E6%AD%A3%E5%B8%B8","repayLeverShow":"-","cCanLoanIn":0,"fLoanIn":"0.00","cOverdraft":"0.00","netConvertUSD":"0.00","loanInConvertCNY":"0.00","key":"qtumqc","fUnitDecimal":8,"repayLevel":0,"showName":"QTUM/QC","loanInConvertUSD":"0.00","level":0,"netConvertCNY":"0.00","cFreeze":"0.00","cUnitTag":"QTUM","couldTransferOutCoin":"0.00","totalConvertUSD":"0.00","cEnName":"QTUM","loanOutConvertCNY":"0.00","fAvailable":"0.00","fEnName":"QC","cShowName":"QTUM","leverMultiple":"3.00","couldTransferOutFiat":"0.00","fFreeze":"0.00","fCanLoanIn":0,"cLoanOut":"0.00","cUnitDecimal":8,"cLoanIn":"0.00","cAvailable":"0.00","fLoanOut":"0.00","repayLock":false,"status":1},{"totalConvertCNY":"0.00","loanOutConvertUSD":"0.00","unwindPrice":"0.00","fOverdraft":"0.00","fShowName":"QC","statusShow":"%E6%AD%A3%E5%B8%B8","repayLeverShow":"-","cCanLoanIn":0,"fLoanIn":"0.00","cOverdraft":"0.00","netConvertUSD":"0.00","loanInConvertCNY":"0.00","key":"btsqc","fUnitDecimal":8,"repayLevel":0,"showName":"BTS/QC","loanInConvertUSD":"0.00","level":0,"netConvertCNY":"0.00","cFreeze":"0.00","cUnitTag":"BTS","couldTransferOutCoin":"0.00","totalConvertUSD":"0.00","cEnName":"BTS","loanOutConvertCNY":"0.00","fAvailable":"0.00","fEnName":"QC","cShowName":"BTS","leverMultiple":"3.00","couldTransferOutFiat":"0.00","fFreeze":"0.00","fCanLoanIn":0,"cLoanOut":"0.00","cUnitDecimal":8,"cLoanIn":"0.00","cAvailable":"0.00","fLoanOut":"0.00","repayLock":false,"status":1},{"totalConvertCNY":"0.00","loanOutConvertUSD":"0.00","unwindPrice":"0.00","fOverdraft":"0.00","fShowName":"QC","statusShow":"%E6%AD%A3%E5%B8%B8","repayLeverShow":"-","cCanLoanIn":0,"fLoanIn":"0.00","cOverdraft":"0.00","netConvertUSD":"0.00","loanInConvertCNY":"0.00","key":"hsrqc","fUnitDecimal":8,"repayLevel":0,"showName":"HC/QC","loanInConvertUSD":"0.00","level":0,"netConvertCNY":"0.00","cFreeze":"0.00","cUnitTag":"HC","couldTransferOutCoin":"0.00","totalConvertUSD":"0.00","cEnName":"HSR","loanOutConvertCNY":"0.00","fAvailable":"0.00","fEnName":"QC","cShowName":"HC","leverMultiple":"3.00","couldTransferOutFiat":"0.00","fFreeze":"0.00","fCanLoanIn":0,"cLoanOut":"0.00","cUnitDecimal":8,"cLoanIn":"0.00","cAvailable":"0.00","fLoanOut":"0.00","repayLock":false,"status":1},{"totalConvertCNY":"0.00","loanOutConvertUSD":"0.00","unwindPrice":"0.00","fOverdraft":"0.00","fShowName":"QC","statusShow":"%E6%AD%A3%E5%B8%B8","repayLeverShow":"-","cCanLoanIn":0,"fLoanIn":"0.00","cOverdraft":"0.00","netConvertUSD":"0.00","loanInConvertCNY":"0.00","key":"btmqc","fUnitDecimal":8,"repayLevel":0,"showName":"BTM/QC","loanInConvertUSD":"0.00","level":0,"netConvertCNY":"0.00","cFreeze":"0.00","cUnitTag":"BTM","couldTransferOutCoin":"0.00","totalConvertUSD":"0.00","cEnName":"BTM","loanOutConvertCNY":"0.00","fAvailable":"0.00","fEnName":"QC","cShowName":"BTM","leverMultiple":"3.00","couldTransferOutFiat":"0.00","fFreeze":"0.00","fCanLoanIn":0,"cLoanOut":"0.00","cUnitDecimal":8,"cLoanIn":"0.00","cAvailable":"0.00","fLoanOut":"0.00","repayLock":false,"status":1},{"totalConvertCNY":"0.00","loanOutConvertUSD":"0.00","unwindPrice":"0.00","fOverdraft":"0.00","fShowName":"QC","statusShow":"%E6%AD%A3%E5%B8%B8","repayLeverShow":"-","cCanLoanIn":0,"fLoanIn":"0.00","cOverdraft":"0.00","netConvertUSD":"0.00","loanInConvertCNY":"0.00","key":"bchsvqc","fUnitDecimal":8,"repayLevel":0,"showName":"BSV/QC","loanInConvertUSD":"0.00","level":0,"netConvertCNY":"0.00","cFreeze":"0.00","cUnitTag":"BSV","couldTransferOutCoin":"0.00","totalConvertUSD":"0.00","cEnName":"BCHSV","loanOutConvertCNY":"0.00","fAvailable":"0.00","fEnName":"QC","cShowName":"BSV","leverMultiple":"3.00","couldTransferOutFiat":"0.00","fFreeze":"0.00","fCanLoanIn":0,"cLoanOut":"0.00","cUnitDecimal":8,"cLoanIn":"0.00","cAvailable":"0.00","fLoanOut":"0.00","repayLock":false,"status":1},{"totalConvertCNY":"0.00","loanOutConvertUSD":"0.00","unwindPrice":"0.00","fOverdraft":"0.00","fShowName":"USDT","statusShow":"%E6%AD%A3%E5%B8%B8","repayLeverShow":"-","cCanLoanIn":0,"fLoanIn":"0.00","cOverdraft":"0.00","netConvertUSD":"0.00","loanInConvertCNY":"0.00","key":"btcusdt","fUnitDecimal":8,"repayLevel":0,"showName":"BTC/USDT","loanInConvertUSD":"0.00","level":0,"netConvertCNY":"0.00","cFreeze":"0.00","cUnitTag":"BTC","couldTransferOutCoin":"0.00","totalConvertUSD":"0.00","cEnName":"BTC","loanOutConvertCNY":"0.00","fAvailable":"0.00","fEnName":"USDT","cShowName":"BTC","leverMultiple":"3.00","couldTransferOutFiat":"0.00","fFreeze":"0.00","fCanLoanIn":0,"cLoanOut":"0.00","cUnitDecimal":8,"cLoanIn":"0.00","cAvailable":"0.00","fLoanOut":"0.00","repayLock":false,"status":1},{"totalConvertCNY":"0.00","loanOutConvertUSD":"0.00","unwindPrice":"0.00","fOverdraft":"0.00","fShowName":"USDT","statusShow":"%E6%AD%A3%E5%B8%B8","repayLeverShow":"-","cCanLoanIn":0,"fLoanIn":"0.00","cOverdraft":"0.00","netConvertUSD":"0.00","loanInConvertCNY":"0.00","key":"ethusdt","fUnitDecimal":8,"repayLevel":0,"showName":"ETH/USDT","loanInConvertUSD":"0.00","level":0,"netConvertCNY":"0.00","cFreeze":"0.00","cUnitTag":"ETH","couldTransferOutCoin":"0.00","totalConvertUSD":"0.00","cEnName":"ETH","loanOutConvertCNY":"0.00","fAvailable":"0.00","fEnName":"USDT","cShowName":"ETH","leverMultiple":"3.00","couldTransferOutFiat":"0.00","fFreeze":"0.00","fCanLoanIn":0,"cLoanOut":"0.00","cUnitDecimal":8,"cLoanIn":"0.00","cAvailable":"0.00","fLoanOut":"0.00","repayLock":false,"status":1},{"totalConvertCNY":"0.00","loanOutConvertUSD":"0.00","unwindPrice":"0.00","fOverdraft":"0.00","fShowName":"USDT","statusShow":"%E6%AD%A3%E5%B8%B8","repayLeverShow":"-","cCanLoanIn":0,"fLoanIn":"0.00","cOverdraft":"0.00","netConvertUSD":"0.00","loanInConvertCNY":"0.00","key":"xrpusdt","fUnitDecimal":8,"repayLevel":0,"showName":"XRP/USDT","loanInConvertUSD":"0.00","level":0,"netConvertCNY":"0.00","cFreeze":"0.00","cUnitTag":"XRP","couldTransferOutCoin":"0.00","totalConvertUSD":"0.00","cEnName":"XRP","loanOutConvertCNY":"0.00","fAvailable":"0.00","fEnName":"USDT","cShowName":"XRP","leverMultiple":"3.00","couldTransferOutFiat":"0.00","fFreeze":"0.00","fCanLoanIn":0,"cLoanOut":"0.00","cUnitDecimal":8,"cLoanIn":"0.00","cAvailable":"0.00","fLoanOut":"0.00","repayLock":false,"status":1},{"totalConvertCNY":"0.00","loanOutConvertUSD":"0.00","unwindPrice":"0.00","fOverdraft":"0.00","fShowName":"USDT","statusShow":"%E6%AD%A3%E5%B8%B8","repayLeverShow":"-","cCanLoanIn":0,"fLoanIn":"0.00","cOverdraft":"0.00","netConvertUSD":"0.00","loanInConvertCNY":"0.00","key":"eosusdt","fUnitDecimal":8,"repayLevel":0,"showName":"EOS/USDT","loanInConvertUSD":"0.00","level":0,"netConvertCNY":"0.00","cFreeze":"0.00","cUnitTag":"EOS","couldTransferOutCoin":"0.00","totalConvertUSD":"0.00","cEnName":"EOS","loanOutConvertCNY":"0.00","fAvailable":"0.00","fEnName":"USDT","cShowName":"EOS","leverMultiple":"3.00","couldTransferOutFiat":"0.00","fFreeze":"0.00","fCanLoanIn":0,"cLoanOut":"0.00","cUnitDecimal":8,"cLoanIn":"0.00","cAvailable":"0.00","fLoanOut":"0.00","repayLock":false,"status":1},{"totalConvertCNY":"0.00","loanOutConvertUSD":"0.00","unwindPrice":"0.00","fOverdraft":"0.00","fShowName":"USDT","statusShow":"%E6%AD%A3%E5%B8%B8","repayLeverShow":"-","cCanLoanIn":0,"fLoanIn":"0.00","cOverdraft":"0.00","netConvertUSD":"0.00","loanInConvertCNY":"0.00","key":"xlmusdt","fUnitDecimal":8,"repayLevel":0,"showName":"XLM/USDT","loanInConvertUSD":"0.00","level":0,"netConvertCNY":"0.00","cFreeze":"0.00","cUnitTag":"XLM","couldTransferOutCoin":"0.00","totalConvertUSD":"0.00","cEnName":"XLM","loanOutConvertCNY":"0.00","fAvailable":"0.00","fEnName":"USDT","cShowName":"XLM","leverMultiple":"3.00","couldTransferOutFiat":"0.00","fFreeze":"0.00","fCanLoanIn":0,"cLoanOut":"0.00","cUnitDecimal":7,"cLoanIn":"0.00","cAvailable":"0.00","fLoanOut":"0.00","repayLock":false,"status":1},{"totalConvertCNY":"0.00","loanOutConvertUSD":"0.00","unwindPrice":"0.00","fOverdraft":"0.00","fShowName":"USDT","statusShow":"%E6%AD%A3%E5%B8%B8","repayLeverShow":"-","cCanLoanIn":0,"fLoanIn":"0.00","cOverdraft":"0.00","netConvertUSD":"0.00","loanInConvertCNY":"0.00","key":"ltcusdt","fUnitDecimal":8,"repayLevel":0,"showName":"LTC/USDT","loanInConvertUSD":"0.00","level":0,"netConvertCNY":"0.00","cFreeze":"0.00","cUnitTag":"LTC","couldTransferOutCoin":"0.00","totalConvertUSD":"0.00","cEnName":"LTC","loanOutConvertCNY":"0.00","fAvailable":"0.00","fEnName":"USDT","cShowName":"LTC","leverMultiple":"3.00","couldTransferOutFiat":"0.00","fFreeze":"0.00","fCanLoanIn":0,"cLoanOut":"0.00","cUnitDecimal":8,"cLoanIn":"0.00","cAvailable":"0.00","fLoanOut":"0.00","repayLock":false,"status":1},{"totalConvertCNY":"0.00","loanOutConvertUSD":"0.00","unwindPrice":"0.00","fOverdraft":"0.00","fShowName":"USDT","statusShow":"%E6%AD%A3%E5%B8%B8","repayLeverShow":"-","cCanLoanIn":0,"fLoanIn":"0.00","cOverdraft":"0.00","netConvertUSD":"0.00","loanInConvertCNY":"0.00","key":"adausdt","fUnitDecimal":8,"repayLevel":0,"showName":"ADA/USDT","loanInConvertUSD":"0.00","level":0,"netConvertCNY":"0.00","cFreeze":"0.00","cUnitTag":"ADA","couldTransferOutCoin":"0.00","totalConvertUSD":"0.00","cEnName":"ADA","loanOutConvertCNY":"0.00","fAvailable":"0.00","fEnName":"USDT","cShowName":"ADA","leverMultiple":"3.00","couldTransferOutFiat":"0.00","fFreeze":"0.00","fCanLoanIn":0,"cLoanOut":"0.00","cUnitDecimal":8,"cLoanIn":"0.00","cAvailable":"0.00","fLoanOut":"0.00","repayLock":false,"status":1},{"totalConvertCNY":"0.00","loanOutConvertUSD":"0.00","unwindPrice":"0.00","fOverdraft":"0.00","fShowName":"USDT","statusShow":"%E6%AD%A3%E5%B8%B8","repayLeverShow":"-","cCanLoanIn":0,"fLoanIn":"0.00","cOverdraft":"0.00","netConvertUSD":"0.00","loanInConvertCNY":"0.00","key":"dashusdt","fUnitDecimal":8,"repayLevel":0,"showName":"DASH/USDT","loanInConvertUSD":"0.00","level":0,"netConvertCNY":"0.00","cFreeze":"0.00","cUnitTag":"DASH","couldTransferOutCoin":"0.00","totalConvertUSD":"0.00","cEnName":"DASH","loanOutConvertCNY":"0.00","fAvailable":"0.00","fEnName":"USDT","cShowName":"DASH","leverMultiple":"3.00","couldTransferOutFiat":"0.00","fFreeze":"0.00","fCanLoanIn":0,"cLoanOut":"0.00","cUnitDecimal":8,"cLoanIn":"0.00","cAvailable":"0.00","fLoanOut":"0.00","repayLock":false,"status":1},{"totalConvertCNY":"0.00","loanOutConvertUSD":"0.00","unwindPrice":"0.00","fOverdraft":"0.00","fShowName":"USDT","statusShow":"%E6%AD%A3%E5%B8%B8","repayLeverShow":"-","cCanLoanIn":0,"fLoanIn":"0.00","cOverdraft":"0.00","netConvertUSD":"0.00","loanInConvertCNY":"0.00","key":"neousdt","fUnitDecimal":8,"repayLevel":0,"showName":"NEO/USDT","loanInConvertUSD":"0.00","level":0,"netConvertCNY":"0.00","cFreeze":"0.00","cUnitTag":"NEO","couldTransferOutCoin":"0.00","totalConvertUSD":"0.00","cEnName":"NEO","loanOutConvertCNY":"0.00","fAvailable":"0.00","fEnName":"USDT","cShowName":"NEO","leverMultiple":"3.00","couldTransferOutFiat":"0.00","fFreeze":"0.00","fCanLoanIn":0,"cLoanOut":"0.00","cUnitDecimal":8,"cLoanIn":"0.00","cAvailable":"0.00","fLoanOut":"0.00","repayLock":false,"status":1},{"totalConvertCNY":"0.00","loanOutConvertUSD":"0.00","unwindPrice":"0.00","fOverdraft":"0.00","fShowName":"USDT","statusShow":"%E6%AD%A3%E5%B8%B8","repayLeverShow":"-","cCanLoanIn":0,"fLoanIn":"0.00","cOverdraft":"0.00","netConvertUSD":"0.00","loanInConvertCNY":"0.00","key":"etcusdt","fUnitDecimal":8,"repayLevel":0,"showName":"ETC/USDT","loanInConvertUSD":"0.00","level":0,"netConvertCNY":"0.00","cFreeze":"0.00","cUnitTag":"ETC","couldTransferOutCoin":"0.00","totalConvertUSD":"0.00","cEnName":"ETC","loanOutConvertCNY":"0.00","fAvailable":"0.00","fEnName":"USDT","cShowName":"ETC","leverMultiple":"3.00","couldTransferOutFiat":"0.00","fFreeze":"0.00","fCanLoanIn":0,"cLoanOut":"0.00","cUnitDecimal":8,"cLoanIn":"0.00","cAvailable":"0.00","fLoanOut":"0.00","repayLock":false,"status":1},{"totalConvertCNY":"0.00","loanOutConvertUSD":"0.00","unwindPrice":"0.00","fOverdraft":"0.00","fShowName":"USDT","statusShow":"%E6%AD%A3%E5%B8%B8","repayLeverShow":"-","cCanLoanIn":0,"fLoanIn":"0.00","cOverdraft":"0.00","netConvertUSD":"0.00","loanInConvertCNY":"0.00","key":"xemusdt","fUnitDecimal":8,"repayLevel":0,"showName":"XEM/USDT","loanInConvertUSD":"0.00","level":0,"netConvertCNY":"0.00","cFreeze":"0.00","cUnitTag":"XEM","couldTransferOutCoin":"0.00","totalConvertUSD":"0.00","cEnName":"XEM","loanOutConvertCNY":"0.00","fAvailable":"0.00","fEnName":"USDT","cShowName":"XEM","leverMultiple":"3.00","couldTransferOutFiat":"0.00","fFreeze":"0.00","fCanLoanIn":0,"cLoanOut":"0.00","cUnitDecimal":8,"cLoanIn":"0.00","cAvailable":"0.00","fLoanOut":"0.00","repayLock":false,"status":1},{"totalConvertCNY":"0.00","loanOutConvertUSD":"0.00","unwindPrice":"0.00","fOverdraft":"0.00","fShowName":"USDT","statusShow":"%E6%AD%A3%E5%B8%B8","repayLeverShow":"-","cCanLoanIn":0,"fLoanIn":"0.00","cOverdraft":"0.00","netConvertUSD":"0.00","loanInConvertCNY":"0.00","key":"dogeusdt","fUnitDecimal":8,"repayLevel":0,"showName":"DOGE/USDT","loanInConvertUSD":"0.00","level":0,"netConvertCNY":"0.00","cFreeze":"0.00","cUnitTag":"DOGE","couldTransferOutCoin":"0.00","totalConvertUSD":"0.00","cEnName":"DOGE","loanOutConvertCNY":"0.00","fAvailable":"0.00","fEnName":"USDT","cShowName":"DOGE","leverMultiple":"3.00","couldTransferOutFiat":"0.00","fFreeze":"0.00","fCanLoanIn":0,"cLoanOut":"0.00","cUnitDecimal":8,"cLoanIn":"0.00","cAvailable":"0.00","fLoanOut":"0.00","repayLock":false,"status":1},{"totalConvertCNY":"0.00","loanOutConvertUSD":"0.00","unwindPrice":"0.00","fOverdraft":"0.00","fShowName":"USDT","statusShow":"%E6%AD%A3%E5%B8%B8","repayLeverShow":"-","cCanLoanIn":0,"fLoanIn":"0.00","cOverdraft":"0.00","netConvertUSD":"0.00","loanInConvertCNY":"0.00","key":"qtumusdt","fUnitDecimal":8,"repayLevel":0,"showName":"QTUM/USDT","loanInConvertUSD":"0.00","level":0,"netConvertCNY":"0.00","cFreeze":"0.00","cUnitTag":"QTUM","couldTransferOutCoin":"0.00","totalConvertUSD":"0.00","cEnName":"QTUM","loanOutConvertCNY":"0.00","fAvailable":"0.00","fEnName":"USDT","cShowName":"QTUM","leverMultiple":"3.00","couldTransferOutFiat":"0.00","fFreeze":"0.00","fCanLoanIn":0,"cLoanOut":"0.00","cUnitDecimal":8,"cLoanIn":"0.00","cAvailable":"0.00","fLoanOut":"0.00","repayLock":false,"status":1},{"totalConvertCNY":"0.00","loanOutConvertUSD":"0.00","unwindPrice":"0.00","fOverdraft":"0.00","fShowName":"USDT","statusShow":"%E6%AD%A3%E5%B8%B8","repayLeverShow":"-","cCanLoanIn":0,"fLoanIn":"0.00","cOverdraft":"0.00","netConvertUSD":"0.00","loanInConvertCNY":"0.00","key":"btsusdt","fUnitDecimal":8,"repayLevel":0,"showName":"BTS/USDT","loanInConvertUSD":"0.00","level":0,"netConvertCNY":"0.00","cFreeze":"0.00","cUnitTag":"BTS","couldTransferOutCoin":"0.00","totalConvertUSD":"0.00","cEnName":"BTS","loanOutConvertCNY":"0.00","fAvailable":"0.00","fEnName":"USDT","cShowName":"BTS","leverMultiple":"3.00","couldTransferOutFiat":"0.00","fFreeze":"0.00","fCanLoanIn":0,"cLoanOut":"0.00","cUnitDecimal":8,"cLoanIn":"0.00","cAvailable":"0.00","fLoanOut":"0.00","repayLock":false,"status":1},{"totalConvertCNY":"0.00","loanOutConvertUSD":"0.00","unwindPrice":"0.00","fOverdraft":"0.00","fShowName":"USDT","statusShow":"%E6%AD%A3%E5%B8%B8","repayLeverShow":"-","cCanLoanIn":0,"fLoanIn":"0.00","cOverdraft":"0.00","netConvertUSD":"0.00","loanInConvertCNY":"0.00","key":"hsrusdt","fUnitDecimal":8,"repayLevel":0,"showName":"HC/USDT","loanInConvertUSD":"0.00","level":0,"netConvertCNY":"0.00","cFreeze":"0.00","cUnitTag":"HC","couldTransferOutCoin":"0.00","totalConvertUSD":"0.00","cEnName":"HSR","loanOutConvertCNY":"0.00","fAvailable":"0.00","fEnName":"USDT","cShowName":"HC","leverMultiple":"3.00","couldTransferOutFiat":"0.00","fFreeze":"0.00","fCanLoanIn":0,"cLoanOut":"0.00","cUnitDecimal":8,"cLoanIn":"0.00","cAvailable":"0.00","fLoanOut":"0.00","repayLock":false,"status":1},{"totalConvertCNY":"0.00","loanOutConvertUSD":"0.00","unwindPrice":"0.00","fOverdraft":"0.00","fShowName":"USDT","statusShow":"%E6%AD%A3%E5%B8%B8","repayLeverShow":"-","cCanLoanIn":0,"fLoanIn":"0.00","cOverdraft":"0.00","netConvertUSD":"0.00","loanInConvertCNY":"0.00","key":"btmusdt","fUnitDecimal":8,"repayLevel":0,"showName":"BTM/USDT","loanInConvertUSD":"0.00","level":0,"netConvertCNY":"0.00","cFreeze":"0.00","cUnitTag":"BTM","couldTransferOutCoin":"0.00","totalConvertUSD":"0.00","cEnName":"BTM","loanOutConvertCNY":"0.00","fAvailable":"0.00","fEnName":"USDT","cShowName":"BTM","leverMultiple":"3.00","couldTransferOutFiat":"0.00","fFreeze":"0.00","fCanLoanIn":0,"cLoanOut":"0.00","cUnitDecimal":8,"cLoanIn":"0.00","cAvailable":"0.00","fLoanOut":"0.00","repayLock":false,"status":1}]}}}`
	var ret LeverAssetsInfo
	err := json.Unmarshal([]byte(s), &ret)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%#v", ret)
}

func TestZB_GetLoans(t *testing.T) {
	z := testZB()
	loans, err := z.GetLoans("qc", 1, 10)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%#v", loans)
}

func TestZB_GetLoanRecords(t *testing.T) {
	z := testZB()
	loanRecords, err := z.GetLoanRecords("",
		"btsqc", 1, 1, 10)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%#v", loanRecords)
}

func TestZB_Borrow(t *testing.T) {
	z := testZB()
	result, err := z.Borrow("btcqc",
		"qc", 100, 0.1, 30, 0)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%v", result)
}