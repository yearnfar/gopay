package wxpay

import (
	"crypto/md5"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"sort"
	"strings"

	"github.com/clbanning/mxj"
)

var (
	RandChar          = []byte("abcdefghijklmnopqrstuvwxyz0123456789")
	RandCharLen int32 = 36
)

// makeNonceStr 生成随机字符串
func makeNonceStr(n int) string {
	sb := new(strings.Builder)
	for i := 0; i < n; i++ {
		sb.WriteByte(RandChar[rand.Int31n(RandCharLen)])
	}
	return sb.String()
}

// 发送请求
func sendRequest(urlStr string, param map[string]interface{}, result interface{}) (err error) {
	data, err := xml.Marshal(param)
	if err != nil {
		return
	}

	body := strings.NewReader(string(data))
	req, err := http.NewRequest(http.MethodPost, urlStr, body)
	if err != nil {
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = xml.Unmarshal(data, result)
	if err != nil {
		return
	}
	return
}

// 生成sign
func makeSign(param map[string]interface{}, appSecret string) string {
	var keys []string
	for k := range param {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	sb := new(strings.Builder)

	for _, k := range keys {
		if v := fmt.Sprint(param[k]); k != "sign" && v != "" {
			sb.WriteString(k + "=" + v + "&")
		}
	}

	sb.WriteString("key=" + appSecret)
	return fmt.Sprintf("%X", md5.Sum([]byte(sb.String())))
}

// 数据校验
func checkSign(data []byte, appSecret string) (err error) {
	param, err := xml2Map(data)
	if err != nil {
		return
	}

	sign, ok := param["sign"]
	if !ok || sign == "" {
		return
	}

	if v := makeSign(param, appSecret); v != sign {
		return
	}
	return
}

func struct2Map(v interface{}) (xmlMap map[string]interface{}, err error) {
	data, err := xml.Marshal(v)
	if err != nil {
		return
	}

	xmlMap, err = xml2Map(data)
	if err != nil {
		return
	}
	return
}

func xml2Map(data []byte) (xmlMap map[string]interface{}, err error) {
	xm, err := mxj.NewMapXml(data)
	if err != nil {
		return
	}

	xmlMap, ok := xm["xml"].(map[string]interface{})
	if !ok {
		return
	}
	return
}
