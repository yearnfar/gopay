package gopay

import (
	"github.com/BurntSushi/toml"
	"github.com/yearnfar/gopay/internal/gopay"
	"github.com/yearnfar/gopay/internal/gopay/adaptor"
)

var registry = &gopay.Registry{}

// LoadConfig加载配置
func LoadConfig(cfgFile string) {
	cfg := &Config{}
	_, err := toml.DecodeFile(cfgFile, cfg)
	if err != nil {
		return
	}

	for _, c := range cfg.Wxpay {
		app := &adaptor.Wxpay{
			MchId:     c.MchId,
			AppId:     c.AppId,
			AppSecret: c.AppSecret,
			NotifyUrl: c.NotifyUrl,
			TradeType: c.TradeType,
		}

		registry.Register(gopay.WxPay, app.TradeType, app)
	}
	return
}

func LoadAliPay(payType string, config map[string]*adaptor.Alipay) {

}

// UnifiedOrder 统一下单
func UnifiedOrder() {

}
