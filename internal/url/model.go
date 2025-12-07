package url

import "time"

type URL struct {
	Id        string    `json:"id"`
	Original  string    `json:"original"`
	Short     string    `json:"short"`
	Slug      string    `json:"slug"`
	SecretKey string    `json:"secret_key"`
	Clicks    int       `json:"clicks"`
	CreatedAt time.Time `json:"created_at"`
}
