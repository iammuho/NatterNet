package dto

// SignInReqDTO represents the required fields for a user to sign in.
// swagger:model SignInReqDTO
type SignInReqDTO struct {
	// The login of the user. It must be alphanumeric and have a length between 3 and 100.
	// required: true
	// example: john_doe
	Login string `json:"login" validate:"required,min=3,max=100,alphanum"`

	// The password of the user. It must have a length between 8 and 32.
	// required: true
	// example: securePassword123
	Password string `json:"password" validate:"required,min=8,max=32"`
}
