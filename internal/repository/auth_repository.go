package repository

import "github.com/AlinaZbk/hospital-system/internal/models"

type Authorization interface {
	CreateUser(user models.User) (int, error)
}
