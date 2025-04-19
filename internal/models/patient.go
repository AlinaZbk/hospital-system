package models

type Patient struct {
	Person
	WardID  int             `json:"ward_id"`
	History []MedicalRecord `json:"history"`
}
