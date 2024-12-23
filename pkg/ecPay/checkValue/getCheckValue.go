package checkValue

import (
	"errors"
	"fmt"
	create_order "member/router/app/content/create-order"
	"reflect"
)

func GetCheckValue(req *create_order.Request) (value1 string) {

	// 將Struct轉為map[string]string
	t := reflect.TypeOf(req)
	v := reflect.ValueOf(req)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	var keys map[string]string

	if t.Kind() == reflect.Struct {
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i).Name
			value := v.Field(i).String()
			if value != "" {
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

	//for key := range keys {
	//
	//}
	fmt.Println(keys)

	// 將傳入的參數按第一個英文字母做排序，遇到相同的則依第二個字母

	//fmt.Println("param: ", param)
	//sort.Strings()
	// 參數前面加上HashKey參數、最後面加上HashIV

	// 整串字串進行URL encode

	// 轉為小寫

	// 再以SHA256加密

	// 再轉大寫

	return
}
