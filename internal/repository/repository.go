package repository

type Repository struct {
	Authorization
	PatientRepository
	DoctorRepository
	WardRepository
	DiseaseRepository
	MedicalRecordRepository
}

func NewRepository() *Repository {
	return &Repository{}
}