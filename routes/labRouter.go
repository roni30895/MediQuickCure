package routes

import (
	controller "Doctor-Appointment-Project/controllers"
	middleware "Doctor-Appointment-Project/middleware"

	"github.com/gin-gonic/gin"
)

func LabtRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())

	incomingRoutes.GET("/lab", controller.Get_Lab_Profile())

	incomingRoutes.POST("/lab", controller.Add_lab_details())

	incomingRoutes.PUT("/lab", controller.Update_lab())

	incomingRoutes.DELETE("/lab", controller.Delete_lab())

	incomingRoutes.DELETE("/cancel_appointment", controller.Cancel_lab_appointment())

}
