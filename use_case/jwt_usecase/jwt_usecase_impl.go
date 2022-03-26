package jwt_usecase

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

type CustomClaim struct {
	jwt.StandardClaims
	UserID string `json:"user_id"`
	Role   string `json:"role"`
}

func (jwtAuth *jwtUsecase) GenerateToken(userId string, role string) (string, error) {
	data, err := jwtAuth.userRepo.GetRoleByRoleId(role)
	if err != nil {
		return "user not found", err
	}

	claim := CustomClaim{
		UserID: userId,
		Role:   data.Title,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Minute * 60).Unix(),
			Issuer:    os.Getenv("APP_NAME"),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claim)
	return token.SignedString([]byte(os.Getenv("SECRET_KEY")))
}

func (jwtAuth *jwtUsecase) ValidateToken(token string) (*jwt.Token, error) {

	if !strings.Contains(token, "Bearer ") {
		log.Print("invalid Log   ", token)
		return nil, fmt.Errorf("Invalid")
	}
	tokenR := strings.Replace(token, "Bearer ", "", -1)
	return jwt.Parse(tokenR, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
}

func (jwtAuth *jwtUsecase) ValidateTokenAndGetUserId(token string) (string, error) {

	validatedToken, err := jwtAuth.ValidateToken(token)
	if err != nil {
		return "", err
	}

	claims, ok := validatedToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("failed to claim token")
	}

	return claims["user_id"].(string), nil
}

func (jwtAuth *jwtUsecase) ValidateTokenAndGetRole(token string) (string, string, error) {
	validatedToken, err := jwtAuth.ValidateToken(token)
	if err != nil {
		return "", "", err
	}

	claims, ok := validatedToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", "", errors.New("failed to claim token")
	}

	userdata, err := jwtAuth.userRepo.GetUser(claims["user_id"].(string))
	if err != nil {
		return "", "", err
	}

	role, err := jwtAuth.userRepo.GetRoleByRoleId(userdata.RoleID)
	if err != nil {
		return "", "", err
	}

	return claims["user_id"].(string), role.Title, nil
}
