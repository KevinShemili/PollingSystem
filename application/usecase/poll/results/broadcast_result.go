package results

import "time"

type BroadcastPoll struct {
	BroadcastType string `json:"type"`
	Data          struct {
		PollID     uint      `json:"poll_id"`
		Title      string    `json:"title"`
		ExpiresAt  time.Time `json:"expires_at"`
		Ended      bool      `json:"ended"`
		Categories []struct {
			ID    uint   `json:"category_id"`
			Name  string `json:"category_name"`
			Votes int    `json:"category_votes"`
		} `json:"categories"`
	} `json:"data"`
}
