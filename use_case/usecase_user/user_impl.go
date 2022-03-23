package usecase_user

import (
	"errors"

	"gorm.io/gorm"
	"main.go/helper"
	"main.go/models"
	"main.go/models/dto/login_dto"
	"main.go/models/dto/user_dto"
	"main.go/models/entity/user"
	"main.go/use_case/jwt_usecase"
)

func (user *userUsecase) GetUsers() models.Response {
	userlist, err := user.userRepo.GetUsers()
	response := []user_dto.UserList{}
	for _, user := range userlist {
		role := models.Role{ID: user.RoleID, Title: user.Title}
		responseData := user_dto.UserList{
			ID:     user.ID,
			Name:   user.Name,
			Role:   role,
			Active: user.Active,
		}
		response = append(response, responseData)
	}

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return helper.ResponseError("Data not found", err, 404)
	} else if err != nil {
		return helper.ResponseError("Internal server error", err, 500)
	}
	return helper.ResponseSuccess("ok", nil, response, 200)
}

func (user *userUsecase) GetUser(id string) models.Response {
	userData, err := user.userRepo.GetUser(id)

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return helper.ResponseError("Data not found", err, 404)
	} else if err != nil {
		return helper.ResponseError("Internal server error", err, 500)
	}

	role := models.Role{
		ID:    userData.RoleID,
		Title: userData.Title,
	}

	userResponse := map[string]interface{}{
		"id":              userData.ID,
		"name":            userData.Name,
		"email":           userData.Email,
		"role":            role,
		"personal_number": userData.PersonalNumber,
		"active":          userData.Active,
	}
	return helper.ResponseSuccess("ok", nil, userResponse, 200)
}

func (users *userUsecase) CreateUser(newUser user_dto.User) models.Response {
	userInsert := user.User{
		ID:             newUser.ID,
		Name:           newUser.Name,
		Email:          newUser.Email,
		PersonalNumber: newUser.PersonalNumber,
		Password:       newUser.Password,
	}

	userData, _, err := users.userRepo.CreateUser(userInsert)
	// fmt.Println()
	if err != nil {
		return helper.ResponseError("Internal server error", err, 500)
	}

	return helper.ResponseSuccess("ok", nil, map[string]interface{}{
		"id": userData.ID}, 201)
}

func (users *userUsecase) UpdateUser(userUpdate user_dto.User, id string) models.Response {
	userInsert := user.User{
		Name:           userUpdate.Name,
		Email:          userUpdate.Email,
		PersonalNumber: userUpdate.PersonalNumber,
		Active:         userUpdate.Active,
		Password:       userUpdate.Password,
		RoleID:         userUpdate.RoleID.ID,
	}
	_, err := users.userRepo.UpdateUser(userInsert, id)

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return helper.ResponseError("Data not found", err, 404)
	} else if err != nil {
		return helper.ResponseError("Internal server error", err, 500)
	}

	userUpdate.ID = id
	return helper.ResponseSuccess("ok", nil, map[string]interface{}{"id": id}, 200)
}

func (user *userUsecase) DeleteUser(id string) models.Response {

	err := user.userRepo.DeleteUser(id)

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return helper.ResponseError("Data not found", err, 404)
	} else if err != nil {
		return helper.ResponseError("Internal server error", err, 500)
	}
	return helper.ResponseSuccess("ok", nil, nil, 200)
}

func (users *userUsecase) UserLogin(userLogin login_dto.UserLogin) models.Response {
	userData, err := users.userRepo.GetUserByPN(userLogin.PersonalNumber)

	if err != nil {
		return helper.ResponseError("User not found", map[string]interface{}{"message": "Personal Number not found"}, 404)
	}

	errPwd := helper.CheckPasswordHash(userLogin.Password, userData.Password)

	if errPwd != nil {
		return helper.ResponseError("User not found", map[string]interface{}{"message": "Wrong Password"}, 404)
	}

	jwt := jwt_usecase.GetJwtUsecase(users.userRepo)

	response, _ := jwt.GenerateToken(userData.ID, userData.RoleID)

	return helper.ResponseSuccess("ok", nil, map[string]interface{}{"token": response}, 200)
}
