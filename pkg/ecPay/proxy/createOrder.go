package proxy

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
)

func CreateOrder() (err error) {
	apiURL := "https://payment-stage.ecpay.com.tw/Cashier/AioCheckOut/V5"
	formData := url.Values{}

	formData.Set("MerchantID", "3002607")
	formData.Set("MerchantTradeNo", "shanni1234567890")
	formData.Set("MerchantTradeDate", "2024/12/19 17:46:00")
	formData.Set("PaymentType", "aio")
	formData.Set("TotalAmount", "100")
	formData.Set("TradeDesc", "Test the order")
	formData.Set("ItemName", "Shanni的商品")
	formData.Set("CheckMacValue", "")
	formData.Set("EncryptType", "1")
	formData.Set("StoreID", "")
	formData.Set("ClientBackURL", "")
	formData.Set("ItemURL", "")
	formData.Set("Remark", "")
	formData.Set("ChooseSubPayment", "")
	formData.Set("OrderResultURL", "")
	formData.Set("NeedExtraPaidInfo", "N")
	formData.Set("IgnorePayment", "")
	formData.Set("PlatformID", "")
	formData.Set("CustomField1", "")
	formData.Set("CustomField2", "")
	formData.Set("CustomField3", "")
	formData.Set("CustomField4", "")
	formData.Set("Language", "")
	// 將表單數據編碼為 application/x-www-form-urlencoded 格式
	formEncoded := formData.Encode()

	// 創建 HTTP 請求
	req, err := http.NewRequest("POST", apiURL, bytes.NewBufferString(formEncoded))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// 設置請求頭，指定 Content-Type
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// 發送請求
	client := &http.Client{}
	resp, err := client.Do(req)
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
	return
}
