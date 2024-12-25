package results

import "time"

type CreatePollResult struct {
	Title      string
	ExpiresAt  time.Time
	IsEnded    bool
	Categories []string
}

type UpdatePollResult struct {
	Title       string
	Description string
	ExpiresAt   time.Time
	Categories  []string
}
