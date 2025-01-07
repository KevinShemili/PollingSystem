// request structs for poll related requests
// validator is used to validate the request fields

package requests

import "gin/application/utility"

type CreatePollRequest struct {
	Title      string   `json:"title" validate:"required"`
	ExpiresAt  string   `json:"expires_at" validate:"required"`
	Categories []string `json:"categories" validate:"required"`
}

type AddVoteRequest struct {
	PollID         uint `json:"-" validate:"required"`
	PollCategoryID uint `json:"poll_category_id" validate:"required"`
}

type UpdatePollRequest struct {
	PollID           uint     `json:"-" validate:"required"`
	Title            string   `json:"title" validate:"required"`
	ExpiresAt        string   `json:"expires_at" validate:"required"`
	Description      string   `json:"description" validate:"required"`
	DeleteCategories []uint   `json:"delete_categories" validate:"omitempty"`
	NewCategories    []string `json:"new_categories" validate:"omitempty"`
}

type GetPollsRequest struct {
	utility.QueryParams

	ShowActiveOnly bool `json:"show_active_only" validate:"omitempty"`
}
