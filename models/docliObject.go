package models

import (
	"time"
	"gopkg.in/mgo.v2/bson"
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
	Id bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	FullName string `json:"full_name"`
	UserId string `json:"user_id"`
	OriginalName string `json:"image_name"`
	UniqueId string `json:"unique_id"`
	Ports []int `json:"ports"`
	ServerPorts []PortObject `json:"server_ports"`
	Networks []string `json:"networks"`
	Volumes []string `json:"volumes"`
	Uploaded time.Time
	ContainerName string
}

type PortObject struct {
	Container int `json:"container"`
	Host int `json:"host"`
}