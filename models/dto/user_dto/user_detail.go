package user_dto

type UserDetail struct {
	ID             string `json:"id"`
	PersonalNumber string `json:"personalNumber"`
	Email          string `json:"email"`
	Name           string `json:"name"`
	RoleID         string `json:"role"`
	Title          string `json:"title"`
	Active         bool   `json:"active"`
}
