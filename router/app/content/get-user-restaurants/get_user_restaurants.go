package get_user_restaurants

type Request struct {
	ID       int64  `query:"id"`
	Type     int64  `query:"type"`
	Name     string `query:"name"`
	Address  string `query:"address"`
	Page     string `query:"page"`
	PageSize string `query:"page_size"`
}

type Response struct {
	List       []List `json:"list"`
	TotalCount int64  `json:"total_count"`
}

type List struct {
	ID      int64  `json:"id"`
	Type    string `json:"type"`
	Name    string `json:"name"`
	Address string `json:"address"`
}
