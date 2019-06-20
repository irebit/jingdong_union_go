package jingdong_union_go

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"sort"
	"strconv"
	"strings"
)

func GetSign(clientSecret string, p map[string]interface{}) string {
	var keys []string
	for k := range p {
		if k != "sign" && k != "access_token" {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)
	signStr := clientSecret
	for _, key := range keys {
		signStr += key + GetString(p[key])
	}
	signStr += clientSecret
	// log.Println(signStr, "=", md5Hash(signStr))
	return md5Hash(signStr)
}

func GetString(i interface{}) string {
	switch v := i.(type) {
	case string:
		return v
	case int:
		return strconv.Itoa(v)
	case bool:
		return strconv.FormatBool(v)
	default:
		bytes, _ := json.Marshal(v)
		return string(bytes)
	}
	return ""
}

func md5Hash(signStr string) string {
	h := md5.New()
	h.Write([]byte(signStr))
	cipherStr := h.Sum(nil)
	return strings.ToUpper(hex.EncodeToString(cipherStr))
}
