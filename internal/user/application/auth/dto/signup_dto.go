package dto

// SignupReqDTO represents the required fields for a user to sign up.
// swagger:model SignupReqDTO
type SignupReqDTO struct {
	// The email address of the user. It must be a valid email format.
	// required: true
	// example: john.doe@example.com
	Email string `json:"email" validate:"required,email"`

	// The desired username for the user. It must be alphanumeric and have a length between 3 and 20.
	// required: true
	// example: john_doe
	Username string `json:"username" validate:"required,min=3,max=20,alphanum"`

	// The desired password for the user. It must have a length between 8 and 32.
	// required: true
	// example: securePassword123
	Password string `json:"password" validate:"required,min=8,max=32"`
}
