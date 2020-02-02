package gopay

// GoPay 支付接口
type GoPay interface {
	UnifiedOrder(*UnifiedOrderRequest) (*UnifiedOrderResponse, error)
	Notify([]byte) (*NotifyResponse, error)
	Refund(*RefundRequest) (*RefundResponse, error)
}
