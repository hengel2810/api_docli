package models

import (
	"time"
)

func DocliObjectValid(docliObject DocliObject) bool {
	if docliObject.FullName == "" {
		return false
	}
	if docliObject.UserId == "" {
		return false
	}
	if docliObject.OriginalName == "" {
		return false
	}
	if docliObject.UniqueId == "" {
		return false
	}
	return true
}

type DocliObject struct {
	FullName string `json:"full_name"`
	UserId string `json:"user_id"`
	OriginalName string `json:"image_name"`
	UniqueId string `json:"unique_id"`
	Ports []PortObject `json:"ports"`
	Networks []string `json:"networks"`
	Volumes []string `json:"volumes"`
	Uploaded time.Time
}

type PortObject struct {
	InternalPort int `json:"ex"`
	ExternalPort int `json:"int"`
}