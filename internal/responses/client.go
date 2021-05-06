package responses

import (
	"time"
)

type Client struct {
	Id    uint64
	Name  string
	Email string

	CreatedAt time.Time
	UpdatedAt time.Time
}
