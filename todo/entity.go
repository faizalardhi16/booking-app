package todo

import "time"

type Todo struct {
	ID        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
