package requests

type CreatePollRequest struct {
	Title      string   `json:"title"`
	ExpiresAt  string   `json:"expires_at"`
	Categories []string `json:"categories"`
}
