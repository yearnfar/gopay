package wxpay

import "encoding/xml"

// Response 接口返回
type Response struct {
	XMLName    xml.Name `xml:"xml"`
	ReturnCode string   `xml:"return_code"`
	ReturnMsg  string   `xml:"return_msg"`
	AppId      string   `xml:"appid"`        // 公众账号ID
	MchId      string   `xml:"mch_id"`       // 商户号
	DeviceInfo string   `xml:"device_info"`  // 设备号
	NonceStr   string   `xml:"nonce_str"`    // 随机字符串
	Sign       string   `xml:"sign"`         // 签名
	SignType   string   `xml:"sign_type"`    // 签名类型
	ResultCode string   `xml:"result_code"`  // 业务结果
	ErrCode    string   `xml:"err_code"`     // 错误码
	ErrCodeDes string   `xml:"err_code_des"` // 错误描述
}
