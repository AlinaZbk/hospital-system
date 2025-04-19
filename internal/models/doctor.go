package models

type Doctor struct {
	Person
	Specialization string `json:"specialization"`
}
