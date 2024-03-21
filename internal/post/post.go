package post

import "time"

type Post struct {
	ID      int
	IDUser  int
	Date    time.Time
	Title   string
	Content string
}
