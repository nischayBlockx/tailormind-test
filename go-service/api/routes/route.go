package routes

import (
	"go-service/api/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/api/v1/students/:id/report", handler.GeneratePDFReport)
	r.POST("/api/v1/students", handler.AddStudentData)
	r.PUT("/api/v1/students", handler.UpdateStudentData)
	r.DELETE("/api/v1/students/:id", handler.DeleteStudentData)

	return r
}
