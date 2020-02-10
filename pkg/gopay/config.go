package gopay

import "github.com/yearnfar/gopay/internal/gopay/adaptor"

// Config 配置
type Config struct {
	Wxpay  map[string]*adaptor.Wxpay  `toml:"wxpay"`
	Alipay map[string]*adaptor.Alipay `toml:"alipay"`
}
