package repository

import "github.com/jmoiron/sqlx"

type Repository struct {
	Authorization
	PatientRepository
	DoctorRepository
	WardRepository
	DiseaseRepository
	MedicalRecordRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
