package models

import "time"

type Note struct {
	Title string `json:"title"`
	Content string `json:"content"`
	CreatedAt time.Time `json:"create_at"`
	Tag string `json:"tag"`
}