package proxy

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	create_order "member/router/app/content/create-order"
	"net/http"
	"net/url"
	"reflect"
)

func CreateOrder(req *create_order.Request) (body []byte, err error) {
	apiURL := "https://payment-stage.ecpay.com.tw/Cashier/AioCheckOut/V5"
	formData := url.Values{}

	t := reflect.TypeOf(req)
	v := reflect.ValueOf(req)

	// 如果是指標的話，先解引用
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	// 確保是Struct結構
	if t.Kind() == reflect.Struct {
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i).Name
			value := v.Field(i).String()
			if value != "" {
				formData.Set(field, value)
			}
		}
	} else {
		err = errors.New("Because it not type of struct, it can't iterate over fields ")
		return
	}

	// 將表單數據編碼為 application/x-www-form-urlencoded 格式
	formEncoded := formData.Encode()

	// 創建 HTTP 請求
	request, err := http.NewRequest("POST", apiURL, bytes.NewBufferString(formEncoded))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// 設置請求頭，指定 Content-Type
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// 發送請求
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// 讀取響應
	if resp.StatusCode == http.StatusOK {
		fmt.Println("Request succeeded!")
	} else {
		fmt.Printf("Request failed with status: %d\n", resp.StatusCode)
	}

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		err = errors.New("Failed to read the body!")
		return
	}

	return
}
