package get_user_restaurants

type Request struct {
	Type    string `query:"type"`
	Name    string `query:"name"`
	Address string `query:"address"`
}

type Response []struct {
	Type    string `json:"type"`
	Name    string `json:"name"`
	Address string `json:"address"`
}
