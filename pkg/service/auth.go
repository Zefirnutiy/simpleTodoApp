package service

import (
	"Zefirnutiy/simpleTodoApp/pkg/repository"
	"Zefirnutiy/simpleTodoApp/structs"
	"crypto/sha1"
	"fmt"
)

const salt = "dsfkslkivyiv285498#&^*&"
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