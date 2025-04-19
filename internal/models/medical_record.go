package models

import "time"

type MedicalRecord struct {
	ID        int       `json:"id"`
	PatientID int       `json:"patient_id"`
	DoctorID  int       `json:"doctor_id"`
	DiseaseID int       `json:"disease_id"`
	Notes     string    `json:"notes"`
	CreatedAt time.Time `json:"created_at"`
}
