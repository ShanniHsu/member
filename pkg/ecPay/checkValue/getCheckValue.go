package checkValue

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	create_order "member/router/app/content/create-order"
	"net/url"
	"reflect"
	"sort"
	"strings"
)

func GetCheckValue(req *create_order.Request) (checkValue string) {
	// 將Struct轉為map[string]string
	t := reflect.TypeOf(req)
	v := reflect.ValueOf(req)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	var keys map[string]string
	var arr []string

	if t.Kind() == reflect.Struct {
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i).Name
			value := v.Field(i).String()
			if value != "" {
				arr = append(arr, field)
				if len(keys) != 0 {
					keys[field] = value
				} else {
					keys = map[string]string{field: value}
				}
			}
		}
	} else {
		err := errors.New("Because it not type of struct, it can't iterate over fields ")
		fmt.Println(err)
		return
	}

	// 將傳入的參數按第一個英文字母做排序，遇到相同的則依第二個字母排序(快速排序法)
	sort.Strings(arr)

	// 重新組合參數及值
	var s string
	for i := 0; i < len(arr); i++ {
		s = fmt.Sprintf("%s&%s=%s", s, arr[i], keys[arr[i]])
	}
	// 參數前面加上HashKey參數、最後面加上HashIV
	s = fmt.Sprintf("HashKey=%s%s&HashIV=%s", "pwFHCqoQZGmho4w6", s, "EkRm7iFT261dpevs")
	// 整串字串進行URL encode
	encodeString := url.QueryEscape(s)
	// 如果要驗證可以使用解碼 url.QueryUnescape(encodeString)

	// 轉為小寫
	lowerString := strings.ToLower(encodeString)
	// 再以SHA256加密
	sha256String := generateSHA256(lowerString)
	// 再轉大寫
	checkValue = strings.ToUpper(sha256String)
	return
}

func generateSHA256(input string) (output string) {
	hash := sha256.Sum256([]byte(input))
	output = hex.EncodeToString(hash[:])
	return
}
