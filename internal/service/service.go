package service

import (
	"github.com/AlinaZbk/hospital-system/internal/models"
	"github.com/AlinaZbk/hospital-system/internal/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
}

type PatientService interface {
}

type DoctorService interface {
}

type WardService interface {
}

type DiseaseService interface {
}

type MedicalRecordService interface {
}

type Service struct {
	Authorization
	PatientService
	DoctorService
	WardService
	DiseaseService
	MedicalRecordService
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthSerice(repos.Authorization),
	}
}
