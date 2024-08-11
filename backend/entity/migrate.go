package entity

import "time"

type MigrateHistory struct {
	ID        uint      `json:"id"`
	Key       string    `json:"key"`
	Hash      string    `json:"hash"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
