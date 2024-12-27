package requests

type RegisterRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
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
