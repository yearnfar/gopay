package adaptor

import (
	"github.com/yearnfar/gopay/internal/gopay"
	"github.com/yearnfar/gopay/internal/pkg/wxpay"
)

// Wxpay 微信支付
type Wxpay struct {
	AppId     string // 公众账号ID
	AppSecret string // AppSecret是APPID对应的接口密码
	MchId     string // 商户号
	NotifyUrl string // 通知地址
	TradeType string // 交易类型
}

var _ gopay.GoPay = &Wxpay{}

// UnifiedOrder 统一下单接口
func (p *Wxpay) UnifiedOrder(req *gopay.UnifiedOrderRequest) (resp *gopay.UnifiedOrderResponse, err error) {
	param := &wxpay.UnifiedOrderRequest{
		Body:       req.Subject,
		OutTradeNo: req.OutTradeNo,
		TotalFee:   req.TotalFee,
	}

	if req.OpenId != "" {
		param.OpenId = req.OpenId
	}

	if req.ClientIp != "" {
		param.SpbillCreateIp = req.ClientIp
	}

	app := wxpay.NewApp(p.TradeType, p.AppId, p.AppSecret, p.MchId, p.NotifyUrl)
	result, err := app.UnifiedOrder(param)
	if err != nil {
		return
	}

	resp = &gopay.UnifiedOrderResponse{
		TradeType: result.TradeType,
		PrepayId:  result.PrepayId,
		CodeUrl:   result.CodeURL,
	}
	return
}

// Notify 支付回调
func (p *Wxpay) Notify(data []byte) (resp *gopay.NotifyResponse, err error) {
	app := wxpay.NewApp(p.TradeType, p.AppId, p.AppSecret, p.MchId, p.NotifyUrl)
	result, err := app.Notify(data)
	if err != nil {
		return
	}

	resp = &gopay.NotifyResponse{
		OutTradeNo:    result.OutTradeNo,
		TransactionId: result.TransactionId,
	}
	return
}

// Refund 退款接口
func (p *Wxpay) Refund(req *gopay.RefundRequest) (resp *gopay.RefundResponse, err error) {
	param := &wxpay.RefundRequest{
		TransactionId: req.TransactionId,
		OutTradeNo:    req.OutTradeNo,
		OutRefundNo:   req.OutRefundNo,
		TotalFee:      req.TotalFee,
		RefundFee:     req.RefundFee,
		NotifyUrl:     req.NotifyUrl,
	}

	app := wxpay.NewApp(p.TradeType, p.AppId, p.AppSecret, p.MchId, p.NotifyUrl)
	result, err := app.Refund(param)
	if err != nil {
		return
	}

	resp = &gopay.RefundResponse{
		TransactionId: result.TransactionId,
		OutTradeNo:    result.OutTradeNo,
		OutRefundNo:   result.OutRefundNo,
		TotalFee:      result.TotalFee,
		RefundId:      result.RefundId,
		RefundFee:     result.RefundFee,
		CashFee:       result.CashFee,
	}
	return
}
