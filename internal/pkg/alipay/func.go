package alipay

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
)

func sendRequest(urlStr string, param map[string]interface{}, result interface{}) (err error) {
	data, err := json.Marshal(param)
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

	err = json.Unmarshal(data, result)
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
		if v := fmt.Sprint(param[k]); k != "sign" && k != "sign_type" && v != "" {
			sb.WriteString(k + "=" + v + "&")
		}
	}

	sb.WriteString("key=" + appSecret)
	return fmt.Sprintf("%X", md5.Sum([]byte(sb.String())))
}
