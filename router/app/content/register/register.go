package register

type Request struct {
	Account  string `json:"account" binding:"max=10,min=6"`
	Password string `json:"password" binding:"max=10,min=6"`
	Nickname string `json:"nickname" binding:"required"`
	Email    string `json:"email" binding:"email"`
}
