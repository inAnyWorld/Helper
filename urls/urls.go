package urls

import (
	"net/url"
	"strings"
)

// ParseUriQueryToMap 解析url参数到map
func ParseUriQueryToMap(query string) map[string]interface{} {
	qm := strings.Split(query, "&")
	convert := make(map[string]interface{}, len(qm))
	for _, v := range qm {
		split := strings.Split(v, "=")
		convert[split[0]] = split[1]
	}
	return convert
}

// GetQueryParams url params 不区分大小写
func GetQueryParams(mValues map[string][]string, key string) string {
	paramsToLowerMap := make(map[string][]string)
	// key 统一转小写
	for pk, pv := range mValues {
		paramsToLowerMap[strings.ToLower(pk)] = pv
	}
	// 获取key是否存在
	if values, ok := paramsToLowerMap[strings.ToLower(key)]; ok {
		return values[0]
	}
	return ""
}

// GetDomain 获取域名
func GetDomain(str string, isMains ...bool) string {
	u, err := url.Parse(str)
	isMain := false
	if len(isMains) > 0 {
		isMain = isMains[0]
	}

	if err != nil {
		return ""
	} else if !isMain {
		return u.Hostname()
	}

	parts := strings.Split(u.Hostname(), ".")
	domain := parts[len(parts)-2] + "." + parts[len(parts)-1]

	return domain
}