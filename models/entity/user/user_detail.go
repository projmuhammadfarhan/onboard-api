package user

type UserDetail struct {
	ID             string `json:"id"`
	PersonalNumber string `json:"personal_number"`
	Email          string `json:"email"`
	Name           string `json:"name"`
	RoleID         string `json:"role_id"`
	Title          string `json:"title"`
	Active         bool   `json:"active"`
}
