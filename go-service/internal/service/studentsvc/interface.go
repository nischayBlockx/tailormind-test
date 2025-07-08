package studentsvc

import "go-service/internal/models"

type StudentServiceInterface interface {
	GenerateReport(id string) ([]byte, error)
	InsertStudent(input StudentDetails) error
	UpdateStudent(input StudentDetails) error
	DeleteStudent(id string) error
	GetStudent(id string) (*models.StudentDetails, error)
}
