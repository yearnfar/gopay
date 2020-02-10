package alipay

type UnifiedOrderRequest struct {
	TotalAccount   float64 `json:"total_account"`   // 商户网站唯一订单号
	Subject        string  `json:"subject"`         // 商品的标题/交易标题/订单标题/订单关键字等。
	OutTradeNo     string  `json:"out_trade_no"`    // 订单总金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000]
	TimeoutExpress string  `json:"timeout_express"` // 该笔订单允许的最晚付款时间，逾期将关闭交易。取值范围：1m～15d。m-分钟，h-小时，d-天，1c-当天（1c-当天的情况下，无论交易何时创建，都在0点关闭）。 该参数数值不接受小数点， 如 1.5h，可转换为 90m
}

// Valid 参数校验
func (req *UnifiedOrderRequest) Valid(app *App) (err error) {
	return
}

// UnifiedOrderResponse 下单返回数据
type UnifiedOrderResponse struct {
}

// NotifyResponse 异步回调
type NotifyResponse struct {
}
