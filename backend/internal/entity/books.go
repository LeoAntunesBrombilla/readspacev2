package entity

import "time"

type Volume struct {
	VolumeInfo VolumeInfo `json:"volumeInfo"`
}

type VolumeInfo struct {
	Title       string     `json:"title"`
	Subtitle    string     `json:"subtitle"`
	Authors     []string   `json:"authors"`
	Publisher   string     `json:"publisher"`
	Description string     `json:"description"`
	PageCount   int        `json:"pageCount"`
	Categories  []string   `json:"categories"`
	Language    string     `json:"language"`
	ImageLinks  ImageLinks `json:"imageLinks"`
}

type ImageLinks struct {
	SmallThumbnail string `json:"smallThumbnail"`
	Thumbnail      string `json:"thumbnail"`
}

type ExternalBook struct {
	Title       string   `json:"title"`
	Subtitle    string   `json:"subtitle"`
	Authors     []string `json:"authors"`
	Publisher   string   `json:"publisher"`
	Description string   `json:"description"`
	PageCount   int      `json:"pageCount"`
	Categories  []string `json:"categories"`
	Language    string   `json:"language"`
	ImageLinks  struct {
		SmallThumbnail string `json:"smallThumbnail"`
		Thumbnail      string `json:"thumbnail"`
	} `json:"imageLinks"`
}

type BookList struct {
	ID          int           `json:"id" db:"id"`
	Name        string        `json:"name" db:"name"`
	CreatedAt   time.Time     `json:"created_at" db:"created_at"`
	Books       []ExternalBook `json:"books"`
}
