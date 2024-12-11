package results

import "time"

type CreatePollResult struct {
	Title      string
	LifeTime   time.Time
	IsEnded    bool
	Categories []string
}
