/*
Structs containing the structure of the broadcast messages sent to the clients
Broadcasts happen when:
- Poll is created
- A vote is cast
- A poll expires
- A poll is deleted
- A poll is updated - uses same structure as poll creation

*/

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

type BroadcastVote struct {
	BroadcastType string `json:"type"`
	Data          struct {
		PollID     uint `json:"poll_id"`
		Categories []struct {
			CategoryID uint `json:"category_id"`
			Votes      int  `json:"category_votes"`
		} `json:"categories"`
	} `json:"data"`
}

type BroadcastExpiry struct {
	BroadcastType string `json:"type"`
	Data          struct {
		PollID uint `json:"poll_id"`
	} `json:"data"`
}

type BroadcastDeletion struct {
	BroadcastType string `json:"type"`
	Data          struct {
		PollID uint `json:"poll_id"`
	} `json:"data"`
}
