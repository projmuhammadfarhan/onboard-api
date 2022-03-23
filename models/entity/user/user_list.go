package user

type UserList struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	RoleID string `json:"role_id"`
	Title  string `json:"title"`
	Active bool   `json:"active"`
}
