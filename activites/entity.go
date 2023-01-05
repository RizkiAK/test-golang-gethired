package activites

import "time"

type Activites struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
