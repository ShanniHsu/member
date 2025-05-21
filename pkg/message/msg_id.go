package message

/*
00: 一般成功/提示訊息
10: 查無資料、Not Found
20: 欄位驗證錯誤
30: 權限或認證錯誤
40: 外部API錯誤
50: 資料庫/邏輯錯誤
90: 未知或系統錯誤
*/
const (
	// auth
	UNAUTHORIZED = "A103001"

	// database
	GETDATAFAILED = "D101001"

	// system
	SYSTEM_ERROR = "S109001"

	// user
	REGISTER_SUCCESSFULLY = "U100001"
	LOGIN_SUCCESSFULLY    = "U100002"
	ACCOUNT_EXISTED       = "U101001"
	REGISTER_FAILED       = "U109001"
)
