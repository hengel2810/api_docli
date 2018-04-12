package models

import "time"

type RequestDockerImage struct {
	Name string
	Path  string
	UserId string
	Uploaded time.Time
}

type RequestDockerImageError struct {
	StatusCode int
	Msg string
}