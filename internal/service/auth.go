package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/AlinaZbk/hospital-system/internal/models"
	"github.com/AlinaZbk/hospital-system/internal/repository"
)

const salt = "hjkls23456bnmsk"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthSerice(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user models.User) (int, error) {
	user.Password = s.generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
