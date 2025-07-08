// handler/report_handler.go
package handler

import (
	"go-service/internal/models"
	"go-service/internal/service/pdfsvc"
	"go-service/internal/service/studentsvc"
	"go-service/pkg/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

var studentService *studentsvc.StudentService
var pdfService *pdfsvc.PDFService

func InitHandlers(store *storage.Store) {
	studentService = studentsvc.NewStudentService(store)
	pdfService = pdfsvc.NewPDFService(store)
}

func GeneratePDFReport(c *gin.Context) {
	id := c.Param("id")
	pdfBytes, err := pdfService.GenerateStudentReport(id)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.Header("Content-Type", "application/pdf")
	c.Header("Content-Disposition", "attachment; filename=student_"+id+"_report.pdf")
	c.Data(http.StatusOK, "application/pdf", pdfBytes)
}

func AddStudentData(c *gin.Context) {
	var studentInput *models.StudentDetails
	if err := c.ShouldBindJSON(&studentInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := studentService.InsertStudent(studentInput); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Student data added successfully"})
}

func UpdateStudentData(c *gin.Context) {
	var studentInput *models.StudentDetails
	if err := c.ShouldBindJSON(&studentInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := studentService.UpdateStudent(studentInput); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Student data updated successfully"})
}

func DeleteStudentData(c *gin.Context) {
	id := c.Param("id")
	if err := studentService.DeleteStudent(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Student data deleted successfully"})
}
