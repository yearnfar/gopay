package wxpay

// RefundRequest 退款请求
type RefundRequest struct {
	TransactionId string `xml:"transaction_id"`
	OutTradeNo    string `xml:"out_trade_no"`
	TotalFee      int    `xml:"total_fee"`
	OutRefundNo   string `xml:"out_refund_no"`
	RefundFee     int    `xml:"refund_fee"`
	NotifyUrl     string `xml:"notify_url,omitempty"`
}

// ToParam 参数转换
func (req *RefundRequest) ToParam(app *App) (param map[string]interface{}, err error) {
	param, err = struct2Map(req)
	if err != nil {
		return
	}

	param["app_id"] = app.AppId
	param["mch_id"] = app.MchId
	param["nonce_str"] = makeNonceStr(20)
	param["sign"] = makeSign(param, app.AppSecret)
	return
}

type RefundResponse struct {
	Response `xml:",innerXml"`

	TransactionId string `xml:"transaction_id"`
	OutTradeNo    string `xml:"out_trade_no"`
	OutRefundNo   string `xml:"out_refund_no"`
	RefundId      string `xml:"refund_id"`
	RefundFee     int    `xml:"refund_fee"`
	TotalFee      int    `xml:"total_fee"`
	CashFee       int    `xml:"cash_fee"`
}
