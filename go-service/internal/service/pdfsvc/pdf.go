package pdfsvc

import (
	"bytes"
	"fmt"
	"go-service/internal/repository/student"
	"go-service/pkg/storage"

	"github.com/jung-kurt/gofpdf"
)

type PDFService struct {
	StudentRepo *student.StudentRepository
}

func NewPDFService(store *storage.Store) *PDFService {
	return &PDFService{StudentRepo: student.NewStudentRepository(store)}
}

func (p *PDFService) GenerateStudentReport(id string) ([]byte, error) {
	student, err := p.StudentRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Student Report")
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

	buf := new(bytes.Buffer)
	err = pdf.Output(buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
