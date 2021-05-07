package main

import "time"

type DockerImage struct {
	Username   string
	Password   string
	Registry   string
	Project    string
	Repository string
}

type Artifact struct {
	Id        int       `json:"id"`
	ProjectID int       `json:"project_id"`
	PullTime  time.Time `json:"pull_time"`
	PushTime  time.Time `json:"push_time"`
	Tags      []Tag     `json:"tags"`
}

type Tag struct {
	Id       int       `json:"id"`
	Name     string    `json:"name"`
	PullTime time.Time `json:"pull_time"`
	PushTime time.Time `json:"push_time"`
}
