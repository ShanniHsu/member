package login

type Request struct {
	Account  string `json:"account" validate:"required"`
	Password string `json:"password" validate:"required"`
}
