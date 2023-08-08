package course

import "time"

type Course struct {
	ID           int
	UserID       int
	Title        string
	Materi       string
	Category     string
	Division     string
	Playlist     string
	Slug         string
	LinkVideo    string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	CourseImages []CourseImage
}

type CourseImage struct {
	ID        int
	CourseID  int
	IsPrimary int
	FileName  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
