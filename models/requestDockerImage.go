package models

import "time"

type RequestDockerImage struct {
	FullName   string      `json:"full_name"`
	UserId   string      `json:"user_id"`
	OriginalName   string      `json:"original_name"`
	UniqueId   string      `json:"unique_id"`
	Uploaded time.Time
}