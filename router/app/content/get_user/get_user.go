package get_user

type Response struct {
	Account  string `json:"account"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	Status   int64  `json:"status"`
}
