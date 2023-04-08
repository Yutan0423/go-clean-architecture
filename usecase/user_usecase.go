package usecase

import (
	"go-clean-architecture/entity"
	"go-clean-architecture/repository"
	"go-clean-architecture/validator"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type IUserUseCase interface {
	SignUp(u entity.User) (entity.UserResponse, error)
	SignIn(u entity.User) (string, error)
}

type userUseCase struct {
	ur repository.IUserRepository
	uv validator.IUserValidator
}

func NewUserUseCase(ur repository.IUserRepository, uv validator.IUserValidator) IUserUseCase {
	return &userUseCase{ur, uv}
}

func (uu *userUseCase) SignUp(u entity.User) (entity.UserResponse, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	if err != nil {
		return entity.UserResponse{}, err
	}

	newUser := entity.User{Email: u.Email, Password: string(hash)}
	if err := uu.ur.CreateUser(&newUser); err != nil {
		return entity.UserResponse{}, err
	}

	resUser := entity.UserResponse{
		ID:    newUser.ID,
		Email: newUser.Email,
	}
	return resUser, nil
}

func (uu *userUseCase) SignIn(u entity.User) (string, error) {
	storedUser := entity.User{}
	if err := uu.ur.GetUserByEmail(&storedUser, u.Email); err != nil {
		return "", err
	}
	err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(u.Password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
