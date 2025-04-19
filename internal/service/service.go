package service

import "github.com/AlinaZbk/hospital-system/internal/repository"

type Authorization interface {
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
	return &Service{}
}