package requests

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TokensRequest struct {
	JWTToken     string `json:"jwt_token"`
	RefreshToken string `json:"refresh_token"`
}

type LogOutRequest struct {
	UserID int `json:"user_id"`
}
