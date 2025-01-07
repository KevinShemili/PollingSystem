// results produced by the authentication use cases

package results

type LoginResult struct {
	JWTToken     string `json:"jwt_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshResult struct {
	JWTToken     string `json:"jwt_token"`
	RefreshToken string `json:"refresh_token"`
}
