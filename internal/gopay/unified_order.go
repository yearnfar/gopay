package gopay

type UnifiedOrderRequest struct {
	Subject    string
	OutTradeNo string
	TotalFee   int
	OpenId     string
	ClientIp   string
}

type UnifiedOrderResponse struct {
	TradeType string //
	PrepayId  string //
	CodeUrl   string //
}
