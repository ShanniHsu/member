package create_order

type Request struct {
	MerchantID        string `json:"MerchantID"`
	MerchantTradeNo   string `json:"MerchantTradeNo"`
	MerchantTradeDate string `json:"MerchantTradeDate"`
	PaymentType       string `json:"PaymentType"`
	TotalAmount       string `json:"TotalAmount"`
	TradeDesc         string `json:"TradeDesc"`
	ItemName          string `json:"ItemName"`
	ReturnURL         string `json:"ReturnURL"`
	ChoosePayment     string `json:"ChoosePayment"`
	CheckMacValue     string `json:"CheckMacValue"`
	EncryptType       string `json:"EncryptType"`
	StoreID           string `json:"StoreID"`
	ClientBackURL     string `json:"ClientBackURL"`
	ItemURL           string `json:"ItemURL"`
	Remark            string `json:"Remark"`
	ChooseSubPayment  string `json:"ChooseSubPayment"`
	OrderResultURL    string `json:"OrderResultURL"`
	NeedExtraPaidInfo string `json:"NeedExtraPaidInfo"`
	IgnorePayment     string `json:"IgnorePayment"`
	PlatformID        string `json:"PlatformID"`
	CustomField1      string `json:"CustomField1"`
	CustomField2      string `json:"CustomField2"`
	CustomField3      string `json:"CustomField3"`
	CustomField4      string `json:"CustomField4"`
	Language          string `json:"Language"`
}
