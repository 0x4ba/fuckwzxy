package utils

import (
	"fmt"
	"math/rand"
	"net/url"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixMicro())
}

func RandomTemperature() string {
	tpl := `["0","36.%d","36.%d","36.%d"]`

	return fmt.Sprintf(tpl, rand.Intn(70), rand.Intn(70), rand.Intn(70))
}

//URL 包里的 encode 对map 进行了 sort ,这个sort 会导致服务器无法使用，故重写
func UrlEncode(v url.Values) string {
	if v == nil {
		return ""
	}
	var buf strings.Builder
	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}
	for _, k := range keys {
		vs := v[k]
		keyEscaped := url.QueryEscape(k)
		for _, v := range vs {
			if buf.Len() > 0 {
				buf.WriteByte('&')
			}
			buf.WriteString(keyEscaped)
			buf.WriteByte('=')
			buf.WriteString(url.QueryEscape(v))
		}
	}
	return buf.String()
}
