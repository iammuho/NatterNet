package dto

type SignInReqDTO struct {
	Login    string `json:"login" validate:"required,min=3,max=100,alphanum"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}
