// handler/report_handler.go
package handler

import (
	"bytes"
	"fmt"
	"go-service/internal/client"
	"go-service/internal/models"
	"go-service/internal/service/pdfsvc"
	"go-service/internal/service/studentsvc"
	"go-service/pkg/storage"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jung-kurt/gofpdf"
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

func GeneratePDFReportViaHttpClient(c *gin.Context) {
	id := c.Param("id")
	httpClient := client.NewClient(os.Getenv("STUDENT_API_BASE_URL")) // You can read baseURL from env/config

	student, err := httpClient.GetStudentByIDHttp(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch student via HTTP: " + err.Error()})
		return
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Student Report (via HTTP)")
	pdf.Ln(12)
	pdf.SetFont("Arial", "", 12)

	fields := map[string]string{
		"ID":                   student.ID,
		"Name":                 student.Name,
		"Email":                student.Email,
		"Phone":                student.Phone,
		"Gender":               student.Gender,
		"DOB":                  student.DOB,
		"Class":                student.Class,
		"Section":              student.Section,
		"Roll":                 student.Roll,
		"Father Name":          student.FatherName,
		"Father Phone":         student.FatherPhone,
		"Mother Name":          student.MotherName,
		"Mother Phone":         student.MotherPhone,
		"Guardian Name":        student.GuardianName,
		"Guardian Phone":       student.GuardianPhone,
		"Relation of Guardian": student.RelationOfGuardian,
		"Current Address":      student.CurrentAddress,
		"Permanent Address":    student.PermanentAddress,
		"Admission Date":       student.AdmissionDate,
		"Reporter Name":        student.ReporterName,
	}

	for label, value := range fields {
		pdf.CellFormat(50, 8, fmt.Sprintf("%s:", label), "", 0, "L", false, 0, "")
		pdf.CellFormat(100, 8, value, "", 1, "L", false, 0, "")
	}

	var buf bytes.Buffer
	if err := pdf.Output(&buf); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate PDF: " + err.Error()})
		return
	}

	c.Header("Content-Type", "application/pdf")
	c.Header("Content-Disposition", "attachment; filename=student_"+id+"_report_http.pdf")
	c.Data(http.StatusOK, "application/pdf", buf.Bytes())
}
