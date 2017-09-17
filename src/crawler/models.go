package crawler

import "time"

// Topic Entity Description
type Topic struct {
	URL       string
	Title     string
	Content   string
	CreatedAt time.Time
}
