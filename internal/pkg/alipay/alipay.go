package alipay

import (
	"encoding/json"
	"time"
)

const (
	Gateway = "https://openapi.alipay.com/gateway.do"
)

const (
	MethodAppPay = "alipay.trade.app.pay" // app支付接口
	MethodWapPay = "alipay.trade.wap.pay" // 手机网站支付接口
	MethodPay    = "alipay.trade.pay"     // 统一收单交易支付接口
)

type App struct {
	AppId     string
	AppSecret string
	Method    string
}

// UnifiedOrder 统一下单
func (a *App) UnifiedOrder(req *UnifiedOrderRequest) (resp *UnifiedOrderResponse, err error) {
	param, err := a.setParam(req)
	if err != nil {
		return
	}

	resp = &UnifiedOrderResponse{}
	err = sendRequest(Gateway, param, resp)
	if err != nil {
		return
	}
	return
}

// Refund 退款
func (a *App) Refund(req *RefundRequest) (resp *RefundResponse, err error) {
	param, err := a.setParam(req)
	if err != nil {
		return
	}

	resp = &RefundResponse{}
	err = sendRequest(Gateway, param, resp)
	if err != nil {
		return
	}
	return
}

// Notify 异步回调
func (a *App) Notify(body []byte) (resp *NotifyResponse, err error) {
	return
}

func (a *App) setParam(v interface{}) (param map[string]interface{}, err error) {
	bizContent, err := a.setBizContent(v)
	if err != nil {
		return
	}

	param = make(map[string]interface{})
	param["app_id"] = a.AppId
	param["method"] = a.Method
	param["charset"] = "utf-8"
	param["sign_type"] = "MD5"
	param["timestamp"] = time.Now().Format("2006-01-02 15:04:05")
	param["biz_content"] = bizContent
	param["sign"] = makeSign(param, a.AppSecret)
	return
}

func (a *App) setBizContent(v interface{}) (string, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
