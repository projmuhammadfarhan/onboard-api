package login_dto

type UserLogin struct {
	PersonalNumber string `json:"personal_number" binding:"required"`
	Password       string `json:"password" binding:"required"`
}
