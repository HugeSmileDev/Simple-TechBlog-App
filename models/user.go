package models

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Expertise string `json:"expertise"`
}

func ValidateUser(user *User) string {
	if user.Name == "" {
		return "Name is required"
	}
	if user.Email == "" {
		return "Email is required"
	}
	return ""
}
