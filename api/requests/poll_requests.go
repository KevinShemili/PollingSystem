package requests

import "gin/application/utility"

type CreatePollRequest struct {
	Title      string   `json:"title"`
	ExpiresAt  string   `json:"expires_at"`
	Categories []string `json:"categories"`
}

type AddVoteRequest struct {
	PollID         uint `json:"-"`
	PollCategoryID uint `json:"poll_category_id"`
}

type UpdatePollRequest struct {
	PollID           uint     `json:"-"`
	Title            string   `json:"title"`
	ExpiresAt        string   `json:"expires_at"`
	Description      string   `json:"description"`
	DeleteCategories []uint   `json:"delete_categories"`
	NewCategories    []string `json:"new_categories"`
}

type GetPollsRequest struct {
	utility.QueryParams

	ShowActiveOnly bool `json:"show_active_only"`
}
