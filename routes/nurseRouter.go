package routes

import (
	controller "Doctor-Appointment-Project/controllers"
	middleware "Doctor-Appointment-Project/middleware"

	"github.com/gin-gonic/gin"
)

func NurseRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())

	incomingRoutes.GET("/get_nurse_profile", controller.Get_nurse_profile())

	incomingRoutes.GET("/nurse_checking_feedback", controller.Nurse_checking_feedback())

	incomingRoutes.GET("/check_nurse_appointment", controller.Check_nurse_appointment())

	incomingRoutes.POST("/nurse", controller.Add_nurse())

	incomingRoutes.PUT("/nurse", controller.Update_nurse())

	incomingRoutes.DELETE("/nurse", controller.Delete_nurse())
}
