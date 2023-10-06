package models

type Post struct {
	UserId int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func ValidatePost(post *Post) string {
	if post.Title == "" {
		return "Title is required"
	}
	if post.Body == "" {
		return "Body is required"
	}
	if post.UserId < 1 {
		return "Invalid UserId"
	}
	return ""
}
