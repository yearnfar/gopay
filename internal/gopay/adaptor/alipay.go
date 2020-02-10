package adaptor

type Alipay struct {
	AppId     string // 公众账号ID
	AppSecret string // AppSecret是APPID对应的接口密码
	NotifyUrl string // 通知地址
	TradeType string // 交易类型
}
