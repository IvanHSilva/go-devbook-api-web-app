package models

import (
	"time"
)

// User model
type User struct {
	ID      uint64    `json:"id,omitempty"`
	Name    string    `json:"name,omitempty"`
	EMail   string    `json:"email,omitempty"`
	Pass    string    `json:"password,omitempty"`
	RegDate time.Time `json:"regdate,omitempty"`
}
