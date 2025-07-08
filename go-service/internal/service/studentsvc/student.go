package studentsvc

import (
	"bytes"
	"fmt"
	"go-service/internal/models"
	"go-service/internal/repository/student"
	"go-service/pkg/storage"

	"github.com/jung-kurt/gofpdf"
)

type StudentService struct {
	StudentRepo *student.StudentRepository
}

func NewStudentService(store *storage.Store) *StudentService {
	return &StudentService{StudentRepo: student.NewStudentRepository(store)}
}

func (s *StudentService) GenerateReport(id string) ([]byte, error) {
	student, err := s.StudentRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Student Report")
	pdf.Ln(12)
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(40, 10, fmt.Sprintf("ID: %s", student.ID))
	pdf.Ln(8)
	pdf.Cell(40, 10, fmt.Sprintf("Name: %s", student.Name))
	pdf.Ln(8)
	pdf.Cell(40, 10, fmt.Sprintf("Email: %s", student.Email))
	pdf.Ln(8)
	pdf.Cell(40, 10, fmt.Sprintf("Phone: %s", student.Phone))
	buf := new(bytes.Buffer)
	err = pdf.Output(buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (s *StudentService) InsertStudent(input *models.StudentDetails) error {
	return s.StudentRepo.Insert(input)
}

func (s *StudentService) UpdateStudent(input *models.StudentDetails) error {
	return s.StudentRepo.Update(input)
}

func (s *StudentService) DeleteStudent(id string) error {
	return s.StudentRepo.Delete(id)
}

func (s *StudentService) GetStudent(id string) (*models.StudentDetails, error) {
	return s.StudentRepo.GetByID(id)
}
