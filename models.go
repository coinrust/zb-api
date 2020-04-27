package zb

// Code=1000 success

type GeneralResult struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type TickerData struct {
	High float64 `json:"high,string"`
	Vol  float64 `json:"vol,string"`
	Last float64 `json:"last,string"`
	Low  float64 `json:"low,string"`
	Buy  float64 `json:"buy,string"`
	Sell float64 `json:"sell,string"`
}

type AllTicker map[string]TickerData

type Ticker struct {
	Date   int64      `json:"date,string"` // 1587952260502
	Ticker TickerData `json:"ticker"`
}

type Depth struct {
	Asks      [][]float64 `json:"asks"`
	Bids      [][]float64 `json:"bids"`
	Timestamp int64       `json:"timestamp"` // 1587953173
}

type Trade struct {
	Amount    float64 `json:"amount,string"`
	Date      int64   `json:"date"`
	Price     float64 `json:"price,string"`
	Tid       int64   `json:"tid"`
	TradeType string  `json:"trade_type"`
	Type      string  `json:"type"`
}

type KLine struct {
	Symbol    string      `json:"symbol"`
	Data      [][]float64 `json:"data"` // 时间戳(ms),Open,High,Low,Close,Volume
	MoneyType string      `json:"moneyType"`
}

type OrderResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	ID      string `json:"id"`
}

type OrderMoreV2Response struct {
	Code    int     `json:"code"`
	Message string  `json:"message"`
	Data    []int64 `json:"data"`
}

type Order struct {
	Currency    string  `json:"currency"`
	ID          string  `json:"id"`
	Price       float64 `json:"price"`
	Status      int     `json:"status"`
	TotalAmount float64 `json:"total_amount"`
	TradeAmount float64 `json:"trade_amount"`
	TradeDate   int64   `json:"trade_date"`
	TradeMoney  float64 `json:"trade_money"`
	Type        int     `json:"type"`
}

type AccountInfoCoin struct {
	IsCanWithdraw bool    `json:"isCanWithdraw"`
	CanLoan       bool    `json:"canLoan"`
	FundsType     int     `json:"fundstype"`
	ShowName      string  `json:"showName"`
	IsCanRecharge bool    `json:"isCanRecharge"`
	CnName        string  `json:"cnName"`
	EnName        string  `json:"enName"`
	Available     float64 `json:"available,string"`
	FreeZ         float64 `json:"freez,string"`
	UnitTag       string  `json:"unitTag"`
	Key           string  `json:"key"`
	UnitDecimal   int     `json:"unitDecimal"`
}

type AccountInfoBase struct {
	AuthGoogleEnabled    bool   `json:"auth_google_enabled"`
	AuthMobileEnabled    bool   `json:"auth_mobile_enabled"`
	TradePasswordEnabled bool   `json:"trade_password_enabled"`
	Username             string `json:"username"`
}

type AccountInfoResult struct {
	Coins   []AccountInfoCoin `json:"coins"`
	Version int64             `json:"version"`
	Base    AccountInfoBase   `json:"base"`
}

type AccountInfo struct {
	Result      AccountInfoResult `json:"result"`
	LeverPerm   bool              `json:"leverPerm"`
	OtcPerm     bool              `json:"otcPerm"`
	AssetPerm   bool              `json:"assetPerm"`
	MoneyPerm   bool              `json:"moneyPerm"`
	SubUserPerm bool              `json:"subUserPerm"`
	EntrustPerm bool              `json:"entrustPerm"`
}

type LeverAssetsInfoLever struct {
	TotalConvertCNY      string  `json:"totalConvertCNY"`
	LoanOutConvertUSD    string  `json:"loanOutConvertUSD"`
	UnwindPrice          string  `json:"unwindPrice"`
	FOverdraft           string  `json:"fOverdraft"`
	FShowName            string  `json:"fShowName"`
	StatusShow           string  `json:"statusShow"`
	RepayLeverShow       string  `json:"repayLeverShow"`
	CCanLoanIn           int     `json:"cCanLoanIn"`
	FLoanIn              string  `json:"fLoanIn"`
	COverdraft           string  `json:"cOverdraft"`
	NetConvertUSD        string  `json:"netConvertUSD"`
	LoanInConvertCNY     string  `json:"loanInConvertCNY"`
	Key                  string  `json:"key"`
	FUnitDecimal         int     `json:"fUnitDecimal"`
	RepayLevel           int     `json:"repayLevel"`
	ShowName             string  `json:"showName"`
	LoanInConvertUSD     string  `json:"loanInConvertUSD"`
	Level                int     `json:"level"`
	NetConvertCNY        string  `json:"netConvertCNY"`
	CFreeze              string  `json:"cFreeze"`
	CUnitTag             string  `json:"cUnitTag"`
	CouldTransferOutCoin string  `json:"couldTransferOutCoin"`
	TotalConvertUSD      string  `json:"totalConvertUSD"`
	CEnName              string  `json:"cEnName"`
	LoanOutConvertCNY    string  `json:"loanOutConvertCNY"`
	FAvailable           string  `json:"fAvailable"`
	FEnName              string  `json:"fEnName"`
	CShowName            string  `json:"cShowName"`
	LeverMultiple        string  `json:"leverMultiple"`
	CouldTransferOutFiat string  `json:"couldTransferOutFiat"`
	FFreeze              string  `json:"fFreeze"`
	FCanLoanIn           float64 `json:"fCanLoanIn"`
	CLoanOut             string  `json:"cLoanOut"`
	CUnitDecimal         float64 `json:"cUnitDecimal"`
	CLoanIn              string  `json:"cLoanIn"`
	CAvailable           string  `json:"cAvailable"`
	FLoanOut             string  `json:"fLoanOut"`
	RepayLock            bool    `json:"repayLock"`
	Status               int     `json:"status"`
}

type LeverAssetsInfoDatas struct {
	LeverPerm bool                   `json:"leverPerm"`
	Levers    []LeverAssetsInfoLever `json:"levers"`
}

type LeverAssetsInfoData struct {
	Des   string               `json:"des"`
	IsSuc bool                 `json:"isSuc"`
	Datas LeverAssetsInfoDatas `json:"datas"`
}

type LeverAssetsInfo struct {
	Code    int                 `json:"code"`
	Message LeverAssetsInfoData `json:"message"`
}

type Loan struct {
	Amount            float64 `json:"amount,string"`
	Balance           float64 `json:"balance,string"`
	CoinName          string  `json:"coinName"` // QC
	RepaymentDay      int     `json:"repaymentDay"`
	InterestRateOfDay float64 `json:"interestRateOfDay,string"`
	LowestAmount      float64 `json:"lowestAmount"`
	RateOfDayShow     string  `json:"rateOfDayShow"` // 0.05 %
}

type LoansResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Result  []Loan `json:"result"`
}

type LoanRecord struct {
	CreateTime           int64  `json:"createTime"`
	StatusShow           string `json:"statusShow"`
	FreezID              string `json:"freezId"`
	Tstatus              int    `json:"tstatus"`
	WithoutLxAmount      string `json:"withoutLxAmount"`
	InvestMark           bool   `json:"investMark"`
	WithoutLxDays        int    `json:"withoutLxDays"`
	HasRepay             string `json:"hasRepay"`
	Amount               string `json:"amount"`
	ID                   int64  `json:"id"`
	FwfScale             string `json:"fwfScale"`
	Rate                 string `json:"rate"`
	MarketName           string `json:"marketName"`
	HasLx                string `json:"hasLx"`
	IsIn                 bool   `json:"isIn"`
	BalanceAmount        string `json:"balanceAmount"`
	FundType             int    `json:"fundType"`
	OutUserID            int64  `json:"outUserId"`
	InUserID             int64  `json:"inUserId"`
	RepayDate            int64  `json:"repayDate"`
	ZheLx                string `json:"zheLx"`
	OutUserFees          string `json:"outUserFees"`
	DikouLx              string `json:"dikouLx"`
	SourceType           int    `json:"sourceType"`
	CoinName             string `json:"coinName"`
	Reward               string `json:"reward"`
	Status               int    `json:"status"`
	ArrearsLx            string `json:"arrearsLx"`
	BalanceWithoutLxDays int    `json:"balanceWithoutLxDays"`
	RiskManage           int    `json:"riskManage"`
	RateAddVal           string `json:"rateAddVal"`
	OutUserName          string `json:"outUserName"`
	InUserName           string `json:"inUserName"`
	InUserLock           bool   `json:"inUserLock"`
	RateForm             int    `json:"rateForm"`
	LoanID               int64  `json:"loanId"`
	NextRepayDate        int64  `json:"nextRepayDate"`
}

type LoanRecordResponse struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Result  []LoanRecord `json:"result"`
}

type Repayment struct {
	BenJin     string `json:"benJin"`
	ID         int64  `json:"id"`
	StatusShow string `json:"statusShow"`
	Status     int    `json:"status"`
	LiXi       string `json:"liXi"`
	ActureDate int64  `json:"actureDate"`
}

type RepaymentsResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Result  []Repayment `json:"result"`
}
