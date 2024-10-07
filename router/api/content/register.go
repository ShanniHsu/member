package content

type Request struct {
	Account  string `json:"account" validate:"max=10,min=6"`
	Password string `json:"password" validate:"max=10,min=6"`
}
