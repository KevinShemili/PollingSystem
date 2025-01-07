// request structs for authentication endpoints
// validator is used to validate the request fields

package requests

type RegisterRequest struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"omitempty"`
	Age       int    `json:"age" validate:"omitempty"`
	Email     string `json:"email" validate:"required"`
	Password  string `json:"password" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type TokensRequest struct {
	JWTToken     string `json:"jwt_token" validate:"required"`
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type LogOutRequest struct {
	UserID int `json:"user_id" validate:"required"`
}
