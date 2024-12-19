package routes

import (
	controller "Doctor-Appointment-Project/controllers"
	middleware "Doctor-Appointment-Project/middleware"

	"github.com/gin-gonic/gin"
)

func PatientRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())

	// Patient

	incomingRoutes.POST("/patient", controller.Add_patient())
	incomingRoutes.GET("/get_patient_details", controller.Get_patient_details())
	incomingRoutes.DELETE("/patient", controller.Delete_patient())

	//Doctor

	incomingRoutes.POST("/doctor_feedback", controller.Doctor_feedback())
	incomingRoutes.GET("/showall/doctors", controller.Get_docter())
	incomingRoutes.GET("/get_doctor_by_city", controller.GetDoctorByLocation())
	incomingRoutes.POST("/book_appointment", controller.Book_doctor_appointment())
	incomingRoutes.DELETE("/cancelAppointment", controller.Cancel_doctor_appointment())

	// Lab

	incomingRoutes.POST("/book_test", controller.Book_test())
	incomingRoutes.GET("/get_lab_by_location", controller.Get_lab_by_location())
	incomingRoutes.POST("/lab_feedback", controller.Lab_feedback())

	// Nurse

	incomingRoutes.GET("/nurse/by_city", controller.Get_nurse_by_city())
	incomingRoutes.GET("/nurse/by_specialisation", controller.Get_nurse_by_specialisation())
	incomingRoutes.GET("/nurse/by_location", controller.Get_nurse_by_location())
	incomingRoutes.GET("/nurse", controller.Get_nurse())
	incomingRoutes.POST("/nurse/feedback", controller.Nurse_feedback())
	incomingRoutes.POST("/nurse/book", controller.Book_nurse_appointment())
	incomingRoutes.DELETE("/nurse/cancel_appointment", controller.Cancel_nurse_appointment())

	// Prescription
	incomingRoutes.GET("/prescription", controller.Get_prescription())

	// order medicines

	incomingRoutes.POST("/medicines", controller.Order_medicines())
	incomingRoutes.DELETE("/medicines", controller.Cancel_ordered_medicines())

	//online_consultancy

	incomingRoutes.GET("/online_consultant", controller.Get_doctor_by_online_consultancy_availability())
	incomingRoutes.POST("/online_consultancy", controller.Book_online_consultancy_appointment())
	incomingRoutes.DELETE("/online_consultancy", controller.Cancel_online_consultancy_appointment())

	// Home Visit

	incomingRoutes.GET("/home_visit", controller.Get_doctor_by_home_visit_availability())
	incomingRoutes.POST("/home_visit", controller.Book_home_visit_appointment())
	incomingRoutes.DELETE("/home_visit", controller.Cancel_home_visit_appointment())

}
