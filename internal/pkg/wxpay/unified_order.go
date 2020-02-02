package wxpay

// UnifiedOrderRequest 支付参数
type UnifiedOrderRequest struct {
	Body           string `xml:"body"`             // 商品描述
	OutTradeNo     string `xml:"out_trade_no"`     // 商户订单号
	TotalFee       int    `xml:"total_fee"`        // 标价金额
	SpbillCreateIp string `xml:"spbill_create_ip"` // 终端IP
	OpenId         string `xml:"openid"`           // 用户标识，JSAPI支付此参数必传
}

// ToParam 参数转换
func (req *UnifiedOrderRequest) ToParam(app *App) (param map[string]interface{}, err error) {
	param, err = struct2Map(req)
	if err != nil {
		return
	}

	param["appid"] = app.AppId
	param["mch_id"] = app.MchId
	param["nonce_str"] = makeNonceStr(20)
	param["trade_type"] = app.TradeType
	param["notify_url"] = app.NotifyUrl
	param["sign"] = makeSign(param, app.AppSecret)
	return
}

// UnifiedOrderResponse 支付返回
type UnifiedOrderResponse struct {
	Response `xml:",innerXml"`

	TradeType string `xml:"trade_type"` // 交易类型
	PrepayId  string `xml:"prepay_id"`  // 预支付交易会话标识
	CodeURL   string `xml:"code_url"`   // 二维码链接
}

// NotifyResponse 回调数据
type NotifyResponse struct {
	Response `xml:",innerXml"`

	OpenId        string `xml:"openid"`         // 用户在商户appid下的唯一标识
	IsSubscribe   string `xml:"is_subscribe"`   // 用户是否关注公众账号，Y-关注，N-未关注，仅在公众账号类型支付有效
	TradeType     string `xml:"trade_type"`     // 调用接口提交的交易类型，取值如下：JSAPI，NATIVE，APP，MICROPAY，详细说明见参数规定
	BankType      string `xml:"bank_type"`      // 银行类型，采用字符串类型的银行标识
	TotalFee      int    `xml:"total_fee"`      // 标价金额
	CashFee       int    `xml:"cash_fee"`       // 现金支付金额订单现金支付金额，详见支付金额
	TransactionId string `xml:"transaction_id"` // 微信支付订单号
	OutTradeNo    string `xml:"out_trade_no"`   // 商户系统内部订单号，要求32个字符内，只能是数字、大小写字母_-|*@ ，且在同一个商户号下唯一。
	TimeEnd       string `xml:"time_end"`       // 支付完成时间，格式为yyyyMMddHHmmss，如2009年12月25日9点10分10秒表示为20091225091010。
}
