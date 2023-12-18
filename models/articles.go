package models

import "time"

type Article struct {
	Title    string   `json:"title"`
	Desc     string   `json:"desc"`
	Tags     []string `json:"tags"`
	Author   string   `json:"author"`
	MusicId  string   `josn:"musicId`
	Path     string
	ShortUrl string
	Category string
	Date     time.Time `json:"date"`
}
