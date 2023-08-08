package course

type CourseFormatter struct {
	ID       int    `json:"id"`
	UserID   int    `json:"user_id"`
	Title    string `json:"title"`
	Category string `json:"category"`
	Division string `json:"division"`
	Playlist string `json:"playlist"`
	Slug     string `json:"slug"`
	ImageURL string `json:"image_url"`
}

func FormatCourse(course Course) CourseFormatter {
	courseFormatter := CourseFormatter{}
	courseFormatter.ID = course.ID
	courseFormatter.UserID = course.UserID
	courseFormatter.Title = course.Title
	courseFormatter.Category = course.Category
	courseFormatter.Division = course.Division
	courseFormatter.Playlist = course.Playlist
	courseFormatter.Slug = course.Slug
	courseFormatter.ImageURL = ""

	if len(course.CourseImages) > 0 {
		courseFormatter.ImageURL = course.CourseImages[0].FileName
	}

	return courseFormatter
}

func FormatCourses(courses []Course) []CourseFormatter {
	coursesFormatter := []CourseFormatter{}

	for _, course := range courses {
		courseFormatter := FormatCourse(course)
		coursesFormatter = append(coursesFormatter, courseFormatter)
	}
	return coursesFormatter
}
