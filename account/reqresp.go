package account

type (
	CreateUserRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	CreateUserResponse struct {
		ID string `json:"id"`
	}

	GetUserRequest struct {
		ID string `json:"id"`
	}
	GetUserResponse struct {
		Email string `json:"email"`
	}
)
