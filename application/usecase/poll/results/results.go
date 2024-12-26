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

type GetPollResult struct {
	ID          uint
	Title       string
	Description string
	ExpiresAt   time.Time
	Categories  []struct {
		CategoryID   uint
		CategoryName string
		Votes        int
	}
}
