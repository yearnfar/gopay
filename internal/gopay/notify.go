package gopay

// NotifyResponse 回调返回
type NotifyResponse struct {
	OutTradeNo    string // 商户订单号
	TransactionId string // 支付订单号
}
