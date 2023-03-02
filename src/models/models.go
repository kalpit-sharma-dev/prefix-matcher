package models

type Response struct {
	LCP string `json:"lcp,omitempty"`
}

type Error struct {
	Id      int    `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}
