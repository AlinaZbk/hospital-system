package models

type Ward struct {
	ID       int       `json:"id"`
	Number   int       `json:"number"`
	Capacity int       `json:"capacity"`
	Patients []Patient `json:"patients"`
}
