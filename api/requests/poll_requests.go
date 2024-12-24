package requests

type CreatePollRequest struct {
	Title      string   `json:"title"`
	ExpiresAt  string   `json:"expires_at"`
	Categories []string `json:"categories"`
}

type AddVoteRequest struct {
	PollID         uint `json:"-"`
	PollCategoryID uint `json:"poll_category_id"`
}
