package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up")
		auth.POST("/sign-in")
	}

	api := router.Group("/api")
	{
		patient := api.Group("/patients")
		{
			patient.GET("/", h.GetAllPatients)       // список всех пациентов
			patient.GET("/:id", h.GetPatientByID)    // карточка конкретного пациента
			patient.POST("/", h.CreatePatient)      // добавление нового пациента
			patient.PUT("/:id", h.UpdatePatient)    // изменить данные пациента
			patient.DELETE("/:id", h.DeletePatient) // удаление пациента
		}

		doctor := api.Group("/doctors")
		{
			doctor.GET("/", h.GetAllDoctors)    // список всех докторов
			doctor.GET("/:id", h.GetDoctorByID) // информация о конкретном докторе
			doctor.POST("/", h.CreateDoctor)   // добавление нового врача
		}

		ward := api.Group("/wards")
		{
			ward.GET("/", h.GetAllWrads)       // список всех палат
			ward.GET("/:id", h.GetWardByID)    // информация о конкретной палате
			ward.POST("/", h.CreateWard)      // добавление новой палаты
			ward.PUT("/:id", h.UpdateWard)    // заселить пациента
			ward.DELETE("/:id", h.DeleteWard) // выселить пациента
		}

		disease := api.Group("/diseases")
		{
			disease.GET("/", h.GetAllDiseases)  // список всех болезней
			disease.POST("/", h.CreateDisease) // добавление новой болезни
		}

		medical_record := api.Group("/medical-records")
		{
			medical_record.GET("/", h.GetAllMedicalRecords)    // весь список истории болезни
			medical_record.GET("/:id", h.GetMedicalRecordByID) // история болезни конкретного пациента
			medical_record.POST("/", h.CreateMedicalRecord)   // добавление записи
		}

	}
	return router
}
