package login

type Request struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}
