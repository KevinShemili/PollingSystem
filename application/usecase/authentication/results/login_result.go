package results

type LoginResult struct {
	AuthenticationToken string `json:"authentication_token"`
	RefreshToken        string `json:"refresh_token"`
}
