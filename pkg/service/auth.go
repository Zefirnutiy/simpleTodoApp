package service

import (
	"Zefirnutiy/simpleTodoApp/pkg/repository"
	"Zefirnutiy/simpleTodoApp/structs"
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	salt = "dsfkslkivyiv285498#&^*&"
	singinKey = "qweqweioyisd*^&$352_"
	tokenTTL = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"userId"`

}

type AuthService struct {
	repo repository.Authorization
}

func newAuthService(repo repository.Authorization)*AuthService{
	return &AuthService{repo}
}

func (s *AuthService) CreateUser(user structs.User)(int, error){
	user.Password = s.generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *AuthService) GenerateToken(userName, password string)(string, error){
	user, err := s.repo.GetUser(userName, s.generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims {
			jwt.StandardClaims{
		ExpiresAt: time.Now().Add(tokenTTL).Unix(),
		IssuedAt: time.Now().Unix() ,
	},
		user.Id,
	})

	return token.SignedString([]byte(singinKey))
}