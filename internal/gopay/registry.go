package gopay

import (
	"fmt"
	"sync"
)

// Registry 支付应用管理
type Registry struct {
	apps sync.Map
}

// Register 注册
func (a *Registry) Register(payType, appId string, app GoPay) (err error) {
	key := payType + "_" + appId
	_, ok := a.apps.LoadOrStore(key, app)
	if ok {
		err = fmt.Errorf("%s-%s已注册", payType, appId)
		return
	}
	return
}

// Get 获取支付
func (a *Registry) Get(payType, appId string) (app GoPay, err error) {
	key := payType + "_" + appId
	val, ok := a.apps.Load(key)
	if !ok {
		err = fmt.Errorf("%s-%s不存在", payType, appId)
		return
	}

	app = val.(GoPay)
	return
}
