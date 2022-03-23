package jwt_usecase

import (
	"github.com/golang-jwt/jwt"
	"main.go/repository/user_repo"
)

type JwtUsecase interface {
	GenerateToken(string, string) (string, error)
	ValidateToken(string) (*jwt.Token, error)
	ValidateTokenAndGetUserId(string) (string, error)
}

type jwtUsecase struct {
	userRepo user_repo.UserRepository
}

func GetJwtUsecase(repo user_repo.UserRepository) JwtUsecase {
	return &jwtUsecase{
		userRepo: repo,
	}
}
