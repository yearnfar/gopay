package wxpay

import (
	"encoding/xml"
)

const (
	// UnifiedOrderApi 统一下单接口
	UnifiedOrderApi = "https://api.mch.weixin.qq.com/adaptor/unifiedorder"

	RefundApi = "https://api.mch.weixin.qq.com/secapi/pay/refund"
)

const (
	// TradeTypeJSAPI 公众号支付
	TradeTypeJSAPI = "JSAPI"

	// TradeTypeNative 扫码支付
	TradeTypeNative = "NATIVE"

	// TradeTypeApp App支付
	TradeTypeApp = "APP"
)

// App 微信支付
type App struct {
	AppId     string // 公众账号ID
	AppSecret string // AppSecret是APPID对应的接口密码
	MchId     string // 商户号
	NotifyUrl string // 通知地址
	TradeType string // 交易类型
}

func NewApp(tradeType, appId, appSecret, mchId, notifyUrl string) *App {
	return &App{
		AppId:     appId,
		AppSecret: appSecret,
		MchId:     mchId,
		NotifyUrl: notifyUrl,
		TradeType: tradeType,
	}
}

// UnifiedOrder 统一下单接口
func (a *App) UnifiedOrder(req *UnifiedOrderRequest) (result *UnifiedOrderResponse, err error) {
	param, err := req.ToParam(a)
	if err != nil {
		return
	}

	result = &UnifiedOrderResponse{}
	err = sendRequest(UnifiedOrderApi, param, result)
	if err != nil {
		return
	}
	return
}

// Notify 支付回调
func (a *App) Notify(data []byte) (result *NotifyResponse, err error) {
	err = checkSign(data, a.AppSecret)
	if err != nil {
		return
	}

	result = &NotifyResponse{}
	err = xml.Unmarshal(data, result)
	if err != nil {
		return
	}
	return
}

// Refund 退款接口
func (a *App) Refund(req *RefundRequest) (result *RefundResponse, err error) {
	param, err := req.ToParam(a)
	if err != nil {
		return
	}

	result = &RefundResponse{}
	err = sendRequest(UnifiedOrderApi, param, result)
	if err != nil {
		return
	}
	return
}
