package gopay

// RefundRequest 退款参数
type RefundRequest struct {
	TransactionId string
	OutTradeNo    string
	TotalFee      int
	OutRefundNo   string
	RefundFee     int
	NotifyUrl     string
}

// RefundResponse 退款返回
type RefundResponse struct {
	TransactionId string
	OutTradeNo    string
	OutRefundNo   string
	RefundId      string
	RefundFee     int
	TotalFee      int
	CashFee       int
}
