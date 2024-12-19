package routes

import (
	controller "Doctor-Appointment-Project/controllers"
	middleware "Doctor-Appointment-Project/middleware"

	"github.com/gin-gonic/gin"
)

func DoctorRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate()) // this protect our routes

	incomingRoutes.GET("/doctor/MyAppointment", controller.Check_my_appointment())

	incomingRoutes.GET("/doctor/get_my_profile", controller.Get_my_profile())

	incomingRoutes.GET("/doctor/feedback", controller.Doctor_Checking_Feedback())

	incomingRoutes.POST("/add_prescription", controller.Add_prescription())

	incomingRoutes.POST("/doctor", controller.Add_docter()) //done

	incomingRoutes.PUT("/doctor", controller.Update_docter()) //done

	incomingRoutes.DELETE("/doctor", controller.Delete_docter()) //done

}
